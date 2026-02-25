package global

import (
	"gin-template/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_CONFIG *config.Server
	GVA_VIPER  *viper.Viper
	GVA_LOG    *zap.Logger
	GVA_DB     *gorm.DB
)
