package middleware

import (
	"gin-template/core/response"
	"gin-template/utils"

	"github.com/gin-gonic/gin"
)

const (
	GinHeaderAuthorizationKey = "Authorization"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req         = c.Request
			tokenString = req.Header.Get(GinHeaderAuthorizationKey)
		)

		if tokenString == "" {
			response.Fail(c, response.Unauthorized, "未授权访问")
			c.Abort()
			return
		}

		claims := &utils.UserClaims{}
		token, err := utils.ParseTokenWithClaims(tokenString, claims)
		if err != nil {
			response.Fail(c, response.Unauthorized, "未授权访问")
			c.Abort()
			return
		}

		if !token.Valid {
			response.Fail(c, response.Unauthorized, "Token无效")
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}

}
