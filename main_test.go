package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	// Create a request to the root endpoint
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the request and response recorder
	http.HandlerFunc(RootHandler).ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status code: %d but got: %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "Hello from server"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected response body: %s, but got: %s", expectedBody, rr.Body.String())
	}
}
func TestHelloHandler(t *testing.T) {
	// Create a request to the root endpoint
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the request and response recorder
	http.HandlerFunc(HelloHandler).ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := "this is the hello route running from docker"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected response body %s but got %s", expectedBody, rr.Body.String())
	}
}
