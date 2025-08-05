package routes

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	msg := hello("")
	want := "Hello, World!"
	if msg != want {
		t.Errorf(`hello("") = %s, want: %s`, msg, want)
	}
}

func TestHelloName(t *testing.T) {
	msg := hello("Gor")
	want := "Hello, Gor!"
	if msg != want {
		t.Errorf(`hello("Gor") = %s, want: %s`, msg, want)
	}
}

func TestGetHello(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(getHello))
	defer ts.Close() // Ensure the server is closed after the test

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	responseBody := string(bodyBytes)

	if responseBody != "Hello, World!" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, "Hello, World!")
	}
}

func TestGetHelloName(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(getHello))
	defer ts.Close() // Ensure the server is closed after the test

	resp, err := http.Get(ts.URL + "?name=Gor")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	responseBody := string(bodyBytes)

	if responseBody != "Hello, Gor!" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, "Hello, Gor!")
	}
}
