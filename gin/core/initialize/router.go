package initialize

import (
	"gin-template/modules/auth"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	Auth auth.AuthRouter
}

func Routers() *gin.Engine {
	Router := gin.New()

	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	routerGroup := RouterGroup{}

	apiGroup := Router.Group("api")

	{
		routerGroup.Auth.InitAuthRouter(apiGroup)
	}

	return Router
}
