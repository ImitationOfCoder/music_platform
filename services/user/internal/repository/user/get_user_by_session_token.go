package user_repository

import (
	"context"
	"errors"
	"fmt"
	"user_microservice/internal/domain"

	"github.com/jackc/pgx/v5"
)

func (r *UserRepository) GetUserBySessionToken(sessionToken string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT u.id, u.name, u.email FROM music_platform.users AS u
		JOIN music_platform.sessions AS s ON u.id = s.user_id
		WHERE s.token=$1;
	`

	var user domain.User
	err := r.pool.QueryRow(
		ctx,
		query,
		sessionToken,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}

		return domain.User{}, fmt.Errorf("UserRepository - GetUserBySessionToken - Scan: %w", err)
	}

	err = user.Validate()
	if err != nil {
		return domain.User{}, fmt.Errorf("UserRepository - GetUserBySessionToken - Validate: %w", err)
	}

	return user, nil
}
