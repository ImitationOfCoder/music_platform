package grpc

import (
	v1 "github.com/ImitationOfCoder/music_platform/proto/gen/go"
)

type ServerApi struct {
	v1.UnimplementedUserServiceServer
}
