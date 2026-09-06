package account_service

import (
	"account_microservice/internal/domain"
	"account_microservice/pkg/security"
	"context"
	"errors"
	"fmt"
)

func (s *AccountService) CreateAccount(ctx context.Context, name string, email string, password string) error {
	// validate data

	id := s.snowflake.GenerateID()

	// Hash password
	passwordHash, err := security.HashPassword(password)
	if err != nil {
		return fmt.Errorf("AuthService - CreateAccount - security.HashPassword: %w", err)
	}

	// Create Account
	err = s.accountRepository.CreateAccount(ctx, id, email, passwordHash)
	if err != nil {
		if errors.Is(err, domain.ErrAccountAlreadyExists) {
			return domain.ErrAccountAlreadyExists
		}

		return fmt.Errorf("AccountService - CreateAccount - s.accountRepository.CreateAccount: %w", err)
	}

	// Create Profile
	_, err = s.profileServiceClient.CreateProfile(ctx, id, name)
	if err != nil {
		return fmt.Errorf("AccountService - CreateAccount - s.profileServiceClient.CreateProfile: %w", err)
	}

	return nil
}
