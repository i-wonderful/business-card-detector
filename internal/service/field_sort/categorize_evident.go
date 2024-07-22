package field_sort

import (
	. "card_detector/internal/service/field_sort/helper"
	"strings"
)

const (
	FIELD_PHONE    = "phone"
	FIELD_TELEGRAM = "telegram"
	FIELD_SKYPE    = "skype"
	FIELD_EMAIL    = "email"
	FIELD_URLS     = "site"
	FIELD_NAME     = "name"
)

// categorizeEvidentFields - распределение очевидных полей
// phones, emails, telegram, skype, urls, name
// @return map[string]interface{}, []string - recognized fields and not recognized fields
func (s *Service) categorizeEvidentFields(data []string) (map[string]interface{}, []int) {
	notDetectItems := []int{}
	phones := []string{}
	emails := []string{}
	urls := []string{}
	telegram := []string{}
	var skype, name string

	for i, line := range data {
		line = ClearTrashSymbols(line)

		if s.processPhone(line, &phones) {
			continue
		}

		if s.processEmail(line, &emails, &skype) {
			continue
		}

		if s.processTelegram(line, &telegram) {
			continue
		}

		if skype == "" && s.processSkype(line, &skype) {
			continue
		}

		if s.processUrls(line, &urls) {
			continue
		}

		if name == "" && s.processNameByExistingNames(line, &name) {
			continue
		}

		notDetectItems = append(notDetectItems, i)
	}

	recognized := map[string]interface{}{
		FIELD_PHONE:    clearPhones(phones),
		FIELD_TELEGRAM: telegram,
		FIELD_SKYPE:    skype,
		FIELD_EMAIL:    emails,
		FIELD_URLS:     urls,
		FIELD_NAME:     name,
	}
	return recognized, notDetectItems
}

func (s *Service) processPhone(line string, phones *[]string) bool {
	if phoneRegex.MatchString(line) && NotContainsLetters(line) {
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

func (s *Service) processTelegram(line string, telegram *[]string) bool {
	if match := telegramRegex.FindString(line); match != "" {
		tg := strings.ReplaceAll(match, " ", "_")
		*telegram = append(*telegram, tg)
		return true
	}
	return false
}

func (s *Service) processSkype(line string, skype *string) bool {
	if sk := ExtractLiveSkype(line); sk != "" {
		*skype = sk
		return true
	}
	if sk := ExtractSkypeSkype(line); sk != "" {
		*skype = sk
		return true
	}
	return false
}

func (s *Service) processUrls(line string, urls *[]string) bool {
	if w := ExtractURL(line); len(w) > 0 {
		*urls = append(*urls, w...)
		return true
	}

	return false
}
