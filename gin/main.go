package main

import (
	"gin-template/core"
	"gin-template/global"
)

func initSystem() {
	global.GVA_VIPER = core.Viper()
	global.GVA_LOG, _ = core.InitLogger()
}

func main() {
	initSystem()
	core.RunServer()
}
