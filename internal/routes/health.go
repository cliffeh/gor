package routes

import (
	"encoding/json"
	"net/http"
)

// TODO have this come from somewhere else
var health Health = Health{Status: "ok"}

// TODO this should probably be defined somewhere else?
type Health struct {
	Status string `json:"status"`
}

func getHealthz(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(health)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func getLivez(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(health)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func getReadyz(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(health)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
