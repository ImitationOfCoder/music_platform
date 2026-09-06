package account_service

import (
	"context"
	"errors"
	"fmt"

	"account_microservice/internal/domain"
	"account_microservice/pkg/security"
)

func (s *AccountService) Login(ctx context.Context, email string, password string) (*domain.Account, error) {
	account, err := s.accountRepository.GetAccountByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrAccountNotFound) {
			return nil, domain.ErrInvalidCredentials
		}

		return nil, fmt.Errorf("AccountService - Login - s.accountRepository.GetAccountByEmail: %w", err)
	}

	ok, err := security.VerifyPasswordHash(password, account.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("AccountService - Login - security.VerifyPassword: %w", err)
	}

	if !ok {
		return nil, domain.ErrInvalidCredentials
	}

	return &domain.Account{
		Id:           account.Id,
		Email:        account.Email,
		PasswordHash: account.PasswordHash,
	}, nil
}
