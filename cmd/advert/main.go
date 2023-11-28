package main

import (
	"callboard/internal/advert/config"
	"callboard/internal/advert/models"
	"callboard/internal/advert/repository"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Advertisement struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    Category    string `json:"category"`
    Price       float64 `json:"price"`
    // Другие поля, связанные с объявлениями
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
	err = models.MigrateAdv(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}
	redisClient := NewRedisClient()
	r := repository.Repository{
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
