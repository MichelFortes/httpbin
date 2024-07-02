package main

import (
	"net/http/httptest"
	"softplan/httpbin-go/handlers"
	"strings"
	"testing"
)

func Test_DefaultHandler(t *testing.T) {

	req := httptest.NewRequest("GET", "http://localhost:8080?first-arg=first-value", nil)
	rec := httptest.NewRecorder()

	var rootHandler = handlers.DefaultHandler{}
	rootHandler.ServeHTTP(rec, req)
	resp := rec.Result()

	scExpected := 200
	sc := resp.StatusCode
	if sc != scExpected {
		t.Fatalf("Expected %d but returned %d", scExpected, sc)
	}

	ctExpected := "application/json"
	ct := resp.Header.Get("content-type")
	if !strings.HasPrefix(ct, ctExpected) {
		t.Fatalf("Expected content-type %s but got %s", ctExpected, ct)
	}

}

func Test_DefaultHandlerWithCustomResponseCode(t *testing.T) {

	req := httptest.NewRequest("GET", "http://localhost:8080?response_status=429", nil)
	rec := httptest.NewRecorder()

	var rootHandler = handlers.DefaultHandler{}
	rootHandler.ServeHTTP(rec, req)
	resp := rec.Result()

	scExpected := 429
	sc := resp.StatusCode
	if sc != scExpected {
		t.Fatalf("Expected %d but returned %d", scExpected, sc)
	}

}
