package handlers

import (
	"fmt"
	"michelfortes/httpbin/internal/constraints"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_ProxyHandler_ShouldReturn400_WhenQueryParamToIsMissing(t *testing.T) {
	req := httptest.NewRequest("GET", "/proxy", nil)
	rec := httptest.NewRecorder()

	handler := ProxyHandler{}
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("Expected status code %d, got %d", http.StatusBadRequest, rec.Code)
	}

	expectedBody := fmt.Sprintf(constraints.TextJsonErrorQueryParamNotFound, constraints.QueryParamTo)
	if !strings.Contains(rec.Body.String(), expectedBody) {
		t.Fatalf("Expected body to contain %q, got %q", expectedBody, rec.Body.String())
	}
}

func Test_ProxyHandler_ShouldReturn500_WhenHttpGetFails(t *testing.T) {
	req := httptest.NewRequest("GET", "/proxy?to=http://invalid-url", nil)
	rec := httptest.NewRecorder()

	handler := ProxyHandler{}
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}

func Test_ProxyHandler_ShouldForwardHeadersToDestination(t *testing.T) {
	destServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom-Header") != "CustomValue" {
			t.Fatalf("Expected header X-Custom-Header to be %q, got %q", "CustomValue", r.Header.Get("X-Custom-Header"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"success"}`))
	}))
	defer destServer.Close()

	req := httptest.NewRequest("GET", fmt.Sprintf("/proxy?to=%s", destServer.URL), nil)
	req.Header.Set("X-Custom-Header", "CustomValue")
	rec := httptest.NewRecorder()

	handler := ProxyHandler{}
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

func Test_ProxyHandler_ShouldHandleNon200ResponseFromDestination(t *testing.T) {
	// Mock destination server
	destServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(`{"error":"bad gateway"}`))
	}))
	defer destServer.Close()

	req := httptest.NewRequest("GET", fmt.Sprintf("/proxy?to=%s", destServer.URL), nil)
	rec := httptest.NewRecorder()

	handler := ProxyHandler{}
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadGateway {
		t.Fatalf("Expected status code %d, got %d", http.StatusBadGateway, rec.Code)
	}

	expectedBody := `{"error":"bad gateway"}`
	if rec.Body.String() != expectedBody {
		t.Fatalf("Expected body %q, got %q", expectedBody, rec.Body.String())
	}
}

func Test_ProxyHandler_ShouldHandleEmptyResponseBodyFromDestination(t *testing.T) {
	destServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer destServer.Close()

	req := httptest.NewRequest("GET", fmt.Sprintf("/proxy?to=%s", destServer.URL), nil)
	rec := httptest.NewRecorder()

	handler := ProxyHandler{}
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("Expected status code %d, got %d", http.StatusNoContent, rec.Code)
	}

	if rec.Body.Len() != 0 {
		t.Fatalf("Expected empty body, got %q", rec.Body.String())
	}
}

func Test_ProxyHandler_ShouldReturn500_WhenDestinationURLIsInvalid(t *testing.T) {
	req := httptest.NewRequest("GET", "/proxy?to=://invalid-url", nil)
	rec := httptest.NewRecorder()

	handler := ProxyHandler{}
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}

// brokenReader simulates a broken response body
type brokenReader struct{}

func (b *brokenReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("broken reader")
}

func (b *brokenReader) Close() error {
	return nil
}
