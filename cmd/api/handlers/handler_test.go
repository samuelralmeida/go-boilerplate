package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handler_writeJsonResponse(t *testing.T) {
	h := &handler{}

	payload := "success"

	type args struct {
		w       http.ResponseWriter
		payload interface{}
		status  int
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantMsg    string
	}{
		{
			name:       "erro to marshal payload",
			args:       args{w: httptest.NewRecorder(), payload: make(chan int), status: 200},
			wantStatus: 500,
			wantMsg:    errMarshalResponse.Error(),
		},
		{
			name:       "success to write json response",
			args:       args{w: httptest.NewRecorder(), payload: payload, status: 200},
			wantStatus: 200,
			wantMsg:    payload,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h.writeJsonResponse(tt.args.w, tt.args.payload, tt.args.status)
			recorder := tt.args.w.(*httptest.ResponseRecorder)

			assert.Equal(t, tt.wantStatus, recorder.Code)
			assert.Contains(t, fmt.Sprint(recorder.Body), tt.wantMsg)

		})
	}
}
