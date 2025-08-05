package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInitMux(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf) // Redirect log output

	// Create a mock handler
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(http.StatusOK) // This is handled by HttpHandler
		w.Write([]byte("OK"))
	})

	// Wrap the mock handler with the logging middleware
	handlerToTest := Logger(mockHandler)

	ts := httptest.NewServer(handlerToTest)
	defer ts.Close() // Ensure the server is closed after the test

	// Create a mock request
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	// Serve the request
	handlerToTest.ServeHTTP(rr, req)

	// Assert the captured logs
	fields := strings.Split(buf.String(), " ")
	// 2025/08/05 08:26:38 [192.0.2.1:1234] GET /test HTTP/1.1 200 2 1.408µs
	if len(fields) < 9 {
		t.Fatalf("Expected log format to contain at least 9 fields, got %d", len(fields))
	}
	if fields[3] != "GET" {
		t.Errorf("Expected method to be 'GET', got '%s'", fields[3])
	}
	if fields[4] != "/test" {
		t.Errorf("Expected path to be '/test', got '%s'", fields[4])
	}
	if fields[6] != "200" {
		t.Errorf("Expected status code to 200, got '%s'", fields[6])
	}
	if fields[7] != "2" {
		t.Errorf("Expected content length to be 2, got '%s'", fields[7])
	}
}
