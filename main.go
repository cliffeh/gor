package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/cliffeh/gor/internal/routes"
)

func main() {
	// TODO make these parameters
	host := "127.0.0.1"
	port := 8080

	mux := routes.InitMux()

	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: mux,
		// Recommended timeouts from
		// https://blog.cloudflare.com/exposing-go-on-the-internet/
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	slog.Info("Server listening", "host", host, "port", port)

	if err := s.ListenAndServe(); err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}
