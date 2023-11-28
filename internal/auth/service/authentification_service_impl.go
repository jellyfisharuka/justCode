package service

import (
	"callboard/internal/auth/config"
	request "callboard/internal/auth/database/request"
	"callboard/internal/auth/helper"
	"callboard/internal/auth/models"
	"callboard/internal/auth/repository"
	"callboard/internal/auth/utils"
	"errors"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.Repository
	Validate        *validator.Validate
}

func NewAuthenticationServiceImpl(usersRepository repository.Repository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// Login implements AuthenticationService.
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	//finding the username in database
	new_user, user_err := a.UsersRepository.FindByUsername(users.UserName)
	if user_err != nil {
		return "", errors.New("invalid username")
	}
	config, _ := config.LoadConfig(".")
	verify_error := utils.VerifyPassword(new_user.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid password")

	}
	//generate token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_user.ID, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil

}

// Register implements AuthenticationService.
func (a *AuthenticationServiceImpl) Register(users request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)
	newUser := models.Users{
		UserName: users.UserName,
		Email:    users.Email,
		Password: hashedPassword,
	}
	a.UsersRepository.Save(newUser)
}
