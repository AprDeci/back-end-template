package core

import (
	"gin-template/config"
	"gin-template/global"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() (logger *zap.Logger, err error) {
	//文件输出
	fileWriteSyncer := getLogWriter(global.GVA_CONFIG.Logger.LogFile)

	encoder := getEncoder()
	//控制台输出
	consoleWriteSyncer := zapcore.AddSync(os.Stdout)

	//level
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(global.GVA_CONFIG.Logger.Level)); err != nil {
		level = zapcore.InfoLevel
	}

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
	zap.ReplaceGlobals(logger)
	return logger, nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logFile config.LogFile) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile.FilePath,
		MaxSize:    logFile.MaxSize,
		MaxBackups: logFile.MaxBackup,
		MaxAge:     logFile.MaxAge,
		Compress:   logFile.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
