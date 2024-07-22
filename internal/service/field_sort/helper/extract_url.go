package helper

import (
	manage_str2 "card_detector/internal/util/str"
	"regexp"
	"strings"
)

var urlRegex = regexp.MustCompile(`(?:https?://)?([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)
var simpleUrlRegex = regexp.MustCompile(`^(?:https?://)?([a-zA-Z0-9])+([\-\.]{1}[a-zA-Z0-9]+)*\.[a-zA-Z]{2,}$`)
var emailCheckRegex = regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`)

func ExtractURL(text string) []string {
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

func ExtractBrokenUrl(text string, domains []string, zone string) string {
	if len(domains) == 0 || zone == "" {
		return ""
	}
	match := simpleUrlRegex.FindString(text)

	if strings.Contains(match, domains[0]) || strings.Contains(match, zone) {
		return match
	}
	return ""
}

// PAYFINANS.COM
func IsSimpleUrlAndCheck(line string) bool {
	if simpleUrlRegex.MatchString(line) && manage_str2.IsValidURL(line) {
		return true
	}
	return false
}

func IsSimpleUrl(line string) bool {
	return simpleUrlRegex.MatchString(line)
}
