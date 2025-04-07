package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"michelfortes/httpbin/internal/constraints"
	"net/http"
)

var logger = log.Default()

type ProxyHandler struct {
}

func (h *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	logger.Println(constraints.TextProxingRequest)

	dst := r.URL.Query().Get(constraints.QueryParamTo)
	if len(dst) == 0 {
		w.Header().Set(constraints.HeaderContentType, constraints.ContentTypeApplicationJson)
		http.Error(w, fmt.Sprintf(constraints.TextJsonErrorQueryParamNotFound, constraints.QueryParamTo), http.StatusBadRequest)
		return
	}

	logger.Printf(constraints.TextSendingReqTo, dst)

	req, err := http.NewRequest(r.Method, dst, r.Body)
	if err != nil {
		logger.Println("Error creating request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	for k, v := range r.Header {
		for _, s := range v {
			req.Header.Add(k, s)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set(constraints.HeaderContentType, constraints.ContentTypeApplicationJson)
		if e := json.NewEncoder(w).Encode(err); e != nil {
			w.Write([]byte(e.Error()))
		}
		logger.Panicln(err)
		return
	}

	logger.Printf(constraints.TextGotResponse, resp.StatusCode, resp.Header.Get(constraints.HeaderContentType))

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
