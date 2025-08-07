package middleware

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"time"
)

const LevelAccess = slog.LevelError * 2

var LevelNames = map[slog.Leveler]string{
	LevelAccess: "ACCESS",
}

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

func Logger(next http.Handler, w io.Writer) http.Handler {
	opts := slog.HandlerOptions{
		Level: LevelAccess,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.LevelKey:
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}

				a.Value = slog.StringValue(levelLabel)
			case "msg":
				// Remove the message key from the log output
				return slog.Attr{}
			}

			return a
		},
	}
	logger := slog.New(slog.NewTextHandler(w, &opts))
	ctx := context.Background()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w}
		next.ServeHTTP(lrw, r) // Serve the actual request
		duration := time.Since(start)

		logger.Log(ctx, LevelAccess, "", // empty message
			"method", r.Method,
			"path", r.URL.Path,
			"addr", r.RemoteAddr,
			"proto", r.Proto,
			"status", lrw.statusCode,
			"size", lrw.bytesWritten,
			"duration", duration)
	})
}
