package manage_str

import "unicode"

// IsOnlyLetters checks if the input string contains only letters.
//
// Parameter: val string - the string to check.
// Return type: bool - true if the string contains only letters, spaces, or hyphens; false otherwise.
func IsOnlyLetters(val string) bool {
	for _, r := range val {
		if !unicode.IsLetter(r) /*  && r != '-' && r != ' '  && r != ',' && r != '.'*/ { // todo ???
			return false
		}
	}
	return true
}
