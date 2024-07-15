package field_sort

import "strings"

const (
	FIELD_PHONE    = "phone"
	FIELD_TELEGRAM = "telegram"
	FIELD_SKYPE    = "skype"
	FIELD_EMAIL    = "email"
	FIELD_URLS     = "site"
)

// categorizeEvidentFields - распределение очевидных полей
// phones, emails, telegram, skype, urls
// @return map[string]interface{}, []string - recognized fields and not recognized fields
func (s *Service) categorizeEvidentFields(data []string) (map[string]interface{}, []string) {
	notDetectItems := []string{}
	phones := []string{}
	emails := []string{}
	urls := []string{}
	telegram := ""
	skype := ""

	for _, line := range data {
		line = clearTrashSymbols(line)

		if s.processPhone(line, &phones) {
			continue
		}

		if s.processEmail(line, &emails, &skype) {
			continue
		}

		if s.processTelegram(line, &telegram) {
			continue
		}

		if s.processSkype(line, &skype) {
			continue
		}

		if s.processUrls(line, &urls) {
			continue
		}

		notDetectItems = append(notDetectItems, line)
	}

	recognized := map[string]interface{}{
		FIELD_PHONE:    clearPhones(phones),
		FIELD_TELEGRAM: telegram,
		FIELD_SKYPE:    skype,
		FIELD_EMAIL:    emails,
		FIELD_URLS:     urls,
	}
	return recognized, notDetectItems
}

func (s *Service) processPhone(line string, phones *[]string) bool {
	if phoneRegex.MatchString(line) && notContainsLetters(line) {
		p := phoneRegex.FindString(line)
		*phones = append(*phones, p)
		return true
	}
	return false
}

func (s *Service) processEmail(line string, emails *[]string, skype *string) bool {
	if match := emailRegex.FindStringSubmatch(line); len(match) > 1 {
		findEmail := strings.Replace(match[1], " ", "", -1)
		if !isContains(findEmail, *emails) && !ContainsIgnoreCase(line, "skype") {
			*emails = append(*emails, findEmail)
		} else if *skype == "" {
			*skype = findEmail
		}
		return true
	}
	return false
}

func (s *Service) processTelegram(line string, telegram *string) bool {
	if match := telegramRegex.FindString(line); match != "" && *telegram == "" {
		*telegram = strings.ReplaceAll(match, " ", "_")
		return true
	}
	return false
}

func (s *Service) processSkype(line string, skype *string) bool {
	if sk := extractLiveSkype(line); sk != "" {
		*skype = sk
		return true
	}
	if sk := extractSkypeSkype(line); sk != "" {
		*skype = sk
		return true
	}
	return false
}

func (s *Service) processUrls(line string, urls *[]string) bool {
	if w := extractURL(line); len(w) > 0 {
		*urls = append(*urls, w...)
		return true
	}

	return false
}
