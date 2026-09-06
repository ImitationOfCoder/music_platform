package current_user_grpc

import (
	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc"
)

type ServerApi struct {
	v1.UnimplementedCurrentUserServiceServer
}

func Register(gRPC *grpc.Server) {
	v1.RegisterCurrentUserServiceServer(gRPC, &ServerApi{})
}
