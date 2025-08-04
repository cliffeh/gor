package routes

import "net/http"

func InitMux() *http.ServeMux {
	mux := http.NewServeMux()

	// hello world endpoint
	mux.HandleFunc("GET /hello", getHello)

	// healthcheck endpoints
	mux.HandleFunc("GET /healthz", getHealthz)

	return mux
}
