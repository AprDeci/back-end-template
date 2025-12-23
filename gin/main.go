package main

import (
	"gin-template/core"
	"gin-template/global"
)

func initSystem() {
	global.GVA_VIPER = core.Viper()
}

func main() {
	initSystem()
	core.RunServer()
}
