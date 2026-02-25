package initialize

import (
	"gin-template/global"
	userModels "gin-template/modules/users/models"

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
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&userModels.User{}); 
	if err != nil {
		global.GVA_LOG.Error("failed to migrate users table", zap.Error(err))
	}
}
