package jwt_lib

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type JWTBlacklist struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewJWTBlacklist(redisClient *redis.Client) *JWTBlacklist {
	return &JWTBlacklist{
		redisClient: redisClient,
		ctx:         context.Background(),
	}
}

func (b *JWTBlacklist) AddToBlacklistWithJTI(claims *JWTClaims) error {
	key := fmt.Sprintf("blacklist:jti:%s", claims.JTI)
	exp := claims.Exp

	return b.redisClient.Set(b.ctx, key, "revoked", exp).Err()
}

func (b *JWTBlacklist) IsBlacklisted(jti string) (bool, error) {
	key := fmt.Sprintf("blacklist:jti:%s", jti)
	val, err := b.redisClient.Get(b.ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil // Токен не в черном списке
		}

		return false, err
	}

	return val == "revoked", nil
}
