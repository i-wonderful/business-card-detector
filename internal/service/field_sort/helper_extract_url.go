package field_sort

import (
	manage_str2 "card_detector/internal/util/str"
	"regexp"
)

var urlRegex = regexp.MustCompile(`(?:https?://)?([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)

func extractURL(text string) []string {
	// Проверяем, содержит ли строка email
	if emailCheckRegex.MatchString(text) {
		return nil
	}
	if ContainsIgnoreCase(text, "skype") {
		return nil
	}

	// Если email не найден, извлекаем URL
	return simpleGetUrls(text)
}

func simpleGetUrls(val string) []string {
	match := urlRegex.FindAllString(val, -1)
	rez := make([]string, 0)
	for _, m := range match {
		if manage_str2.IsValidURL(m) {
			rez = append(rez, m)
		}
	}
	return rez
}

// PAYFINANS.COM
var simpleUrlRegex = regexp.MustCompile(`^[a-zA-Z0-9]+([\-\.]{1}[a-zA-Z0-9]+)*\.[a-zA-Z]{2,}$`)

func isSimpleUrl(line string) bool {
	if simpleUrlRegex.MatchString(line) && manage_str2.IsValidURL(line) {
		return true
	}
	return false
}
