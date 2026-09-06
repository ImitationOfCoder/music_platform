package profile_repository

import (
	"context"
	"errors"
	"fmt"
	"user_microservice/internal/domain"

	"github.com/jackc/pgx/v5"
)

func (r *ProfileRepository) GetProfileById(id int64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT account_id, name FROM music_platform.profiles WHERE account_id=$1;
	`

	var user domain.User
	err := r.pool.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&user.Id,
		&user.Name,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, fmt.Errorf("UserRepository - GetUserById - Scan: %w", err)
	}

	err = user.Validate()
	if err != nil {
		return nil, fmt.Errorf("UserRepository - GetUserById - Validate: %w", err)
	}

	return &user, nil
}
