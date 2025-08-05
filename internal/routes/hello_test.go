package routes

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var helloTests = []struct {
	path       string
	body       string
	statusCode int
}{
	{path: "/hello", body: "Hello, World!", statusCode: 200},
	{path: "/hello?name=Gor", body: "Hello, Gor!", statusCode: 200},
}

func TestHelloRoutes(t *testing.T) {
	ts := httptest.NewServer(InitMux())
	defer ts.Close() // Ensure the server is closed after the tests run

	for _, tt := range helloTests {
		t.Run(fmt.Sprintf("GET %s", tt.path), func(t *testing.T) {
			resp, err := http.Get(ts.URL + tt.path)
			if err != nil {
				t.Fatalf("Failed to make GET request: %v", err)
			}

			if resp.StatusCode != tt.statusCode {
				t.Errorf("Expected status code %d, got %d", tt.statusCode, resp.StatusCode)
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}
			responseBody := string(bodyBytes)

			if responseBody != tt.body {
				t.Errorf("handler returned unexpected body: got '%v' want '%v'",
					responseBody, tt.body)
			}
		})
	}
}
