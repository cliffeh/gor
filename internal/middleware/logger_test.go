package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInitMux(t *testing.T) {
	var buf bytes.Buffer

	// Create a mock handler
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(http.StatusOK) // This is handled by HttpHandler
		w.Write([]byte("OK"))
	})

	// Wrap the mock handler with the logging middleware
	handlerToTest := Logger(mockHandler, &buf)

	ts := httptest.NewServer(handlerToTest)
	defer ts.Close() // Ensure the server is closed after the test

	// Create a mock request
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	// Serve the request
	handlerToTest.ServeHTTP(rr, req)

	// Assert the captured logs
	fields := strings.Split(buf.String(), " ")
	// time=2025-08-06T22:37:31.811-04:00 level=ACCESS method=GET path=/hello
	// addr=127.0.0.1:34352 proto=HTTP/1.1 status=200 size=13 duration=22.499µs
	if len(fields) < 9 {
		t.Fatalf("Expected log to contain at least 9 fields, got %d (%v)", len(fields), fields)
	}

	for _, field := range fields {
		pieces := strings.Split(field, "=")
		switch pieces[0] {
		case "level":
			if pieces[1] != "ACCESS" {
				t.Errorf("Expected log level ACCESS, got %s", pieces[1])
			}
		case "method":
			if pieces[1] != "GET" {
				t.Errorf("Expected method GET, got %s", pieces[1])
			}
		case "path":
			if pieces[1] != "/test" {
				t.Errorf("Expected path /test, got %s", pieces[1])
			}
		case "status":
			if pieces[1] != "200" {
				t.Errorf("Expected status 200, got %s", pieces[1])
			}
		case "size":
			if pieces[1] != "2" {
				t.Errorf("Expected size 2, got %s", pieces[1])
			}
		}
	}
}
