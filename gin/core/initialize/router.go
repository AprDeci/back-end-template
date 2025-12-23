package initialize

import (
	"gin-template/modules/auth"

	"github.com/gin-gonic/gin"

	openapiui "github.com/PeterTakahashi/gin-openapi/openapiui"
)

type RouterGroup struct {
	Auth auth.AuthRouter
}

func Routers() *gin.Engine {
	Router := gin.New()
	routerGroup := RouterGroup{}

	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	Router.GET("/docs/*any", openapiui.WrapHandler(openapiui.Config{
		SpecURL:      "/docs/openapi.json",
		SpecFilePath: "./docs/swagger.json",
		Title:        "Example API",
		Theme:        "light", // or "dark"
	}))

	apiGroup := Router.Group("api")

	{
		routerGroup.Auth.InitAuthRouter(apiGroup)
	}

	return Router
}
