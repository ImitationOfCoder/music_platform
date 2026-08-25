package user_service

import (
	"errors"
	"fmt"

	"user_microservice/internal/domain"
	"user_microservice/pkg/security"
)

func (s *UserService) CreateUser(name string, email string, password string) error {
	id := s.snowflake.GenerateID()

	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return fmt.Errorf("UserService - CreateUser - security.HashPassword: %w", err)
	}

	err = s.userRepository.CreateUser(id, name, email, hashedPassword)
	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return domain.ErrUserAlreadyExists
		}

		return fmt.Errorf("UserService - CreateUser - s.userRepository.CreateUser: %w", err)
	}

	return nil
}
