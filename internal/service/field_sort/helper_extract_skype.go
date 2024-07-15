package field_sort

import "regexp"

var skypeRegex = regexp.MustCompile(`^[a-z]{4,31}\.[a-z]{4,31}`)
var skypeRegexLive = regexp.MustCompile(`(?i)live\s*[:.]{0,2}\s*cid\.([\da-f]+)`)
var skypeRegexLiveUser = regexp.MustCompile(`(?i)live:([\w\-\.]+)`)
var skypeSkypeRegex = regexp.MustCompile(`(?i)(skype\s*[:\.]?\s*|s:)([a-zA-Z0-9\.\-_]+(?:@\w+\.[\w.]+)?)`)

func extractSkypeSkype(text string) string {
	subs := skypeSkypeRegex.FindStringSubmatch(text)
	if len(subs) > 2 {
		return subs[2]
	}
	return ""
}

func extractLiveSkype(line string) string {
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

func extractSimpleSkype(skype string, line string) string {
	if skype != "" {
		return ""
	}

	if skypeRegex.MatchString(line) {
		return line
	}

	return ""
}
