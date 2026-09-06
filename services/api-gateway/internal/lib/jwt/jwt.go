package jwt_lib

import (
	"api_gateway_microservice/config"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret    string
	blacklist *JWTBlacklist
}

func NewJWTManager(cfg *config.Config, redis *redis.Client) *JWTManager {
	blacklist := NewJWTBlacklist(redis)

	return &JWTManager{
		secret:    cfg.JWT.Secret,
		blacklist: blacklist,
	}
}

func (j *JWTManager) NewToken(accountId int64, email string) (string, error) {
	claims := NewJWTClaims(accountId, email)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (j *JWTManager) ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims.(*JWTClaims), nil
}

func (j *JWTManager) AddToBlacklist(claims *JWTClaims) error {
	return j.blacklist.AddToBlacklistWithJTI(claims)
}

func (j *JWTManager) IsBlacklisted(jti string) (bool, error) {
	return j.blacklist.IsBlacklisted(jti)
}
