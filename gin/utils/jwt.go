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

// GenerateTokenWithUserInfo 根据用户信息生成JWT令牌
func GenerateTokenWithUserInfo(userID int32, username string, role string) (string, error) {

	claims := &UserClaims{
		Username: username,
		UserID:   userID,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.GVA_CONFIG.JWT.Expired) * time.Second)),
		},
	}

	secret := global.GVA_CONFIG.JWT.Secret

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(secret))

	return ss, err
}

// GenerateToken 生成JWT令牌(保留旧函数用于兼容性)
func GenerateToken() (string, error) {
	return GenerateTokenWithUserInfo(1, "test", "admin")
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
