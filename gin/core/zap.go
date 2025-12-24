package core

import (
	"gin-template/global"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() (logger *zap.Logger, err error) {
	//文件输出
	fileWriteSyncer := getLogWriter()

	encoder := getEncoder()
	//控制台输出
	consoleWriteSyncer := zapcore.AddSync(os.Stdout)

	//level

	var core zapcore.Core

	cores := []zapcore.Core{}

	if global.GVA_CONFIG.Logger.WriteInFile {
		cores = append(cores, zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel))
	}

	if global.GVA_CONFIG.Logger.WriteInConsole {
		cores = append(cores, zapcore.NewCore(encoder, consoleWriteSyncer, zapcore.DebugLevel))
	}

	core = zapcore.NewTee(cores...)

	logger = zap.New(core, zap.AddCaller())
	return logger, nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "log/gin.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
