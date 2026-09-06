package user_service_test

import (
	"api_gateway_microservice/internal/client"
	domain "api_gateway_microservice/internal/domain"
	mocks "api_gateway_microservice/internal/mock"
	user_service "api_gateway_microservice/internal/service/user"
	"api_gateway_microservice/pkg/logger"
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"github.com/stretchr/testify/require"
)

func newTestLogger() *logger.Logger {
	return &logger.Logger{
		Logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

func TestGetUserById(t *testing.T) {
	t.Parallel()

	successMock := &mocks.ProfileClientMock{
		GetProfileByIdFunc: func(ctx context.Context, id int64) (*v1.GetProfileByIdResponse, error) {
			return &v1.GetProfileByIdResponse{Id: id, Name: "Alice"}, nil
		},
	}

	notFoundMock := &mocks.ProfileClientMock{
		GetProfileByIdFunc: func(ctx context.Context, id int64) (*v1.GetProfileByIdResponse, error) {
			return nil, domain.ErrProfileNotFound
		},
	}

	unexpectedErr := errors.New("connection refused")
	unexpectedErrMock := &mocks.ProfileClientMock{
		GetProfileByIdFunc: func(ctx context.Context, id int64) (*v1.GetProfileByIdResponse, error) {
			return nil, unexpectedErr
		},
	}

	tests := []struct {
		name         string
		profileMock  *mocks.ProfileClientMock
		accountId    int64
		expectedUser domain.User
		expectedErr  error
	}{
		{
			name:         "success",
			profileMock:  successMock,
			accountId:    1,
			expectedUser: domain.User{Id: 1, Name: "Alice"},
			expectedErr:  nil,
		},
		{
			name:         "profile not found",
			profileMock:  notFoundMock,
			accountId:    2,
			expectedUser: domain.User{},
			expectedErr:  domain.ErrUserNotFound,
		},
		{
			name:         "unexpected error is wrapped",
			profileMock:  unexpectedErrMock,
			accountId:    3,
			expectedUser: domain.User{},
			expectedErr:  nil, // checked via require.ErrorIs on the wrapped chain below
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			svc := user_service.New(newTestLogger(), &client.GrpcClients{
				Profile: tt.profileMock,
			})

			user, err := svc.GetUserById(context.Background(), tt.accountId)

			require.Equal(t, 1, tt.profileMock.GetProfileByIdCalls)
			require.Equal(t, tt.accountId, tt.profileMock.GetProfileByIdId)

			if tt.name == "unexpected error is wrapped" {
				require.Error(t, err)
				require.ErrorIs(t, err, unexpectedErr)
				require.NotErrorIs(t, err, domain.ErrUserNotFound)

				return
			}

			require.ErrorIs(t, err, tt.expectedErr)
			require.Equal(t, tt.expectedUser, user)
		})
	}
}
