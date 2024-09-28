package handlers

import (
	"bytes"
	"fmt"
	"michelfortes/httpbin/internal/constraints"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func Test_ShouldReturn200AndContentTypeJson_WhenNothingIsSet(t *testing.T) {

	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	var rootHandler = DefaultHandler{}
	rootHandler.ServeHTTP(rec, req)
	resp := rec.Result()

	scExpected := 200
	sc := resp.StatusCode
	if sc != scExpected {
		t.Fatalf("Expected %d but returned %d", scExpected, sc)
	}

	ctExpected := constraints.ContentTypeAppJsonUtf8
	ct := resp.Header.Get(constraints.HeaderContentType)
	if !strings.HasPrefix(ct, ctExpected) {
		t.Fatalf("Expected content-type %s but got %s", ctExpected, ct)
	}

}

func Test_ShouldReturn429_WhenSetStatus429IsSet(t *testing.T) {

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

func Test_ShouldHaveLatency2_WhenSleep2(t *testing.T) {

	sleep := 2.0

	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	req.Header.Add(constraints.SettingSleep, fmt.Sprintf("%0.0f", sleep))
	rec := httptest.NewRecorder()

	var rootHandler = DefaultHandler{}
	startTime := time.Now()
	rootHandler.ServeHTTP(rec, req)

	dur := time.Since(startTime)
	if dur.Seconds() < sleep {
		t.Fatalf("Expected %f but returned %f", sleep, dur.Seconds())
	}

}

func Test_ShouldReturn200_WhenSettingContentTypeIsEqualsToSendByClient(t *testing.T) {

	ctSetting := "application/json"

	req := httptest.NewRequest("POST", "http://localhost:8080", bytes.NewBufferString("{}"))
	req.Header.Add(constraints.SettingContentType, ctSetting)
	req.Header.Add(constraints.HeaderContentType, ctSetting)
	rec := httptest.NewRecorder()

	var rootHandler = DefaultHandler{}
	rootHandler.ServeHTTP(rec, req)
	resp := rec.Result()

	scExpected := 200
	sc := resp.StatusCode
	if sc != scExpected {
		t.Fatalf("Expected %d but returned %d", scExpected, sc)
	}

}

func Test_ShouldReturn415_WhenSettingContentTypeIsNotEqualsToSendByClient(t *testing.T) {

	ctSetting := "application/json"
	ctClient := "text/html"

	req := httptest.NewRequest("POST", "http://localhost:8080", bytes.NewBufferString("{}"))
	req.Header.Add(constraints.SettingContentType, ctSetting)
	req.Header.Add(constraints.HeaderContentType, ctClient)
	rec := httptest.NewRecorder()

	var rootHandler = DefaultHandler{}
	rootHandler.ServeHTTP(rec, req)
	resp := rec.Result()

	scExpected := 415
	sc := resp.StatusCode
	if sc != scExpected {
		t.Fatalf("Expected %d but returned %d", scExpected, sc)
	}

}
