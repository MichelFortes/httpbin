package handlers

import (
	"encoding/json"
	"log"
	"michelfortes/httpbin/internal/constraints"
	"michelfortes/httpbin/pkg/model"
	"net/http"
	"os"
	"strconv"
	"time"
)

type DefaultHandler struct {
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	result := model.ResponseBody{
		ServiceId:   os.Getenv(constraints.EnvServiceId),
		RemoteAddr:  r.RemoteAddr,
		Method:      r.Method,
		Path:        r.URL.Path,
		QueryParams: r.URL.Query(),
		Headers:     r.Header,
	}

	if val, err := strconv.Atoi(r.URL.Query().Get(constraints.QueryParamResponseStatus)); err == nil {
		w.WriteHeader(val)
	}

	if slp, err := strconv.Atoi(r.URL.Query().Get(constraints.QueryParamSleep)); err == nil {
		time.Sleep(time.Second * time.Duration(slp))
	}

	w.Header().Set(constraints.HeaderContentTypeKey, constraints.HeaderContentTypeValueJson)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Default().Fatalln(err)
	}
}
