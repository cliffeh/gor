package main

import (
	"context"
	"flag"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cliffeh/gor/internal/middleware"
	"github.com/cliffeh/gor/internal/routes"
)

func main() {
	bind := ":8080"

	flag.StringVar(&bind, "bind", bind, "interface and port to bind to")

	flag.Parse()

	handler := middleware.Logger(routes.InitMux())

	s := &http.Server{
		Handler: handler,
		// Recommended timeouts from
		// https://blog.cloudflare.com/exposing-go-on-the-internet/
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		l, err := net.Listen("tcp", bind)
		if err != nil {
			slog.Error("Failed to listen", "bind", bind, "error", err)
			os.Exit(1)
		}
		slog.Info("Server listening", "addr", l.Addr().String())

		if err := s.Serve(l); err != http.ErrServerClosed {
			slog.Error("Server failed to start", "error", err)
		}
	}()

	// Block until a signal is received
	sig := <-quit
	slog.Info("Signal received; shutting down server...", "signal", sig)

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Release resources associated with the context

	if err := s.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed", "error", err)
		os.Exit(1)
	}

	slog.Info("Server gracefully shut down.")
}
