package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractDomainName(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		expected string
	}{
		{"HTTP with www", "http://www.alianspay.com", "alianspay"},
		{"HTTPS without www", "https://alianspay.com", "alianspay"},
		{"WWW without protocol", "www.alianspay.com", "alianspay"},
		{"Domain only", "alianspay.com", "alianspay"},
		{"Short domain", "mycompany.by", "mycompany"},
		{"HTTP short domain", "http://mycompany.by", "mycompany"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ExtractDomainFromUrl(tc.url)
			assert.Equal(t, tc.expected, result, "URL: %s", tc.url)
		})
	}
}
