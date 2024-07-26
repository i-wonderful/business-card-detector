package manage_str

import "strings"

// IsContains - check if string in list
func IsContains(s string, list []string) bool {
	isOk, _ := IsContainsWith(s, list)
	return isOk
}

// IsContainsWith - check if string in list and return found
func IsContainsWith(s string, list []string) (bool, string) {
	s = strings.ToLower(s)
	for _, p := range list {
		if strings.Contains(s, strings.ToLower(p)) {
			return true, p
		}
	}
	return false, ""
}
