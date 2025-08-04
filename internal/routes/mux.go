package routes

import "net/http"

func InitMux() *http.ServeMux {
	mux := http.NewServeMux()

	// hello world endpoint
	mux.HandleFunc("GET /hello", getHello)

	// healthcheck endpoints
	mux.HandleFunc("GET /healthz", getHealthz)
	mux.HandleFunc("GET /livez", getLivez)
	mux.HandleFunc("GET /readyz", getReadyz)

	return mux
}
