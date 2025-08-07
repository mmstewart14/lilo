package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs HTTP requests
func LoggingMiddleware(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Create a custom response writer to capture the status code
			rw := &responseWriter{w, http.StatusOK}

			// Call the next handler
			next.ServeHTTP(rw, r)

			// Log the request
			duration := time.Since(start)
			logger.Printf(
				"%s %s %s %d %s",
				r.RemoteAddr,
				r.Method,
				r.URL.Path,
				rw.statusCode,
				duration,
			)
		})
	}
}

// Custom response writer to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
