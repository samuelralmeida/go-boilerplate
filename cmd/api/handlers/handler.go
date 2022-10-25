package handlers

import (
	"fmt"
	"net/http"

	"github.com/samuelralmeida/go-boilerplate/internal/service"
)

type Handler interface {
	Home(w http.ResponseWriter, r *http.Request)
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

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")

	type resp struct {
		Msg string `json:"msg"`
	}

	payload := resp{
		Msg: msg,
	}

	fmt.Println("handler calling service")
	h.service.Home()

	h.writeJsonResponse(w, payload, http.StatusOK)
}
