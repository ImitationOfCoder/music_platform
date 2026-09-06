package profile_service

import (
	"errors"
	"fmt"

	"user_microservice/internal/domain"
)

func (s *ProfileService) CreateProfile(accountId int64, name string) error {
	err := s.profileRepository.CreateProfile(accountId, name)
	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return domain.ErrUserAlreadyExists
		}

		return fmt.Errorf("ProfileService - CreateProfile - s.profileRepository.CreateProfile: %w", err)
	}

	return nil
}
