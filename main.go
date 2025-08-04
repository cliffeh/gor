package main

import (
	"flag"
	"log/slog"
	"net/http"
	"time"

	"github.com/cliffeh/gor/internal/middleware"
	"github.com/cliffeh/gor/internal/routes"
)

func main() {
	bind := "0.0.0.0:8080"

	flag.StringVar(&bind, "bind", bind, "interface and port to bind to")

	flag.Parse()

	handler := middleware.Logger(routes.InitMux())
	// mux := routes.InitMux()

	s := &http.Server{
		Addr:    bind,
		Handler: handler,
		// Recommended timeouts from
		// https://blog.cloudflare.com/exposing-go-on-the-internet/
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	slog.Info("Server listening", "bind", bind)

	if err := s.ListenAndServe(); err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}
