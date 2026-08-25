package user_service

import (
	"fmt"
	"user_microservice/internal/domain"
)

func (s *UserService) GetUserById(userId int64) (domain.User, error) {
	user, err := s.userRepository.GetUserById(userId)
	if err != nil {
		return domain.User{}, fmt.Errorf("UserService - GetUserById - s.userRepository.GetUserById: %w", err)
	}

	return user, nil
}
