package account_repository

import (
	"account_microservice/internal/domain"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *AccountRepository) GetAccountByEmail(ctx context.Context, email string) (*domain.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, email, password_hash FROM music_platform.accounts WHERE email=$1;
	`

	var account domain.Account

	err := r.pool.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&account.Id,
		&account.Email,
		&account.PasswordHash,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrAccountNotFound
		}

		return nil, fmt.Errorf("AccountRepository - GetAccountByEmail - Scan: %w", err)
	}

	return &account, nil
}
