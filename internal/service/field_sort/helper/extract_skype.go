package helper

import "regexp"

var skypeRegex = regexp.MustCompile(`^[a-z]{4,31}\.[a-z]{4,31}`)
var skypeRegexLive = regexp.MustCompile(`(?i)live\s*[:.]{0,2}\s*cid\.([\da-f]+)`)
var skypeRegexLiveUser = regexp.MustCompile(`(?i)live:([\w\-\.]+)`)
var skypeSkypeRegex = regexp.MustCompile(`(?i)(?:\bskype\s*:\s*|\bs\s*:\s*|\bskype\s+)([\w\.\-_@]+)\b`)

// ExtractSkypeSkype extracts Skype ID from the given text using a predefined regex pattern
func ExtractSkypeSkype(text string) string {
	subs := skypeSkypeRegex.FindStringSubmatch(text)
	if len(subs) > 1 {
		return subs[1]
	}
	return ""
}

func ExtractLiveSkype(line string) string {
	matches := skypeRegexLive.FindStringSubmatch(line)
	if len(matches) > 1 {
		return "live:cid." + matches[1]
	}
	matches = skypeRegexLiveUser.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[0]
	}

	return ""
}

func ExtractSimpleSkype(skype string, line string) string {
	if skype != "" {
		return ""
	}

	if skypeRegex.MatchString(line) {
		return line
	}

	return ""
}
