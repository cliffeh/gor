package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// usage: go run integration.go /path/to/gor
	// Execute a simple command and check for errors
	cmd := exec.Command(os.Args[1], "-bind", "127.0.0.1:0")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	_, err = stderr.Read(buf)

	if err != nil {
		log.Fatalf("Failed to read command output: %v", err)
	}

	lines := strings.Split(string(buf), "\n")
	pieces := strings.Split(lines[0], "=")
	addr := pieces[len(pieces)-1]

	resp, err := http.Get("http://" + addr + "/hello")
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	responseBody := string(bodyBytes)

	if responseBody != "Hello, World!" {
		log.Fatalf("handler returned unexpected body: got '%v' want '%v'",
			responseBody, "Hello, World!")
	}

	cmd.Process.Signal(os.Interrupt) // Gracefully stop the server
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("Command did not complete successfully: %v", err)
	}
}
