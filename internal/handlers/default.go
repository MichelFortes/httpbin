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

	handleSleepSetting(r)
	handleStatusCodeSetting(r, w)

	w.Header().Set(constraints.ContentTypeKey, constraints.ContentTypeValueJson)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Default().Fatalln(err)
	}
}

func handleSleepSetting(r *http.Request) {
	if slp, err := strconv.Atoi(r.Header.Get(constraints.SettingSleep)); err == nil {
		time.Sleep(time.Second * time.Duration(slp))
	}
}

func handleStatusCodeSetting(r *http.Request, w http.ResponseWriter) {
	if val, err := strconv.Atoi(r.Header.Get(constraints.SettingResponseStatus)); err == nil {
		w.WriteHeader(val)
	}
}
