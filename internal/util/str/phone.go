package manage_str

import "regexp"

func FindAndRemovePhone(strs []string) string {
	phoneRegex := regexp.MustCompile(`^\+\d[\d\-\s()+]+$`)

	for i, s := range strs {
		if phoneRegex.MatchString(s) {
			phone := strs[i]
			strs = append(strs[:i], strs[i+1:]...)
			return phone
		}
	}

	return ""
}
