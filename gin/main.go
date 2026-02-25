package main

import (
	"gin-template/core"
	"gin-template/global"

	"go.uber.org/zap"
)

func initSystem() {
	global.GVA_VIPER = core.Viper()
	global.GVA_LOG, _ = core.InitLogger()
}

func main() {
	initSystem()

	global.GVA_LOG.Info("后端地址", zap.String("address", "http://127.0.0.1:8080"))
	global.GVA_LOG.Info("文档地址", zap.String("address", "http://127.0.0.1:8080/docs"))

	core.RunServer()
}
