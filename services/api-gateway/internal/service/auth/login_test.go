package auth_service_test

import (
	"api_gateway_microservice/internal/client"
	domain "api_gateway_microservice/internal/domain"
	mocks "api_gateway_microservice/internal/mock"
	auth_service "api_gateway_microservice/internal/service/auth"
	"context"
	"errors"
	"testing"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLogin(t *testing.T) {
	t.Parallel()

	successMock := &mocks.AccountClientMock{
		LoginFunc: func(ctx context.Context, email string, password string) (*v1.LoginResponse, error) {
			return &v1.LoginResponse{Id: 42, Email: email}, nil
		},
	}

	invalidCredentialsMock := &mocks.AccountClientMock{
		LoginFunc: func(ctx context.Context, email string, password string) (*v1.LoginResponse, error) {
			return nil, status.Error(codes.Unauthenticated, "wrong password")
		},
	}

	internalErrMock := &mocks.AccountClientMock{
		LoginFunc: func(ctx context.Context, email string, password string) (*v1.LoginResponse, error) {
			return nil, status.Error(codes.Internal, "db is down")
		},
	}

	unexpectedErr := errors.New("connection refused")
	unexpectedErrMock := &mocks.AccountClientMock{
		LoginFunc: func(ctx context.Context, email string, password string) (*v1.LoginResponse, error) {
			return nil, unexpectedErr
		},
	}

	tests := []struct {
		name        string
		accountMock *mocks.AccountClientMock
		email       string
		password    string
		expectedErr error
	}{
		{
			name:        "success",
			accountMock: successMock,
			email:       "user@example.com",
			password:    "password",
			expectedErr: nil,
		},
		{
			name:        "invalid credentials",
			accountMock: invalidCredentialsMock,
			email:       "user@example.com",
			password:    "wrong-password",
			expectedErr: domain.ErrInvalidCredentials,
		},
		{
			name:        "internal grpc error is wrapped",
			accountMock: internalErrMock,
			email:       "user@example.com",
			password:    "password",
			expectedErr: nil, // checked via require.Error below
		},
		{
			name:        "non-grpc error is wrapped",
			accountMock: unexpectedErrMock,
			email:       "user@example.com",
			password:    "password",
			expectedErr: nil, // checked via require.Error below
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			jwtManager := mocks.NewTestJWTManager(t, "")
			svc := auth_service.New(mocks.NewTestLogger(), jwtManager, &client.GrpcClients{
				Account: tt.accountMock,
			})

			token, err := svc.Login(context.Background(), tt.email, tt.password)

			require.Equal(t, 1, tt.accountMock.LoginCalls)

			switch tt.name {
			case "success":
				require.NoError(t, err)
				require.NotEmpty(t, token)

				claims, parseErr := jwtManager.ParseToken(token)
				require.NoError(t, parseErr)
				require.Equal(t, int64(42), claims.AccountId)
				require.Equal(t, "user@example.com", claims.Email)
			case "internal grpc error is wrapped":
				require.Error(t, err)
				require.Contains(t, err.Error(), "db is down")
				require.Empty(t, token)
			case "non-grpc error is wrapped":
				require.Error(t, err)
				require.ErrorIs(t, err, unexpectedErr)
				require.Empty(t, token)
			default:
				require.ErrorIs(t, err, tt.expectedErr)
				require.Empty(t, token)
			}
		})
	}
}
