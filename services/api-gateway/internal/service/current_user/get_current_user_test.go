package current_user_service_test

import (
	"api_gateway_microservice/internal/client"
	domain "api_gateway_microservice/internal/domain"
	jwt_lib "api_gateway_microservice/internal/lib/jwt"
	mocks "api_gateway_microservice/internal/mock"
	current_user_service "api_gateway_microservice/internal/service/current_user"
	"context"
	"errors"
	"testing"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"github.com/stretchr/testify/require"
)

func TestGetCurrentUser(t *testing.T) {
	t.Parallel()

	successAccountMock := func() *mocks.AccountClientMock {
		return &mocks.AccountClientMock{
			GetAccountByIdFunc: func(ctx context.Context, accountId int64) (*v1.GetAccountByIdResponse, error) {
				return &v1.GetAccountByIdResponse{Id: accountId, Email: "user@example.com"}, nil
			},
		}
	}
	successProfileMock := func() *mocks.ProfileClientMock {
		return &mocks.ProfileClientMock{
			GetProfileByIdFunc: func(ctx context.Context, id int64) (*v1.GetProfileByIdResponse, error) {
				return &v1.GetProfileByIdResponse{Id: id, Name: "Alice"}, nil
			},
		}
	}

	accountErrMock := func() *mocks.AccountClientMock {
		return &mocks.AccountClientMock{
			GetAccountByIdFunc: func(ctx context.Context, accountId int64) (*v1.GetAccountByIdResponse, error) {
				return nil, errors.New("connection refused")
			},
		}
	}
	profileErrMock := func() *mocks.ProfileClientMock {
		return &mocks.ProfileClientMock{
			GetProfileByIdFunc: func(ctx context.Context, id int64) (*v1.GetProfileByIdResponse, error) {
				return nil, errors.New("connection refused")
			},
		}
	}

	tests := []struct {
		name           string
		accountMock    func() *mocks.AccountClientMock
		profileMock    func() *mocks.ProfileClientMock
		expectedUser   domain.CurrentUser
		expectError    bool
		expectedErr    error
		expectAcctCall bool
		expectProfCall bool
	}{
		{
			name:           "success",
			accountMock:    successAccountMock,
			profileMock:    successProfileMock,
			expectedUser:   domain.CurrentUser{Id: 42, Name: "Alice", Email: "user@example.com"},
			expectAcctCall: true,
			expectProfCall: true,
		},
		{
			name:           "account lookup fails",
			accountMock:    accountErrMock,
			profileMock:    successProfileMock,
			expectError:    true,
			expectedErr:    domain.ErrUnauthenticated,
			expectAcctCall: true,
			expectProfCall: false,
		},
		{
			name:           "profile lookup fails",
			accountMock:    successAccountMock,
			profileMock:    profileErrMock,
			expectError:    true,
			expectedErr:    domain.ErrUnauthenticated,
			expectAcctCall: true,
			expectProfCall: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			accountMock := tt.accountMock()
			profileMock := tt.profileMock()

			svc := current_user_service.New(nil, &client.GrpcClients{
				Account: accountMock,
				Profile: profileMock,
			})

			claims := jwt_lib.NewJWTClaims(42, "user@example.com")

			user, err := svc.GetCurrentUser(context.Background(), claims)

			expectedAcctCalls, expectedProfCalls := 0, 0
			if tt.expectAcctCall {
				expectedAcctCalls = 1
			}
			if tt.expectProfCall {
				expectedProfCalls = 1
			}

			require.Equal(t, expectedAcctCalls, accountMock.GetAccountByIdCalls)
			require.Equal(t, expectedProfCalls, profileMock.GetProfileByIdCalls)

			if tt.expectError {
				require.ErrorIs(t, err, tt.expectedErr)
				require.Equal(t, domain.CurrentUser{}, user)

				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expectedUser, user)
		})
	}
}
