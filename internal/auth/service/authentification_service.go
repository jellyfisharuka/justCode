package service

import (
    req "callboard/internal/auth/database/request"
	
)

type AuthenticationService interface {
	Login(users req.LoginRequest)(string, error) 
	Register(users req.CreateUserRequest)
}
