package jwt_lib

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTClaims struct {
	JTI       string
	AccountId int64
	Email     string
	Exp       time.Duration
	jwt.RegisteredClaims
}

func NewJWTClaims(accountId int64, email string) *JWTClaims {
	jti := uuid.New().String()
	exp := 24 * time.Hour * 30

	return &JWTClaims{
		jti,
		accountId,
		email,
		exp,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}
