package repository

import (
	"callboard/internal/advert/models"

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func (r *Repository) CreateAd(c *fiber.Ctx) error {
	advert := models.Advertisement{}
     
	err := c.BodyParser(&advert)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&advert).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create advertisement"})
		return err
	}
	var AdModel []models.Advertisement
	if err := r.DB.Find(&AdModel).Error; err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get advertisements"})
		return err
	}
	// Сериализация и сохранение в Redis
	serializedData, err := json.Marshal(AdModel)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "could not serialize data"})
		return err
	}
	ctx := c.Context()
	// Сохранить сериализованные данные книг в Redis
	err = r.RedisClient.Set(ctx, "AdModel", serializedData, 2*time.Minute).Err()

	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "could not set data in Redis"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "advertisement has been added"})
	return nil
}

func (r *Repository) DeleteAd(context *fiber.Ctx) error {
	AdModel := models.Advertisement{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(AdModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete advertisement",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "advertisement delete successfully",
	})
	return nil
}

func (r *Repository) GetAds(context *fiber.Ctx) error {
	ctx := context.Context()
	adData, err := r.RedisClient.Get(ctx, "AdModel").Result()
	if err != nil {
		// Если данных нет в Redis, получить данные из Postgres
		var AdModel []models.Advertisement
		if err := r.DB.Find(&AdModel).Error; err != nil {
			context.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"message": "could not get advertisements"})
			return err
		}

		// Сериализация и сохранение в Redis
		serializedData, err := json.Marshal(AdModel)
		if err != nil {
			context.Status(http.StatusInternalServerError).JSON(
				&fiber.Map{"message": "could not serialize data"})
			return err
		}

		ctx := context.Context()

		err = r.RedisClient.Set(ctx, "AdModel", serializedData, 2*time.Minute).Err()

		if err != nil {
			context.Status(http.StatusInternalServerError).JSON(
				&fiber.Map{"message": "could not set data in Redis"})
			return err
		}

		// Вернуть данные
		context.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "advertisements fetched successfully",
			"data":    AdModel,
		})
	} else {
		// Вернуть данные из Redis
		var AdModel []models.Advertisement
		err = json.Unmarshal([]byte(adData), &AdModel)
		if err != nil {
			context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get advertisements"})
			return err
		}

		context.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "advertisements fetched successfully",
			"data":    AdModel,
		})
	}
	return nil
}

func (r *Repository) GetAdByID(context *fiber.Ctx) error {

	id := context.Params("id")
	AdModel := &models.Advertisement{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(AdModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the advertisement"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "advertisement id fetched successfully",
		"data":    AdModel,
	})
	return nil
}
func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Адрес Redis сервера
		Password: "",               // Пароль (если есть)
		DB:       0,                // Номер базы данных
	})
	return client
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_ad", r.CreateAd)
	api.Delete("/delete_ad/:id", r.DeleteAd)
	api.Get("/get_ad/:id", r.GetAdByID)
	api.Get("/ads", r.GetAds)
}