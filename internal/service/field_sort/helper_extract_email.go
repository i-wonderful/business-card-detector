package field_sort

import (
	"regexp"
	"strings"
)

var emailExtractRegex = regexp.MustCompile(`[\S\s]+@[\S\s]+\.[\S\s]+`)

func extractEmail(item string) string {
	match := emailExtractRegex.FindStringSubmatch(item)
	if len(match) > 0 {
		return strings.ReplaceAll(match[0], " ", "")
	}
	return ""
}
