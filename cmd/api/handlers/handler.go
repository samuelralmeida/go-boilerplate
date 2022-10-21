package handlers

import (
	"net/http"
)

type Handler interface {
	Home(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	// service service.Servicer
}

type Options struct {
	// Service service.Servicer
}

func New(options Options) Handler {
	handler := new(handler)
	// handler.service = options.Service
	return handler
}

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")

	type resp struct {
		Msg string `json:"msg"`
	}

	payload := resp{
		Msg: msg,
	}

	h.writeJsonResponse(w, payload, http.StatusOK)
}
