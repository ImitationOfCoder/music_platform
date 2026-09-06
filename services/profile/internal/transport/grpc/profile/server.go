package profile_grpc

import (
	"user_microservice/internal/domain"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc"
)

type ProfileService interface {
	CreateProfile(accountId int64, name string) error
	GetProfileById(accountId int64) (*domain.User, error)
}

type ServerApi struct {
	v1.UnimplementedProfileServiceServer
	profileService ProfileService
}

func Register(gRPC *grpc.Server, profileService ProfileService) {
	v1.RegisterProfileServiceServer(gRPC, &ServerApi{
		profileService: profileService,
	})
}
