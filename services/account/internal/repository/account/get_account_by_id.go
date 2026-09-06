package account_repository

import (
	"account_microservice/internal/domain"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *AccountRepository) GetAccountById(ctx context.Context, id int64) (*domain.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, email, password_hash FROM music_platform.accounts WHERE id=$1;
	`

	var account domain.Account

	err := r.pool.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&account.Id,
		&account.Email,
		&account.PasswordHash,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrAccountNotFound
		}

		return nil, fmt.Errorf("AccountRepository - GetAccountById - Scan: %w", err)
	}

	return &account, nil
}
