package initializers

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() *zap.Logger {
	Logger := &lumberjack.Logger{
		Filename:   os.Getenv("LOGGER_PATH"),
		MaxSize:    50,
		MaxBackups: 7,
		MaxAge:     30,
		Compress:   true,
	}

	fileSyncer := zapcore.AddSync(Logger)
	fileEncoderCfg := zap.NewProductionEncoderConfig()
	fileEncoderCfg.TimeKey = "time"
	fileEncoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(fileEncoderCfg),
		fileSyncer,
		zap.InfoLevel,
	)

	consoleSyncer := zapcore.AddSync(os.Stdout)

	consoleEncoderCfg := zap.NewDevelopmentEncoderConfig()
	consoleEncoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleEncoderCfg),
		consoleSyncer,
		zap.DebugLevel,
	)

	core := zapcore.NewTee(
		fileCore,
		consoleCore,
	)

	return zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}

func LoggerFields() []string {
	return []string{
		"method",
		"pid",
		"route",
		"url",
		"reqHeaders",
		"requestId",
		"body",
		"queryParams",
		"status",
		"latency",
		"host",
		"ips",
		"ua",
		"resBody",
		"error",
	}
}
