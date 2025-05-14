package core

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestCheckResponse(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		body           string
		expectedError  error
		expectedErrMsg string
	}{
		{
			name:          "No error for 200 status code",
			statusCode:    200,
			body:          "",
			expectedError: nil,
		},
		{
			name:          "No error for 201 status code",
			statusCode:    201,
			body:          "",
			expectedError: nil,
		},
		{
			name:          "Error for 404 status code",
			statusCode:    404,
			body:          "",
			expectedError: ErrNotFound,
		},
		{
			name:           "Error with unknown format body",
			statusCode:     500,
			body:           "Internal Server Error",
			expectedError:  &ErrorResponse{},
			expectedErrMsg: "failed to parse unknown error format: Internal Server Error",
		},
		{
			name:           "Error with JSON body",
			statusCode:     400,
			body:           `{"error": "Invalid request"}`,
			expectedError:  &ErrorResponse{},
			expectedErrMsg: "{error: Invalid request}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := io.NopCloser(bytes.NewReader([]byte(tt.body)))
			response := &http.Response{
				StatusCode: tt.statusCode,
				Body:       body,
				Request:    &http.Request{Method: "POST", URL: &url.URL{Scheme: "https", Host: "example.com", Path: "/test"}},
			}

			err := CheckResponse(response)

			if tt.expectedError == nil {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error, got nil")
				} else {
					errorResponse, ok := err.(*ErrorResponse)
					if ok && tt.expectedErrMsg != "" && errorResponse.Message != tt.expectedErrMsg {
						t.Errorf("expected error message %q, got %q", tt.expectedErrMsg, errorResponse.Message)
					}
				}
			}
		})
	}
}
func TestParseError(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "String input",
			input:    "An error occurred",
			expected: "An error occurred",
		},
		{
			name:     "Array of strings",
			input:    []interface{}{"Error 1", "Error 2"},
			expected: "[Error 1, Error 2]",
		},
		{
			name: "Array of mixed types",
			input: []interface{}{
				"Error 1",
				map[string]interface{}{"key": "value"},
			},
			expected: "[Error 1, {key: value}]",
		},
		{
			name: "Map with string values",
			input: map[string]interface{}{
				"error1": "Error 1",
				"error2": "Error 2",
			},
			expected: "{error1: Error 1, error2: Error 2}",
		},
		{
			name: "Map with nested structures",
			input: map[string]interface{}{
				"error1": "Error 1",
				"nested": map[string]interface{}{
					"key": "value",
				},
			},
			expected: "{error1: Error 1, nested: {key: value}}",
		},
		{
			name:     "Unexpected type",
			input:    123,
			expected: "failed to parse unexpected error type: int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseError(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

