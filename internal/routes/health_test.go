package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHealthz(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(getHealthz))
	defer ts.Close() // Ensure the server is closed after the test

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
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
	if health.Status != "ok" {
		t.Errorf("Expected health status 'ok', got '%v'", health.Status)
	}
}

func TestGetLivez(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(getLivez))
	defer ts.Close() // Ensure the server is closed after the test

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
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
	if health.Status != "ok" {
		t.Errorf("Expected health status 'ok', got '%v'", health.Status)
	}
}

func TestGetReadyz(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(getReadyz))
	defer ts.Close() // Ensure the server is closed after the test

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
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
	if health.Status != "ok" {
		t.Errorf("Expected health status 'ok', got '%v'", health.Status)
	}
}
