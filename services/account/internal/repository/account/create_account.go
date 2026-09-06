package account_repository

import (
	"account_microservice/internal/domain"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *AccountRepository) CreateAccount(ctx context.Context, accountId int64, email string, passwordHash string) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO music_platform.accounts (id, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		accountId,
		email,
		passwordHash,
	)
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "accounts_email_key" {
				return domain.ErrAccountAlreadyExists
			}
		}

		return fmt.Errorf("AccountRepository - CreateAccount - Scan: %w", err)
	}

	return nil
}
