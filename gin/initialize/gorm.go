package initialize

import (
	"gin-template/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	dsn := global.GVA_CONFIG.DB.DSN()

	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
		SkipInitializeWithVersion: false,
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		global.GVA_LOG.Error("failed to connect database", zap.Error(err))
	}
	return db
}
