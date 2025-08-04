package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	headerWritten bool
	bytesWritten  int
	statusCode    int
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	if !lrw.headerWritten {
		lrw.WriteHeader(http.StatusOK)
	}

	n, err := lrw.ResponseWriter.Write(b)
	lrw.bytesWritten += n

	return n, err
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.ResponseWriter.WriteHeader(code)
	lrw.statusCode = code
	lrw.headerWritten = true
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w}
		next.ServeHTTP(lrw, r) // Serve the actual request
		duration := time.Since(start)

		log.Printf("[%s] %s %s %s %d %d %s\n",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			r.Proto,
			lrw.statusCode,
			lrw.bytesWritten,
			duration,
		)
	})
}
