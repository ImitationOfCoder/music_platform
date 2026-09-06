package profile_service

import (
	"errors"
	"fmt"
	"user_microservice/internal/domain"
)

func (s *ProfileService) GetProfileById(profileId int64) (*domain.User, error) {
	user, err := s.profileRepository.GetProfileById(profileId)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, domain.ErrUserNotFound
		}

		return nil, fmt.Errorf("ProfileService - GetProfileById - s.profileRepository.GetProfileById: %w", err)
	}

	return user, nil
}
