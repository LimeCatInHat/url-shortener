package logger

import (
	"fmt"

	"go.uber.org/zap"
)

func createLogger(level string) (*zap.SugaredLogger, error) {
	logger, err := configureLogger(level)
	if err != nil {
		return &zap.SugaredLogger{}, fmt.Errorf("logger initialization failed: %w", err)
	}
	defer logger.Sync()
	return logger.Sugar(), nil
}

func configureLogger(level string) (*zap.Logger, error) {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, fmt.Errorf("incorrect atomic level %q. error: %w", level, err)
	}

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"}
	cfg.Level = lvl

	zl, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("attempt to build logger config failed: %w", err)
	}
	return zl, nil
}
