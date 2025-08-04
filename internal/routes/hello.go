package routes

import (
	"fmt"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	name := "World"
	if r.URL.Query().Has("name") {
		name = r.URL.Query().Get("name")
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}
