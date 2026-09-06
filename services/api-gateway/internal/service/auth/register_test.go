package auth_service_test

import (
	"api_gateway_microservice/internal/client"
	domain "api_gateway_microservice/internal/domain"
	mocks "api_gateway_microservice/internal/mock"
	auth_service "api_gateway_microservice/internal/service/auth"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	t.Parallel()

	successMock := &mocks.AccountClientMock{
		CreateAccountFunc: func(ctx context.Context, name string, email string, password string) error {
			return nil
		},
	}

	alreadyExistsMock := &mocks.AccountClientMock{
		CreateAccountFunc: func(ctx context.Context, name string, email string, password string) error {
			return domain.ErrAccountAlreadyExists
		},
	}

	unexpectedErr := errors.New("connection refused")
	unexpectedErrMock := &mocks.AccountClientMock{
		CreateAccountFunc: func(ctx context.Context, name string, email string, password string) error {
			return unexpectedErr
		},
	}

	tests := []struct {
		name          string
		accountMock   *mocks.AccountClientMock
		expectedErr   error
		expectedCalls int
	}{
		{
			name:          "success",
			accountMock:   successMock,
			expectedErr:   nil,
			expectedCalls: 1,
		},
		{
			name:          "account already exists",
			accountMock:   alreadyExistsMock,
			expectedErr:   domain.ErrAccountAlreadyExists,
			expectedCalls: 1,
		},
		{
			name:          "unexpected error is wrapped",
			accountMock:   unexpectedErrMock,
			expectedErr:   nil, // checked via require.ErrorIs on the wrapped chain below
			expectedCalls: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			svc := auth_service.New(mocks.NewTestLogger(), nil, &client.GrpcClients{
				Account: tt.accountMock,
			})

			err := svc.Register(context.Background(), "Alice", "user@example.com", "password")

			require.Equal(t, tt.expectedCalls, tt.accountMock.CreateAccountCalls)

			if tt.name == "unexpected error is wrapped" {
				require.Error(t, err)
				require.ErrorIs(t, err, unexpectedErr)
				require.NotErrorIs(t, err, domain.ErrAccountAlreadyExists)

				return
			}

			require.ErrorIs(t, err, tt.expectedErr)
		})
	}
}
