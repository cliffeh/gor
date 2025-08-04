package routes

import (
	"fmt"
	"net/http"
)

const defaultName string = "World"

func hello(name string) string {
	if name == "" {
		name = defaultName
	}
	return fmt.Sprintf("Hello, %s!", name)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	message := hello(r.URL.Query().Get("name"))

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(message))
}
