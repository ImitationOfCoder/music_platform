package account_service

import (
	"account_microservice/internal/domain"
	"context"
	"fmt"
)

func (s *AccountService) GetAccountById(ctx context.Context, id int64) (*domain.Account, error) {
	account, err := s.accountRepository.GetAccountById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("AccountService - GetAccountById - s.accountRepository.GetAccountByEmail: %w", err)
	}

	return account, nil
}
