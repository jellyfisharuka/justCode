package database

type CreateUserRequest struct {
	UserName string `validate:"required, min=2, max=100" json:"username"`
	Email    string `validate:"required, min=2, max=100" json:"email"`
	Password string `validate:"required, min=2, max=100" json:"password"`
}
type UpdateUserRequest struct {
	ID       int    `validate:"required"`
	UserName string `validate:"required, min=2, max=100" json:"username"`
	Email    string `validate:"required, min=2, max=100" json:"email"`
	Password string `validate:"required, min=2, max=100" json:"password"`
}
type LoginRequest struct {
	UserName string `validate:"required, min=200, max=2" json:"username"`
	Password string `validate:"required, min=2, max=100" json:"password"`
}
