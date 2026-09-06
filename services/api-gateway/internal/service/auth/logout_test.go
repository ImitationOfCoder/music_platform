package auth_service_test

import (
	jwt_lib "api_gateway_microservice/internal/lib/jwt"
	"api_gateway_microservice/internal/mock"
	auth_service "api_gateway_microservice/internal/service/auth"
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/require"
)

func TestLogout(t *testing.T) {
	t.Parallel()

	mr := miniredis.RunT(t)
	jwtManager := mock.NewTestJWTManager(t, mr.Addr())

	svc := auth_service.New(mock.NewTestLogger(), jwtManager, nil)

	claims := jwt_lib.NewJWTClaims(42, "user@example.com")

	err := svc.Logout(context.Background(), claims)
	require.NoError(t, err)

	blacklisted, err := jwtManager.IsBlacklisted(claims.JTI)
	require.NoError(t, err)
	require.True(t, blacklisted)
}

func TestLogoutRedisError(t *testing.T) {
	t.Parallel()

	mr := miniredis.RunT(t)
	mr.SetError("redis is down")

	jwtManager := mock.NewTestJWTManager(t, mr.Addr())

	svc := auth_service.New(mock.NewTestLogger(), jwtManager, nil)

	claims := jwt_lib.NewJWTClaims(42, "user@example.com")

	err := svc.Logout(context.Background(), claims)
	require.Error(t, err)
	require.Contains(t, err.Error(), "AddToBlacklist")
}
