package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samuelralmeida/go-boilerplate/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_handler_Home(t *testing.T) {
	mockService := mocks.NewServicer(t)

	mockService.On("Home").Return(nil).Twice()

	h := &handler{service: mockService}

	tests := []struct {
		name string
		msg  string
	}{
		{
			name: "should get msg correctly 1",
			msg:  "success",
		},
		{
			name: "should get msg correctly 2",
			msg:  "yep",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", fmt.Sprintf("http://www.boilerplate.com/home?msg=%s", tt.msg), nil)

			h.Home(w, r)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Contains(t, fmt.Sprint(w.Body), fmt.Sprintf("\"msg\": \"%s\"", tt.msg))
		})
	}
}
