package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/cliffeh/gor/internal/routes"
)

func doRequest(url string, expectedStatusCode int) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make GET request to %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatusCode {
		log.Fatalf("Expected status code %d, got %d",
			expectedStatusCode, resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body for %s: %v", url, err)
	}
	return bodyBytes
}

func main() {
	// usage: go run integration.go /path/to/gor
	// Execute a simple command and check for errors
	cmd := exec.Command(os.Args[1], "-bind", "127.0.0.1:0")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	_, err = stdout.Read(buf)

	if err != nil {
		log.Fatalf("Failed to read command output: %v", err)
	}

	lines := strings.Split(string(buf), "\n")
	pieces := strings.Split(lines[0], "=")
	addr := pieces[len(pieces)-1]

	responseBody := string(doRequest("http://"+addr+"/hello", 200))

	if responseBody != "Hello, World!" {
		log.Fatalf("handler returned unexpected body: got '%v' want '%v'",
			responseBody, "Hello, World!")
	}

	var health routes.Health
	expectedHealth := routes.Health{Status: "ok"}

	for _, path := range []string{"/healthz", "/livez", "/readyz"} {
		responseBody := doRequest("http://"+addr+path, 200)
		err = json.Unmarshal(responseBody, &health)
		if err != nil {
			log.Fatalf("Failed to unmarshal response body for %s: %v", path, err)
		}
		if health != expectedHealth {
			log.Fatalf("Expected health '%s', got '%v'", expectedHealth, health)
		}
	}

	cmd.Process.Signal(os.Interrupt) // Gracefully stop the server
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("Command did not complete successfully: %v", err)
	}
}
