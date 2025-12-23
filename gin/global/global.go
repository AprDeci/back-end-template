package global

import (
	"gin-template/config"

	"github.com/spf13/viper"
)

var (
	GVA_CONFIG *config.Server
	GVA_VIPER  *viper.Viper
)
