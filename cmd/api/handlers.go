package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	type resp struct {
		Msg string `json:"message"`
	}

	payload := resp{
		Msg: "ok",
	}

	app.writeJsonResponse(w, payload, http.StatusOK)
}

func (app *application) writeJsonResponse(w http.ResponseWriter, payload interface{}, status int) {
	out, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		log.Println(fmt.Errorf("error to marshal response: %w", err))
		http.Error(w, "error to marshal response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
