package logger

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type (
	RequestLogger struct {
		logHandler *logHandler
		baseLogger *zap.SugaredLogger
	}

	logHandler struct {
		http.ResponseWriter
		responseData *responseData
		request      *http.Request
	}

	responseData struct {
		status int
		size   int
	}
)

func (requestLogger *RequestLogger) HandleRequest(request *http.Request, w http.ResponseWriter) {
	h := &logHandler{
		ResponseWriter: w,
		request:        request,
		responseData: &responseData{
			status: 0,
			size:   0,
		},
	}
	requestLogger.logHandler = h
}

func (requestLogger *RequestLogger) GetLoggingResponseWriter() http.ResponseWriter {
	return requestLogger.logHandler
}

func (requestLogger *RequestLogger) LogRequestInfo(message string, duration time.Duration) {
	request := requestLogger.logHandler.request

	requestLogger.baseLogger.Infow(
		message,
		"uri", request.RequestURI,
		"method", request.Method,
		"status", requestLogger.logHandler.responseData.status,
		"duration", duration,
		"size", requestLogger.logHandler.responseData.size,
	)
}

func (r *logHandler) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *logHandler) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func CreateRequestLogger() (RequestLogger, error) {
	baseLogger, err := createLogger("info")
	return RequestLogger{baseLogger: baseLogger}, err
}
