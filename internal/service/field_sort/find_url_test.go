package field_sort

import (
	"card_detector/internal/service/field_sort/helper"
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
				"site": []string{"www.tornado-games.com"},
			},
		},
		{
			"Concated",
			[]string{
				"slotscalendar.com|betbrain.com",
			},
			map[string]interface{}{
				"site": []string{
					"slotscalendar.com",
					"betbrain.com",
				},
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
			result := service.Sort(wrapForSort(tc.input), nil)
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
		expected []string
	}{
		{
			name:     "Valid URL",
			input:    "Visit https://example.com for more information.",
			expected: []string{"https://example.com"},
		},
		{
			name:     "URL without scheme",
			input:    "The site is example.org.",
			expected: []string{"example.org"},
		},
		{
			name:     "URL with path",
			input:    "Check out https://www.example.net/path/to/page",
			expected: []string{"https://www.example.net/path/to/page"},
		},
		{
			name:     "URL with query parameters",
			input:    "https://search.example.com?q=query&page=2",
			expected: []string{"https://search.example.com?q=query&page=2"},
		},
		{
			name:     "Invalid URL",
			input:    "This is not a URL.",
			expected: []string{},
		}, {
			name:     "without http or https",
			input:    "www.admill.io",
			expected: []string{"www.admill.io"},
		},
		{
			name:     "with site world",
			input:    "Site: Linebet.com Telegram:",
			expected: []string{"Linebet.com"},
		},
		{
			name:     "email not URL",
			input:    "Mail:  Skype: partners@Linebet.com",
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := helper.ExtractURL(tc.input)
			for k, v := range tc.expected {
				assert.Equal(t, v, result[k])
			}
		})
	}
}
