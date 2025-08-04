package routes

import (
	"net/http"
)

func getHealthz(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// TODO some kind of healthcheck struct
	w.Write([]byte(`{"status": "ok"}`))
}

// TODO livez and readyz endpoints
