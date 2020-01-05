package utils

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware is a middleware that logs requests to stdout
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := &responseLogger{w: w, status: 0, size: 0}
		start := time.Now()
		next.ServeHTTP(logger, r)
		log.Printf("http: %s %s %d %d %v", r.Method, r.URL.Path, logger.status, logger.size, time.Since(start))
	})
}

type responseLogger struct {
	w      http.ResponseWriter
	status int
	size   int
}

func (l *responseLogger) Header() http.Header {
	return l.w.Header()
}

func (l *responseLogger) Write(b []byte) (int, error) {
	size, err := l.w.Write(b)
	l.size += size
	return size, err
}

func (l *responseLogger) WriteHeader(s int) {
	l.w.WriteHeader(s)
	l.status = s
}

func (l *responseLogger) Status() int {
	return l.status
}

func (l *responseLogger) Size() int {
	return l.size
}

func (l *responseLogger) Flush() {
	f, ok := l.w.(http.Flusher)
	if ok {
		f.Flush()
	}
}
