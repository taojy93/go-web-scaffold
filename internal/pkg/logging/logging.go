package logging

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	Logger, err = config.Build()
	if err != nil {
		log.Fatalf("Error initializing zap logger: %v", err)
	}
}

func SyncLogger() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}
