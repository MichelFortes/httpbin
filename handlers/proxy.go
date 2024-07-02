package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var logger = log.Default()

type ProxyHandler struct {
}

func (h *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	logger.Println("iniciando proxy")

	dst := r.URL.Query().Get("to")
	if len(dst) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		http.Error(w, "{ \"error\": \"query param 'to' no found\" }", http.StatusBadRequest)
		return
	}

	logger.Printf("enviando requisição para %s \n", dst)

	resp, err := http.Get(dst)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if e := json.NewEncoder(w).Encode(err); e != nil {
			w.Write([]byte(e.Error()))
		}
		logger.Panicln(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if e := json.NewEncoder(w).Encode(err); e != nil {
			w.Write([]byte(e.Error()))
		}
		logger.Panicln(err)
		return
	}

	logger.Printf("got response with status: %d and content-type: %s\n", resp.StatusCode, resp.Header.Get("Content-Type"))

	for k, v := range resp.Header {
		for _, s := range v {
			w.Header().Add(k, s)
		}
	}

	w.WriteHeader(resp.StatusCode)

	if len(body) > 0 {
		w.Write(body)
	}

}
