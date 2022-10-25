//go:generate mockery --output ../../../internal/mocks --name Handler
package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/samuelralmeida/go-boilerplate/internal/service"
)

type Handler interface {
	homeHandler
}

type handler struct {
	service service.Servicer
}

type Options struct {
	Service service.Servicer
}

func New(options Options) Handler {
	handler := new(handler)
	handler.service = options.Service
	return handler
}

// helpers methods

var (
	errMarshalResponse = errors.New("error to marshal response")
)

func (h *handler) writeJsonResponse(w http.ResponseWriter, payload interface{}, status int) {
	out, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		log.Println(fmt.Errorf("%s: %w", err, errMarshalResponse))
		http.Error(w, errMarshalResponse.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
