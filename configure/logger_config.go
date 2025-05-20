package configure

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger() *zap.Logger {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	return logger
}
