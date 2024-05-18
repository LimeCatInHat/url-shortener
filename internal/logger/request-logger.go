package logger

import (
	"time"

	"go.uber.org/zap"
)

type RequestLogger struct {
	baseLogger zap.SugaredLogger
}

func (requestLogger RequestLogger) LogRequestInfo(uri string, method string, duration time.Duration) {
	requestLogger.baseLogger.Infoln(
		"uri", uri,
		"method", method,
		"duration", duration,
	)
}

func CreateRequestLogger() (RequestLogger, error) {
	baseLogger, err := createLogger()
	return RequestLogger{baseLogger}, err
}
