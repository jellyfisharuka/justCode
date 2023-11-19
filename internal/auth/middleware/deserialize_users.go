package middleware

import (
	"callboard/internal/auth/config"
	"callboard/internal/auth/helper"
	"callboard/internal/auth/repository"
	"callboard/internal/auth/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeserializeUser(userRepository repository.Repository) gin.HandlerFunc{
	return func (ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)
		if len(fields) != 0 && fields[0]=="Bearer" {
			token=fields[1]
		}
		if token=="" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status":"fail","message":"you are not logged in"})
			return 
		}
		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err!=nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"status":"fail", "message":err.Error()})
		}
		id, err_id := strconv.Atoi(fmt.Sprint(sub))
		helper.ErrorPanic(err_id)
		result, err := userRepository.FindById(id)
		if err!=nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status":"fail", "message":"the user belonging to this token no longer exist"})
			return
		}
		ctx.Set("currentUser", result.UserName)
		ctx.Next()
		
	}

}