package logger

import (
	"fmt"

	"go.uber.org/zap"
)

func createLogger() (zap.SugaredLogger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return zap.SugaredLogger{}, fmt.Errorf("getting full url failed: %w", err)
	}
	defer logger.Sync()
	return *logger.Sugar(), nil
}
