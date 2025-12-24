package utils

import (
	"gin-template/global"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Username             string `json:"username"`
	UserID               int32  `json:"user_id"`
	Role                 string `json:"role"`
	jwt.RegisteredClaims `json:"registered_claims"`
}

func GenerateToken() (string, error) {

	claims := &UserClaims{
		Username: "test",
		UserID:   1,
		Role:     "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.GVA_CONFIG.JWT.Expired) * time.Second)),
		},
	}

	secret := global.GVA_CONFIG.JWT.Secret

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(secret))

	return ss, err
}

func ParseTokenWithClaims(tokenString string, claims *UserClaims) (*jwt.Token, error) {
	secret := global.GVA_CONFIG.JWT.Secret
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
