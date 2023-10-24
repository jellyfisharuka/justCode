package repository

import (
	"callboard/internal/user/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *Repo) CreateUser(ctx context.Context, user entity.User) error {
	q := `
INSERT INTO "user" (first_name, last_name, phone, login, password)
VALUES ($1, $2, $3, $4, $5);
`
	_, err := r.main.ExecContext(ctx, q, user.FirstName, user.LastName, user.Phone, user.Login, user.Password)
	if err != nil {
		return fmt.Errorf("db exec query failed: %w", err)
	}

	return nil
}

func (r *Repo) ConfirmUser(ctx context.Context, userId int) error {
	// Ваша логика для подтверждения пользователя в PostgreSQL
	return nil
}

func (r *Repo) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	q := `
SELECT id, first_name, last_name, phone, login, password, is_confirmed, is_deleted, created_at, updated_at
FROM "user"
WHERE login = $1;
`
	var row entity.User

	if err := r.replica.GetContext(ctx, &row, q, login); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("db query failed: %w", err)
	}

	return &row, nil
}