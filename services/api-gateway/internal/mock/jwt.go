package mock

import (
	"api_gateway_microservice/config"
	jwt_lib "api_gateway_microservice/internal/lib/jwt"
	"testing"

	"github.com/go-redis/redis/v8"
)

func NewTestJWTManager(t *testing.T, redisAddr string) *jwt_lib.JWTManager {
	t.Helper()

	cfg := &config.Config{}
	cfg.JWT.Secret = "test-secret"

	rdb := redis.NewClient(&redis.Options{Addr: redisAddr})
	t.Cleanup(func() {
		_ = rdb.Close()
	})

	return jwt_lib.NewJWTManager(cfg, rdb)
}
