package main

import (
	"callboard/internal/auth/config"
	"callboard/internal/auth/controller"
	"callboard/internal/auth/helper"
	"callboard/internal/auth/models"
	"callboard/internal/auth/repository"
	"callboard/internal/auth/service"
	"log"
	"github.com/go-playground/validator/v10"
	"net/http"
	router "callboard/internal/auth/server/router"
)

func main() {
	
	loadConfig, err := config.LoadConfig(".")
   if err != nil {
    log.Fatal("Could not load environment variables", err)
   }
	//db

db:=config.ConnectionDB(&loadConfig)
validate := validator.New()
db.Table("users").AutoMigrate(&models.Users{})
//init repository
userRepository := repository.NewUsersRepositoryImpl(db)
//init service
authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)
//init controller
authenticationController := controller.NewAuthenticationController(authenticationService)
usersController := controller.NewUsersController(userRepository)
routes := router.NewRouter(userRepository, authenticationController, usersController)
  server := &http.Server{
	Addr:    ":8888",
	Handler: routes,
}
server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
   }

 

