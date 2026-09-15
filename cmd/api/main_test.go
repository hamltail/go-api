package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	handler := newHandler(nil, "test-api-key")

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	if response.Body.String() != "OK" {
		t.Fatalf("expected body OK, got %q", response.Body.String())
	}
}

func TestPostsRequiresAPIKey(t *testing.T) {
	handler := newHandler(nil, "test-api-key")

	request := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", response.Code)
	}
}

func TestPostsWithAPIKey(t *testing.T) {
	handler := newHandler(nil, "test-api-key")

	request := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)
	request.Header.Set("X-API-Key", "test-api-key")

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}
