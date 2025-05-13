package core

import (
	"bytes"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRequest(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		token      string
		opt        any
		statusCode int
		response   string
		expectErr  bool
	}{
		{
			name:       "GET request with valid token",
			method:     http.MethodGet,
			path:       "/test-path",
			token:      "valid-token",
			opt:        nil,
			statusCode: http.StatusOK,
			response:   `{"message":"success"}`,
			expectErr:  false,
		},
		{
			name:       "POST request with valid token and body",
			method:     http.MethodPost,
			path:       "/test-path",
			token:      "valid-token",
			opt:        map[string]string{"key": "value"},
			statusCode: http.StatusCreated,
			response:   `{"message":"created"}`,
			expectErr:  false,
		},
		{
			name:       "Invalid token",
			method:     http.MethodGet,
			path:       "/test-path",
			token:      "invalid-token",
			opt:        nil,
			statusCode: http.StatusUnauthorized,
			response:   `{"error":"unauthorized"}`,
			expectErr:  true,
		},
		{
			name:       "Invalid method",
			method:     "INVALID",
			path:       "/test-path",
			token:      "valid-token",
			opt:        nil,
			statusCode: http.StatusMethodNotAllowed,
			response:   `{"error":"method not allowed"}`,
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock HTTP server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("Authorization") != "Bearer "+tt.token {
					http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
					return
				}
				if r.Method != tt.method {
					http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
					return
				}
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.response))
			}))
			defer server.Close()

			// Prepare request body if applicable
			var opt any
			if tt.opt != nil {
				opt = tt.opt
			}

			// Call NewRequest
			result, err := NewRequest(tt.method, server.URL+tt.path, tt.token, opt)

			// Validate results
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got: %v", err)
				}
				if !bytes.Equal(result, []byte(tt.response)) {
					t.Errorf("expected response %s, got %s", tt.response, string(result))
				}
			}
		})
	}
}

func TestNewRequest_InvalidJSON(t *testing.T) {
	// Mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"success"}`))
	}))
	defer server.Close()

	// Invalid JSON input
	invalidOpt := make(chan int)

	_, err := NewRequest(http.MethodPost, server.URL+"/test-path", "valid-token", invalidOpt)
	if err == nil {
		t.Errorf("expected an error for invalid JSON input but got none")
	}
}

func TestNewRequest_CheckResponseError(t *testing.T) {
	// Mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"internal server error"}`))
	}))
	defer server.Close()

	_, err := NewRequest(http.MethodGet, server.URL+"/test-path", "valid-token", nil)
	if err == nil {
		t.Errorf("expected an error for server error but got none")
	}
}


