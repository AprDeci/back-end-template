package global

import (
	"gin-template/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	GVA_CONFIG *config.Server
	GVA_VIPER  *viper.Viper
	GVA_LOG    *zap.Logger
)
