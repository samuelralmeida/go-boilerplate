package handlers

import (
	"fmt"
	"net/http"
)

type homeHandler interface {
	Home(w http.ResponseWriter, r *http.Request)
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
