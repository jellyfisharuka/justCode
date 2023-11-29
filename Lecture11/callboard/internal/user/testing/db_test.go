package testing

import (
	"callboard/internal/user/database"
	"callboard/internal/user/entity"
	"callboard/internal/user/repository"

	"context"
	"testing"
)

func TestCreateUserAndGetUser(t *testing.T) {
	// Инициализация репозитория и базы данных
	mainDB := database.NewMyDatabase()
	replicaDB := database.NewMyDatabase()

	repo := repository.NewRepository(mainDB, replicaDB) // Замените на инициализацию вашего репозитория
	ctx := context.Background()

	// Создание тестового пользователя
	user := entity.User{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "123456789",
		Login:     "johndoe",
		Password:  "password123",
	}

	err := repo.CreateUser(ctx, user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Получение созданного пользователя
	retrievedUser, err := repo.GetUserByLogin(ctx, "johndoe")
	if err != nil {
		t.Fatalf("Failed to retrieve user: %v", err)
	}

	// Проверка, что полученные данные соответствуют ожидаемым
	if retrievedUser.FirstName != "John" || retrievedUser.LastName != "Doe" {
		t.Errorf("User data doesn't match the expected values")
	}
}
