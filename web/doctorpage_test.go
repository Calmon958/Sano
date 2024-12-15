package web

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	Template_dir = "assets/template"
	Templates    *template.Template
)

// GET request returns 200 status code and renders DoctorDashboard template
func TestDoctorPageHandlerGetRequest(t *testing.T) {
	tmpl, err := template.ParseGlob("../assets/template/*.html")
	if err != nil {
		t.Fatalf("Failed to parse templates: %v", err)
	}
	Templates = tmpl

	tests := []struct {
		name                 string
		method               string
		expected_status_code int
	}{
		{
			name:                 "GET request",
			method:               http.MethodGet,
			expected_status_code: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/DashboardDoctor", nil)
			w := httptest.NewRecorder()
			DoctorPageHandler(w, request)
			fmt.Println(w.Code)
			if w.Code != tc.expected_status_code {
				t.Errorf("expected status %v, got %v", tc.expected_status_code, w.Code)
			}
			if !strings.Contains(w.Body.String(), "DoctorDashboard") {
				t.Error("response does not contain expected template content")
			}
		})
	}
}
