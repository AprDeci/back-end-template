package auth

import (
	"gin-template/modules/auth/api"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (ar *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("auth")

	{
		authRouter.POST("login", api.Login)
	}

}
