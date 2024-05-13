package text_recognize

import "testing"

func TestIsPhoneString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid string with 4 digits and 2 letters",
			input:    "123a,c4",
			expected: true,
		},
		{
			name:     "Invalid string with more than 2 letters",
			input:    "1234 abcd",
			expected: false,
		},
		{
			name:     "Invalid string with less than 4 digits",
			input:    "123",
			expected: false,
		},
		{
			name:     "Valid string with only digits",
			input:    "12345",
			expected: true,
		},
		{
			name:     "Invalid string with no digits",
			input:    "abc",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPhone(tc.input)
			if result != tc.expected {
				t.Errorf("isValidString(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
