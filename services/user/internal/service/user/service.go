package user_service

import (
	"user_microservice/internal/domain"
	"user_microservice/pkg/snowflake"
)

type UserService struct {
	userRepository UserRepository
	snowflake      *snowflake.SnowflakeGenerator
}

type UserRepository interface {
	CreateUser(
		id int64,
		name string,
		email string,
		password string,
	) error
	GetUserById(id int64) (domain.User, error)
}

func NewService(
	userRepository UserRepository,
	snowflake *snowflake.SnowflakeGenerator,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		snowflake:      snowflake,
	}
}
