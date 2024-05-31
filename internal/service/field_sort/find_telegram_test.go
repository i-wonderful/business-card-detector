package field_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_FindTelegram(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected map[string]interface{}
	}{
		{
			"With spaces",
			[]string{"@Eska8 Aff"},
			map[string]interface{}{
				"telegram": "@Eska8_Aff",
			},
		},
		{
			"Without @",
			[]string{"TELEGRAM ADV_ADSCOMPASS"},
			map[string]interface{}{
				"telegram": "ADV_ADSCOMPASS",
			},
		},
		{
			"By url",
			[]string{
				"https://t.me/Nicola_an",
				"www.admillio",
				"erkin@admillio",
			},
			map[string]interface{}{
				"telegram": "https://t.me/Nicola_an",
			},
		},
		{
			"Url without http",
			[]string{"t.me/Taras_CoinsPaid"},
			map[string]interface{}{
				"telegram": "t.me/Taras_CoinsPaid",
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

func TestExtractTelegram(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid telegram",
			input:    "My telegram : johndoe",
			expected: "johndoe",
		},
		{
			name:     "Valid telegram handle with whitespace",
			input:    "My telegram :    johndoe123",
			expected: "johndoe123",
		},
		{
			name:     "Case insensitive",
			input:    " tELEGRAM: janedoe",
			expected: "janedoe",
		},
		{
			name:     "No telegram handle",
			input:    "This string does not contain a telegram ",
			expected: "",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractTelegram(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, actual)
			}
		})
	}
}
