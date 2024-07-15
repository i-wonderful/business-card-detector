package field_sort

import "regexp"

var telegramWorldRegex = regexp.MustCompile(`(?i)(telegram\s*:?\s*)(\w+)`)

func extractTelegram(text string) string {
	matches := telegramWorldRegex.FindStringSubmatch(text)
	if len(matches) > 2 {
		return matches[2]
	}
	return ""
}
