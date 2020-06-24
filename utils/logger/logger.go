package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() (*zap.Logger, *zap.SugaredLogger) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	slogger := logger.Sugar()

	defer logger.Sync()
	defer slogger.Sync()
	return logger, slogger
}
