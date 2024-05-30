package field_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_FindSite(t *testing.T) {

	testCases := []struct {
		name     string
		input    []string
		expected map[string]interface{}
	}{
		{
			"Simple site",
			[]string{
				"www.tornado-games.com",
			},
			map[string]interface{}{
				"site": "www.tornado-games.com",
			},
		},
	}
	service := NewService(
		"../../../config/professions.txt",
		"../../../config/companies.txt",
		"../../../config/names.txt",
		true)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := service.Sort(tc.input)
			for k, v := range tc.expected {
				assert.Equal(t, v, result[k])
			}
		})
	}
}
func TestExtractURL(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid URL",
			input:    "Visit https://example.com for more information.",
			expected: "example.com",
		},
		{
			name:     "URL without scheme",
			input:    "The site is example.org.",
			expected: "example.org",
		},
		{
			name:     "URL with path",
			input:    "Check out https://www.example.net/path/to/page",
			expected: "www.example.net",
		},
		{
			name:     "URL with query parameters",
			input:    "https://search.example.com?q=query&page=2",
			expected: "search.example.com",
		},
		{
			name:     "Invalid URL",
			input:    "This is not a URL.",
			expected: "",
		}, {
			name:     "without http or https",
			input:    "www.admill.io",
			expected: "www.admill.io",
		},
		{
			name:     "with site world",
			input:    "Site: Linebet.com Telegram:",
			expected: "Linebet.com",
		},
		{
			name:     "email not URL",
			input:    "Mail:  Skype: partners@Linebet.com",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := extractURL(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
