package configure

import (
	"go.uber.org/zap"
)

type LoggerConfigure struct {
}

func (l *LoggerConfigure) InitLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return logger
}
