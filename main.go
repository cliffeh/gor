package main

import (
	"context"
	"flag"
	"log"
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
			log.Fatalf("Failed to listen on %s: %v", bind, err)
		}
		log.Printf("Server listening on %s", l.Addr().String())

		if err := s.Serve(l); err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Block until a signal is received
	sig := <-quit
	log.Println("Signal received (%v); shutting down server...", sig)

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Release resources associated with the context

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Printf("Server gracefully shut down.")
}
