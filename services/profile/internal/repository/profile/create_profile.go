package profile_repository

import (
	"context"
	"errors"
	"fmt"
	"user_microservice/internal/domain"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *ProfileRepository) CreateProfile(accountId int64, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO music_platform.profiles (account_id, name) VALUES ($1, $2);
	`

	var user domain.User
	err := r.pool.QueryRow(
		ctx,
		query,
		accountId,
		name,
	).Scan(
		&user.Id,
		&user.Name,
	)
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "profiles_account_id_key" {
				return domain.ErrUserAlreadyExists
			}
		}

		return fmt.Errorf("ProfileRepository - CreateProfile - Scan: %w", err)
	}

	return nil
}
