package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "GET request to non-root path",
			method:         http.MethodGet,
			url:            "/nonexistent",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 Page Not Found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			// create a response recorder

			rr := httptest.NewRecorder()

			HomeHandler(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if !strings.Contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

func TestAsciiArtHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		formData       map[string]string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "POST request with empty text",
			method:         http.MethodPost,
			url:            "/ascii-art",
			formData:       map[string]string{"text": "", "banner": "standard"},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "",
		},
		{
			name:           "POST request with empty banner",
			method:         http.MethodPost,
			url:            "/ascii-art",
			formData:       map[string]string{"text": "Hello", "banner": ""},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "",
		},
		{
			name:           "PUT request (not allowed)",
			method:         http.MethodPut,
			url:            "/ascii-art",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			var err error

			if tt.method == http.MethodPost {
				formData := url.Values{}
				for key, value := range tt.formData {
					formData.Set(key, value)
				}
				req, err = http.NewRequest(tt.method, tt.url, strings.NewReader(formData.Encode()))
				if err != nil {
					t.Fatalf("could not create request: %v", err)
				}
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			} else {
				req, err = http.NewRequest(tt.method, tt.url, nil)
				if err != nil {
					t.Fatalf("could not create request: %v", err)
				}
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(AsciiArtHandler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if !strings.Contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

