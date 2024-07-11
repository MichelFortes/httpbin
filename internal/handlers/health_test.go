package handlers

import (
	"fmt"
	"michelfortes/httpbin/internal/constraints"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HealthHandlerOnSuccess(t *testing.T) {

	req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:8080%s", constraints.EndpointHealth), nil)
	rec := httptest.NewRecorder()

	var rootHandler = DefaultHandler{}
	rootHandler.ServeHTTP(rec, req)
	resp := rec.Result()

	scExpected := http.StatusOK
	sc := resp.StatusCode
	if sc != scExpected {
		t.Fatalf("Expected %d but returned %d", scExpected, sc)
	}

}
