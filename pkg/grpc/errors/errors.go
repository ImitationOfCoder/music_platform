package grpc_client

import (
	"errors"
)

var (
	ErrAccountNotFound      = errors.New("account not found")
	ErrAccountAlreadyExists = errors.New("account already exists")

	ErrProfileNotFound = errors.New("profile not found")

	ErrUnauthenticated    = errors.New("unauthenticated")
	ErrInvalidCredentials = errors.New("invalid credentials")

	ErrUserNotFound = errors.New("user not found")
)
