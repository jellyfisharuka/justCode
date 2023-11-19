package repository

import (
	"callboard/internal/auth/models"
	"errors"

)



type Repository interface {
	Save(users models.Users)
	Update(users models.Users)
	Delete(usersId int)
	FindById(usersId int) (models.Users, error)
	FindAll() []models.Users
	FindByUsername(username string) (models.Users, error)
}

var (
	ErrNotFound = errors.New("not found")
)


