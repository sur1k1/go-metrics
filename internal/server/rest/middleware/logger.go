package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type responseData struct {
		status int
		size int
}

type loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
}

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	// записываем ответ, используя оригинальный http.ResponseWriter
	size, err := r.ResponseWriter.Write(b) 
	r.responseData.size += size // захватываем размер
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	// записываем код статуса, используя оригинальный http.ResponseWriter
	r.ResponseWriter.WriteHeader(statusCode) 
	r.responseData.status = statusCode // захватываем код статуса
} 

func Logger(log *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			responseData := &responseData {
				status: 0,
				size: 0,
		}
		lw := loggingResponseWriter {
				ResponseWriter: w,
				responseData: responseData,
		}

			next.ServeHTTP(&lw, r)

			handlingTime := time.Since(start)

			log.Info(
				"got incoming HTTP request",
				zap.String("uri", r.RequestURI),
				zap.String("method", r.Method),
				zap.String("handling time", handlingTime.String()),
				zap.Int("status", lw.responseData.status),
				zap.Int("size", lw.responseData.size),
			)
		}

		return http.HandlerFunc(fn)
	}
}