package user_repository

import (
	"context"
	"errors"
	"fmt"

	"user_microservice/internal/domain"
	"user_microservice/pkg/security"

	"github.com/jackc/pgx/v5"
)

func (r *UserRepository) Login(email string, password string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, password_hash FROM music_platform.users WHERE email=$1;
	`

	var userId int64
	var passwordHash string

	err := r.pool.QueryRow(ctx, query, email).Scan(&userId, &passwordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, domain.ErrInvalidCredentials
		}

		return 0, fmt.Errorf("UserRepository - Login - Scan: %w", err)
	}

	success, err := security.VerifyPassword(password, passwordHash)
	if err != nil {
		return 0, fmt.Errorf("UserRepository - Login - security.VerifyPassword: %w", err)
	}

	if !success {
		return 0, domain.ErrInvalidCredentials
	}

	return userId, nil
}
