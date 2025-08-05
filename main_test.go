package main

import (
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestRunServer(t *testing.T) {
	// var err error
	// listener, err = net.Listen("tcp", "127.0.0.1:0")
	// if err != nil {
	// 	t.Fatalf("Failed to create a local listener: %v", err)
	// }
	// defer listener.Close() // Ensure the listener is closed after the test

	os.Args = []string{"gor", "-bind", "127.0.0.1:0"}

	go main()

	for listener == nil {
		time.Sleep(10 * time.Millisecond) // Give the server a moment to start
	}

	req, err := http.NewRequest("GET", "http://"+listener.Addr().String()+"/hello", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	responseBody := string(bodyBytes)

	if responseBody != "Hello, World!" {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'",
			responseBody, "Hello, World!")
	}
}
