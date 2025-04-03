package handlers

import (
	"bytes"
	"encoding/json"
	"michelfortes/httpbin/internal/constraints"
	"michelfortes/httpbin/pkg/model"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestDefaultHandler_ServeHTTP(t *testing.T) {
	os.Setenv(constraints.EnvServiceId, "test-service-id")
	defer os.Unsetenv(constraints.EnvServiceId)

	handler := &DefaultHandler{}

	t.Run("basic request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test-path", nil)
		req.RemoteAddr = "127.0.0.1:12345"
		req.Header.Set("Test-Header", "HeaderValue")

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rec.Code)
		}

		var response model.ResponseBody
		if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if response.ServiceId != "test-service-id" {
			t.Errorf("expected ServiceId 'test-service-id', got '%s'", response.ServiceId)
		}
		if response.Path != "/test-path" {
			t.Errorf("expected Path '/test-path', got '%s'", response.Path)
		}
		if response.Headers["Test-Header"][0] != "HeaderValue" {
			t.Errorf("expected header 'Test-Header' to be 'HeaderValue', got '%s'", response.Headers["Test-Header"][0])
		}
	})

	t.Run("custom status code", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(constraints.HeaderSettingResponseStatus, "418")

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusTeapot {
			t.Errorf("expected status 418, got %d", rec.Code)
		}
	})

	t.Run("sleep behavior", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(constraints.HeaderSettingSleep, "1")

		start := time.Now()
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		duration := time.Since(start)

		if duration < time.Second {
			t.Errorf("expected sleep of at least 1 second, got %v", duration)
		}
	})

	t.Run("unsupported media type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("test payload"))
		req.Header.Set(constraints.HeaderSettingContentType, "application/json")
		req.Header.Set("Content-Type", "text/plain") // Fixed header key

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusUnsupportedMediaType {
			t.Errorf("expected status 415, got %d", rec.Code)
		}
	})
}
