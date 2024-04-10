package logger

import (
	"log"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func NewLogger() Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to get logger with error: %s", err.Error())
	}

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	return Logger{logger}
}

func (l Logger) StartGinLogger() gin.HandlerFunc {
	return ginzap.Ginzap(l.Logger, time.RFC3339, true)
}
