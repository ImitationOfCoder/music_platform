package profile_service

import (
	"user_microservice/internal/domain"
)

type ProfileService struct {
	profileRepository ProfileRepository
}

type ProfileRepository interface {
	CreateProfile(accountId int64, name string) error
	GetProfileById(id int64) (*domain.User, error)
}

func NewService(
	profileRepository ProfileRepository,
) *ProfileService {
	return &ProfileService{
		profileRepository: profileRepository,
	}
}
