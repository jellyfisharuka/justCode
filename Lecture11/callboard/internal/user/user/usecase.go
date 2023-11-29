package user

import (
	"context"
	"callboard/internal/user/entity"
)

type UseCase interface {
	GetUserByLogin(ctx context.Context, login string) (*entity.User, error)
}
