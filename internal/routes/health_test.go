package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var healthTests = []struct {
	path       string
	health     Health
	statusCode int
}{
	{path: "/healthz", health: Health{Status: "ok"}, statusCode: 200},
	{path: "/livez", health: Health{Status: "ok"}, statusCode: 200},
	{path: "/readyz", health: Health{Status: "ok"}, statusCode: 200},
}

func TestHealthRoutes(t *testing.T) {
	ts := httptest.NewServer(InitMux())
	defer ts.Close() // Ensure the server is closed after the tests run

	for _, tt := range healthTests {
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

			var health Health
			err = json.Unmarshal(bodyBytes, &health)
			if err != nil {
				t.Fatalf("Failed to unmarshal response body: %v", err)
			}
			if health.Status != tt.health.Status {
				t.Errorf("Expected health status '%s', got '%v'", tt.health.Status, health.Status)
			}
		})
	}
}
