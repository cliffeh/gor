package routes

import "net/http"

func InitMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", getHello)

	return mux
}
