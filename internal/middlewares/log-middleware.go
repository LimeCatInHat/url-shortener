package middlewares

import (
	"net/http"
	"time"
)

type RequestLogger interface {
	LogRequestInfo(uri string, method string, duration time.Duration)
}

func WithLogging(h http.Handler, log RequestLogger) http.HandlerFunc {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		uri := r.RequestURI
		method := r.Method
		h.ServeHTTP(w, r)

		duration := time.Since(start)

		log.LogRequestInfo(uri, method, duration)
	}
	return http.HandlerFunc(logFn)
}
