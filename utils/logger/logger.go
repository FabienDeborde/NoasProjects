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

// Usage examples:

// logger.Info("failed to fetch URL",
// 	// Structured context as strongly typed Field values.
// 	zap.String("Prefork", os.Getenv("PREFORK")),
// 	zap.Int("attempt", 3),
// 	zap.Duration("backoff", time.Second),
// )

// slogger.Infow("failed to fetch URL",
// 	// Structured context as loosely typed key-value pairs.
// 	"Prefork", os.Getenv("PREFORK"),
// 	"attempt", 3,
// 	"backoff", time.Second,
// )
