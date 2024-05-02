package field_sort

import (
	"testing"
)

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

func TestExtractTelegram(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid telegram handle",
			input:    "My telegram handle is: Telegram: johndoe",
			expected: "johndoe",
		},
		{
			name:     "Valid telegram handle with whitespace",
			input:    "My telegram handle is: Telegram:   johndoe123",
			expected: "johndoe123",
		},
		{
			name:     "Case insensitive",
			input:    "My Telegram handle is: tELEGRAM: janedoe",
			expected: "janedoe",
		},
		{
			name:     "No telegram handle",
			input:    "This string does not contain a telegram handle",
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

func TestExtractSkype(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid Skype ID",
			input:    "My Skype ID is: skype.sample_id",
			expected: "sample_id",
		},
		//{
		//	name:     "Valid Skype ID with spaces",
		//	input:    "My Skype ID is: skype.live: sample_id",
		//	expected: "sample_id",
		//},
		{
			name:     "Valid Skype ID with case-insensitive pattern",
			input:    "My Skype ID is: Skype: SAMPLE_ID",
			expected: "SAMPLE_ID",
		},
		{
			name:     "No Skype ID",
			input:    "This text does not contain a Skype ID",
			expected: "",
		},
		{
			name:     "Multiple Skype IDs",
			input:    "My Skype IDs are: skype:id1 and skype:id2",
			expected: "id1",
		},
		{
			name:     "Extract with @",
			input:    "Mail: b2b@lLinebet.com Skype: partners@Linebet.com",
			expected: "partners@Linebet.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := extractSkype(test.input)
			if result != test.expected {
				t.Errorf("Expected %q, but got %q", test.expected, result)
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
func TestService_Sort(t *testing.T) {

	testCases := []struct {
		name     string
		input    []string
		expected map[string]interface{}
	}{
		{
			// len:3, cap:3
			name: "Sort names",
			input: []string{
				"Site: Linebet.com Telegram: @linebet partners bot",
				"Mail: b2b@lLinebet.com Skype: partners@Linebet.com",
				"B2B Department",
			},
			expected: map[string]interface{}{
				"site":     "Linebet.com",
				"telegram": "@linebet",
				"email":    "b2b@lLinebet.com",
				"skype":    "partners@Linebet.com",
				"company":  "B2B Department",
			},
		},
		{
			name: "Email with space",
			input: []string{
				"Martin@369gaming. media",
			},
			expected: map[string]interface{}{
				"email": "Martin@369gaming.media",
			},
		},
		{
			name: "Email with 'E-mail'",
			input: []string{
				"E-mail: christoffer.froberg@qpgames.se",
			},
			expected: map[string]interface{}{
				"email": "christoffer.froberg@qpgames.se",
			},
		},
		{
			// ivk@colibrix.io
			name: "Simple Email",
			input: []string{
				"ivk@colibrix.io",
			},
			expected: map[string]interface{}{
				"email": "ivk@colibrix.io",
			},
		},
		{
			"Simpe site",
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
				if result[k] != v {
					t.Errorf("Field %q. Expected %q, got %q", k, v, result[k])
				}
			}
		})
	}
}
