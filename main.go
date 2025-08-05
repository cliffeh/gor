package main

import (
	"flag"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/cliffeh/gor/internal/middleware"
	"github.com/cliffeh/gor/internal/routes"
)

var listener net.Listener

func runServer() error {
	handler := middleware.Logger(routes.InitMux())

	s := &http.Server{
		Handler: handler,
		// Recommended timeouts from
		// https://blog.cloudflare.com/exposing-go-on-the-internet/
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return s.Serve(listener)
}

func main() {
	bind := ":8080"

	flag.StringVar(&bind, "bind", bind, "interface and port to bind to")

	flag.Parse()

	var err error
	listener, err = net.Listen("tcp", bind)
	if err != nil {
		slog.Error("Failed to listen", "error", err)
		os.Exit(1)
	}

	slog.Info("Listening on", "addr", listener.Addr())

	err = runServer()
	if err != nil {
		slog.Error("Failed to run server", "error", err)
		os.Exit(1)
	}
	slog.Info("Server stopped")
}
