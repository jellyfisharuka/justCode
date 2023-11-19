package database

type UserResponse struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
