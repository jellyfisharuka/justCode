package controller

import (
	request "callboard/internal/auth/database/request"
	response "callboard/internal/auth/database/response"
	"callboard/internal/auth/helper"
	"callboard/internal/auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationService service.AuthenticationService
}

func NewAuthenticationController(service service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: service}
}
func (controller *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.ErrorPanic(err)
	token, err_token := controller.authenticationService.Login(loginRequest)
	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username>>or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUsersRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	helper.ErrorPanic(err)

	controller.authenticationService.Register(createUsersRequest)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
