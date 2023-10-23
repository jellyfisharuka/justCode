package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/jellyfisharuka/go-fiber-postgres/models"
	"github.com/jellyfisharuka/go-fiber-postgres/storage"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}
	var bookModel []models.Books
	if err := r.DB.Find(&bookModel).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}
	// Сериализация и сохранение в Redis
	serializedData, err := json.Marshal(bookModel)
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "could not serialize data"})
		return err
	}

	// Сохранить сериализованные данные книг в Redis
	err = r.RedisClient.Set(context.Context(), "bookModel", serializedData, 2*time.Minute).Err()
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "could not set data in Redis"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has been added"})
	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(bookModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete book",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book delete successfully",
	})
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	ctx := context.Context()
	booksData, err := r.RedisClient.Get(ctx, "bookModel").Result()
	if err != nil {
		// Если данных нет в Redis, получить данные из Postgres
		var bookModel []models.Books
		if err := r.DB.Find(&bookModel).Error; err != nil {
			context.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"message": "could not get books"})
			return err
		}

		// Сериализация и сохранение в Redis
		serializedData, err := json.Marshal(bookModel)
		if err != nil {
			context.Status(http.StatusInternalServerError).JSON(
				&fiber.Map{"message": "could not serialize data"})
			return err
		}

		err = r.RedisClient.Set(ctx, "bookModel", serializedData, 2*time.Minute).Err()
		if err != nil {
			context.Status(http.StatusInternalServerError).JSON(
				&fiber.Map{"message": "could not set data in Redis"})
			return err
		}

		// Вернуть данные
		context.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "books fetched successfully",
			"data":    bookModel,
		})
	} else {
		// Вернуть данные из Redis
		var bookModel []models.Books
		err = json.Unmarshal([]byte(booksData), &bookModel)
		if err != nil {
			context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get books"})
			return err
		}

		context.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "books fetched successfully",
			"data":    bookModel,
		})
	}
	return nil
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {

	id := context.Params("id")
	bookModel := &models.Books{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book id fetched successfully",
		"data":    bookModel,
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}
	redisClient := NewRedisClient()
	r := Repository{
		DB:          db,
		RedisClient: redisClient,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Адрес Redis сервера
		Password: "",               // Пароль (если есть)
		DB:       0,                // Номер базы данных
	})
	return client
}
