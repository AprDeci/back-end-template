package initialize

import (
	"gin-template/global"
	"gin-template/middleware"
	"gin-template/modules/auth"
	"time"

	ginzap "github.com/gin-contrib/zap"
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
		Title:        "Gin Template",
		Theme:        "light", // or "dark"
	}))

	logger := global.GVA_LOG

	Router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	Router.Use(ginzap.RecoveryWithZap(logger, true))

	Router.Use(middleware.JWTMiddleware())

	apiGroup := Router.Group("api")

	{
		routerGroup.Auth.InitAuthRouter(apiGroup)
	}

	return Router
}
