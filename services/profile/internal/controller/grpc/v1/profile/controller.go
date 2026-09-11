package v1

import (
	"profile_microservice/internal/service"

	"github.com/ImitationOfCoder/music_platform/pkg/logger"
	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
)

type ProfileController struct {
	v1.UnimplementedProfileServiceServer

	log            *logger.Logger
	profileService service.ProfileService
}
