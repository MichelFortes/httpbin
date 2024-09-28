package handlers

import (
	"michelfortes/httpbin/internal/constraints"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_ShouldReturnStatusCode200AndContentTypeAppJson_OnSimpleGetRequest(t *testing.T) {

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

	ctExpected := constraints.ContentTypeValueJson
	ct := resp.Header.Get(constraints.ContentTypeKey)
	if !strings.HasPrefix(ct, ctExpected) {
		t.Fatalf("Expected content-type %s but got %s", ctExpected, ct)
	}

}

func Test_ShouldReturnAppropriateStatus_WhenSet(t *testing.T) {

	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	req.Header.Add(constraints.SettingResponseStatus, "429")
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
