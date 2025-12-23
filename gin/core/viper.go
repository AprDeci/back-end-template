package core

import (
	"fmt"

	"gin-template/core/internal"
	"gin-template/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	config := getConfigPath()

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	return v
}

func getConfigPath() (config string) {

	switch gin.Mode() { // 根据 gin 模式文件名
	case gin.DebugMode:
		config = internal.ConfigDevFile
	case gin.ReleaseMode:
		config = internal.ConfigProdFile
	case gin.TestMode:
		config = internal.ConfigTestFile
	}
	fmt.Printf("您正在使用 gin 的 %s 模式运行, config 的路径为 %s\n", gin.Mode(), config)

	_, err := os.Stat(config)
	if err != nil || os.IsNotExist(err) {
		config = internal.ConfigFile
		fmt.Printf("配置文件路径不存在, 使用默认配置文件路径: %s\n", config)
	}

	return
}
