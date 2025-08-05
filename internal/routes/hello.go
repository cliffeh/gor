package routes

import (
	"fmt"
	"net/http"
)

const defaultName string = "World"

func getHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = defaultName
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}
