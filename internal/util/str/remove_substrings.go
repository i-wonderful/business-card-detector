package manage_str

import (
	"strings"
)

// RemoveSubstrings removes all substrings from the array
func RemoveSubstrings(arr []string) []string {
	var result []string

	containsSubstring := func(s string, arr []string) bool {
		for _, str := range arr {
			if strings.Contains(str, s) && s != str {
				return true
			}
		}
		return false
	}

	for _, str := range arr {
		if !containsSubstring(str, arr) || IsValidURL(str) {
			result = append(result, str)
		}
	}

	return result
}
