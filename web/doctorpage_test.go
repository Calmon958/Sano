package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// GET request returns successful response with HTTP 200
func TestDoctorPageHandlerGetRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/doctor", nil)
	w := httptest.NewRecorder()

	DoctorPageHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
