package user_repository

import (
	"context"
	"errors"
	"fmt"
	"user_microservice/internal/domain"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *UserRepository) CreateUser(
	id int64,
	name string,
	email string,
	password string,
) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO music_platform.users (id, name, email, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email;
	`

	var userModel UserModel
	err := r.pool.QueryRow(
		ctx,
		query,
		id,
		name,
		email,
		password,
	).Scan(
		&userModel.Id,
		&userModel.Name,
		&userModel.Email,
	)
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "users_email_key" {
				return domain.ErrUserAlreadyExists
			}
		}

		return fmt.Errorf("UserRepository - CreateUser - Scan: %w", err)
	}

	return nil
}
