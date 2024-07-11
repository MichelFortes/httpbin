package handlers

import (
	"net/http"
)

type HealthHandler struct {
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
