package repository

import (
	req "callboard/internal/auth/database/request"
	"callboard/internal/auth/helper"
	"callboard/internal/auth/models"
	"errors"

	"gorm.io/gorm"
)

type RepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements Repository.
func (u *RepositoryImpl) Delete(usersId int) {
	var users models.Users
	result := u.Db.Where("id=?", usersId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

// FindAll implements Repository.
func (u *RepositoryImpl) FindAll() []models.Users {
	var users []models.Users
	result := u.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements Repository.
func (u *RepositoryImpl) FindById(usersId int) (models.Users, error) {
	var users models.Users
	result := u.Db.Find(&users, usersId)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("user is not found")
	}
}

// FindByUsername implements Repository.
func (u *RepositoryImpl) FindByUsername(username string) (models.Users, error) {
	var users models.Users
	result := u.Db.First(&users, "user_name=?", username)
	if result.Error != nil {
		return users, errors.New("invalid username or password")
	}
	return users, nil
}

// Save implements Repository.
func (u *RepositoryImpl) Save(users models.Users) {
	result := u.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Update implements Repository.
func (u *RepositoryImpl) Update(users models.Users) {
	var updateUsers = req.UpdateUserRequest{
		ID:       users.ID,
		UserName: users.UserName,
		Email:    users.Email,
		Password: users.Password,
	}
	result := u.Db.Model(&users).Updates(updateUsers)
	helper.ErrorPanic(result.Error)
}

func NewUsersRepositoryImpl(DB *gorm.DB) Repository {
	return &RepositoryImpl{Db: DB}
}
