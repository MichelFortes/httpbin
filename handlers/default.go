package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type DefaultHandler struct {
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := defaultData{
		Container:   os.Getenv("CONTAINER"),
		RemoteAddr:  r.RemoteAddr,
		Method:      r.Method,
		Path:        r.URL.Path,
		QueryParams: r.URL.Query(),
		Headers:     r.Header,
	}

	if val, err := strconv.Atoi(r.URL.Query().Get("response_status")); err == nil {
		w.WriteHeader(val)
	}

	if slp, err := strconv.Atoi(r.URL.Query().Get("sleep")); err == nil {
		time.Sleep(time.Second * time.Duration(slp))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Default().Fatalln(err)
	}
}

type defaultData struct {
	Container   string              `json:"container"`
	RemoteAddr  string              `json:"remoteAddr"`
	Method      string              `json:"method"`
	Path        string              `json:"path"`
	QueryParams map[string][]string `json:"queryParams"`
	Headers     map[string][]string `json:"headers"`
}
