package field_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_FindEmail(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected map[string]interface{}
	}{
		{
			name: "Email with space",
			input: []string{
				"Martin@369gaming. media",
			},
			expected: map[string]interface{}{
				"email": []string{"Martin@369gaming.media"},
			},
		},
		{
			name: "Email with 'E-mail'",
			input: []string{
				"E-mail: christoffer.froberg@qpgames.se",
			},
			expected: map[string]interface{}{
				"email": []string{"christoffer.froberg@qpgames.se"},
			},
		},
		{
			name: "Simple Email",
			input: []string{
				"ivk@colibrix.io",
			},
			expected: map[string]interface{}{
				"email": []string{"ivk@colibrix.io"},
			},
		},
		{
			"Many emails",
			[]string{
				"SUPPORT@HUGE.PARTNERS",
				"INFO@HUGE.PARTNERS",
			},
			map[string]interface{}{
				"email": []string{
					"SUPPORT@HUGE.PARTNERS",
					"INFO@HUGE.PARTNERS",
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
			result := service.Sort(tc.input)
			for k, v := range tc.expected {
				assert.Equal(t, v, result[k])
			}
		})
	}
}

func TestExtractEmail(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "With spaces",
			input:    "Martin@369gaming. media",
			expected: "Martin@369gaming.media",
		},
		//{
		//	"With 'E-mail'",
		//	"E-mail:christoffer.froberg@qpgames.se",
		//	"christoffer.froberg@qpgames.se",
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := extractEmail(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
