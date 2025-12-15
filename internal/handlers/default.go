package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"michelfortes/httpbin/internal/constraints"
	"michelfortes/httpbin/pkg/model"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type DefaultHandler struct {
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse sleep setting first to determine timeout
	timeout := 5 * time.Second
	if slpRaw := r.Header.Get(constraints.HeaderSettingSleep); slpRaw != "" {
		if slp, err := strconv.Atoi(slpRaw); err == nil {
			timeout = time.Duration(slp)*time.Second + 5*time.Second
		}
	}

	// Criar um contexto com timeout
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	// Substituir o contexto da requisição pelo novo contexto
	r = r.WithContext(ctx)

	result := model.ResponseBody{
		ServiceId:   os.Getenv(constraints.EnvServiceId),
		RemoteAddr:  r.RemoteAddr,
		Method:      r.Method,
		Path:        r.URL.Path,
		QueryParams: r.URL.Query(),
		Headers:     r.Header,
	}

	if r.Body != nil {
		defer r.Body.Close()
		bodyBytes, err := io.ReadAll(r.Body)
		if err == nil {
			result.Payload = string(bodyBytes)
		}
	}

	select {
	case <-ctx.Done():
		http.Error(w, "Request timed out", http.StatusGatewayTimeout)
		return
	default:
	}

	handleSleepSetting(ctx, r)
	handleStatusCodeSetting(r, w)

	if settingContentType := r.Header.Get(constraints.HeaderSettingContentType); len(settingContentType) > 0 {

		if clientContentType := r.Header.Get(constraints.HeaderContentType); !strings.EqualFold(clientContentType, settingContentType) {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
	}

	w.Header().Set(constraints.HeaderContentType, constraints.ContentTypeApplicationJson)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
}

func handleSleepSetting(ctx context.Context, r *http.Request) {
	if slp, err := strconv.Atoi(r.Header.Get(constraints.HeaderSettingSleep)); err == nil {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second * time.Duration(slp)):
		}
	}
}

func handleStatusCodeSetting(r *http.Request, w http.ResponseWriter) {
	if status, err := strconv.Atoi(r.Header.Get(constraints.HeaderSettingResponseStatus)); err == nil {
		w.WriteHeader(status)
	}
}
