package core

import (
	"testing"
)

func TestReplaceUrl(t *testing.T) {
	tests := []struct {
		name      string
		url       string
		projectId string
		region    string
		expected  string
	}{
		{
			name:      "Valid placeholders",
			url:       "https://zt.{Region}.sms.api.sinch.com/xms/v1/{ProjectId}",
			projectId: "12345",
			region:    "us",
			expected:  "https://zt.us.sms.api.sinch.com/xms/v1/12345",
		},
		{
			name:      "No placeholders",
			url:       "https://example.com",
			projectId: "12345",
			region:    "us",
			expected:  "https://example.com",
		},
		{
			name:      "Partial placeholders",
			url:       "https://zt.{Region}.sms.api.sinch.com/xms/v1/static",
			projectId: "12345",
			region:    "eu",
			expected:  "https://zt.eu.sms.api.sinch.com/xms/v1/static",
		},
		{
			name:      "Empty URL",
			url:       "",
			projectId: "12345",
			region:    "us",
			expected:  "",
		},
		{
			name:      "Empty placeholders",
			url:       "https://zt.{Region}.sms.api.sinch.com/xms/v1/{ProjectId}",
			projectId: "",
			region:    "",
			expected:  "https://zt..sms.api.sinch.com/xms/v1/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceUrl(tt.url, tt.projectId, tt.region)
			if result != tt.expected {
				t.Errorf("ReplaceUrl() = %v, want %v", result, tt.expected)
			}
		})
	}
}