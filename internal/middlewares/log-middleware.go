package middlewares

import (
	"net/http"
	"time"

	"github.com/LimeCatInHat/url-shortener/internal/logger"
)

func WithLogging(h http.Handler, log logger.RequestLogger) http.HandlerFunc {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.HandleRequest(r, w)

		h.ServeHTTP(log.GetLoggingResponseWriter(), r)

		duration := time.Since(start)

		log.LogRequestInfo(duration)
	}
	return http.HandlerFunc(logFn)
}
