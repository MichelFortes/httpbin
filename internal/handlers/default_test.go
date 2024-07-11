package handlers

import (
	"fmt"
	"michelfortes/httpbin/internal/constraints"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_DefaultHandler(t *testing.T) {

	req := httptest.NewRequest("GET", "http://localhost:8080?first-arg=first-value", nil)
	rec := httptest.NewRecorder()

	var rootHandler = DefaultHandler{}
	rootHandler.ServeHTTP(rec, req)
	resp := rec.Result()

	scExpected := 200
	sc := resp.StatusCode
	if sc != scExpected {
		t.Fatalf("Expected %d but returned %d", scExpected, sc)
	}

	ctExpected := constraints.HeaderContentTypeValueJson
	ct := resp.Header.Get(constraints.HeaderContentTypeKey)
	if !strings.HasPrefix(ct, ctExpected) {
		t.Fatalf("Expected content-type %s but got %s", ctExpected, ct)
	}

}

func Test_DefaultHandlerWithCustomResponseCode(t *testing.T) {

	req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:8080?%s=429", constraints.QueryParamResponseStatus), nil)
	rec := httptest.NewRecorder()

	var rootHandler = DefaultHandler{}
	rootHandler.ServeHTTP(rec, req)
	resp := rec.Result()

	scExpected := 429
	sc := resp.StatusCode
	if sc != scExpected {
		t.Fatalf("Expected %d but returned %d", scExpected, sc)
	}

}
