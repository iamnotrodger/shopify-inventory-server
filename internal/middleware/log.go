package middleware

import (
	"log"
	"net/http"
	"time"
)

type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{ResponseWriter: w, statusCode: 200}
}

func (w *LogResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *LogResponseWriter) Write(body []byte) (int, error) {
	return w.ResponseWriter.Write(body)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		logResponseWriter := NewLogResponseWriter(w)
		next.ServeHTTP(logResponseWriter, r)

		log.Println(r.Method, r.RequestURI, logResponseWriter.statusCode, time.Since(startTime).String())
	})
}
