package field_sort

import (
	manage_file "card_detector/internal/util/file"
	"log"
	"regexp"
	"strings"
	"time"
	"unicode"
)

type Service struct {
	professions []string
	companies   []string
	names       []string
	isLogTime   bool
}

func NewService(pathProfessions, pathCompanies, pathNames string, isLogTime bool) *Service {
	return &Service{
		professions: manage_file.ReadFile(pathProfessions),
		companies:   manage_file.ReadFile(pathCompanies),
		names:       manage_file.ReadFile(pathNames),
		isLogTime:   isLogTime,
	}
}

func (s *Service) Sort(data []string) map[string]interface{} {
	if s.isLogTime {
		start := time.Now()
		defer func() {
			log.Printf(">>> Time sort fields: %s", time.Since(start))
		}()
	}

	var name, company, jobTitle, skype, address string
	var mailName, domain, zone string
	phones := []string{}
	emails := []string{}
	websites := []string{}
	telegram := []string{}

	notDetectItems := []string{}

	recognized, data := s.categorizeEvidentFields(data)
	if _, ok := recognized[FIELD_TELEGRAM]; ok {
		telegram = recognized[FIELD_TELEGRAM].([]string)
	}
	if _, ok := recognized[FIELD_EMAIL]; ok {
		emails = recognized[FIELD_EMAIL].([]string)
		if len(emails) > 0 {
			mailName, domain, zone = extractMainNameDomainAndZone(emails[0])
		}
	}
	if _, ok := recognized[FIELD_SKYPE]; ok {
		skype = recognized[FIELD_SKYPE].(string)
	}

	if _, ok := recognized[FIELD_URLS]; ok {
		websites = recognized[FIELD_URLS].([]string)
	}

	for _, line := range data {
		line = clearTrashSymbols(line)

		if name == "" && isContainsWithSpace(line, s.names) {
			name = line
		} else if sk := extractSimpleSkype(skype, line); sk != "" {
			skype = sk
		} else if isContains(line, s.professions) {
			jobTitle += " " + line
		} else if company == "" && isContains(line, s.companies) {
			company = line
		} else if address == "" && ContainsIgnoreCase(line, "address") {
			address = line
		} else {
			notDetectItems = append(notDetectItems, line)
		}
	}

	other := []string{}
	for _, item := range notDetectItems {
		item = strings.TrimSpace(item)

		if len(emails) == 0 {
			if s := extractEmail(item); s != "" {
				emails = append(emails, s)
				continue
			}
		}
		if len(websites) == 0 {
			if isSimpleUrl(item) {
				websites = append(websites, item)
				continue
			}
			if s := extractBrokenUrl(item, domain, zone); s != "" {
				websites = append(websites, s)
				continue
			}
		}

		//if company == "" {
		if CheckIsPartOfDomain(item, domain) {
			if company != "" {
				other = append(other, item)
				continue
			}
			company = item
			continue
		}
		//if math := companyRegex.MatchString(item); math {
		//	company += " " + item
		//	continue
		//}
		//}

		if len(telegram) == 0 {
			if s := extractTelegram(item); s != "" {
				telegram = append(telegram, s)
				continue
			}
		}

		if name == "" {
			if m := nameRegex.MatchString(item); m {
				name = item
				continue
			}
		}
		if !NameIsFull(name) && ContainsIgnoreCase(item, mailName) || ContainsIgnoreCase(mailName, item) {
			name += " " + item
			continue
		}

		if m := phoneRegexExtract.MatchString(item); m {
			phNew := phoneRegexExtract.FindAllString(item, -1)
			phones = append(phones, phNew...)
			continue
		}

		other = append(other, item)
	}

	person := recognized
	person[FIELD_EMAIL] = emails
	person[FIELD_PHONE] = append(person[FIELD_PHONE].([]string), clearPhones(phones)...)
	person["name"] = strings.TrimSpace(name)
	person[FIELD_SKYPE] = skype
	person["company"] = strings.TrimSpace(company)
	person[FIELD_TELEGRAM] = telegram
	person[FIELD_URLS] = websites
	person["jobTitle"] = strings.TrimSpace(jobTitle)
	person["other"] = strings.Join(other, ";") + address

	return person
}

func isContains(s string, list []string) bool {
	s = strings.ToLower(s)
	for _, p := range list {
		if strings.Contains(s, p) {
			return true
		}
	}
	return false
}

// После найденной подстроки должны идти пробелы или конец строки.
func isContainsWithSpace(s string, list []string) bool {
	s = strings.ToLower(s)
	for _, p := range list {
		if strings.Contains(s, p) {
			ind := strings.Index(s, p) + len(p)
			if ind >= len(s) {
				return true
			} else if s[ind] == ' ' {
				return true
			} else if s[ind] == '-' {
				return true
			}
		}
	}
	return false
}

func extractBrokenUrl(text string, domain, zone string) string {
	//url := extractURL(text)
	//if url != "" {
	//	return url
	//}
	if domain == "" || zone == "" {
		return ""
	}

	urlRegex := regexp.MustCompile(`(?i)(www\.)?` + domain + `\s*\.?\s*` + zone + `\s*`)

	// Поиск совпадения в тексте
	match := urlRegex.FindString(text)
	if match == "" {
		return ""
	}

	if strings.HasPrefix(match, "www") {
		return "www." + domain + "." + zone
	}
	return domain + "." + zone
}

func extractMainNameDomainAndZone(email string) (string, string, string) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", "", "" // Неверный формат адреса электронной почты
	}
	domainParts := strings.Split(parts[1], ".")
	if len(domainParts) < 2 {
		return "", "", "" // Неверный формат домена
	}
	return parts[0], domainParts[0], domainParts[1]
}

func clearTrashSymbols(val string) string {
	val = strings.TrimFunc(val, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '+' && r != '@'
	})
	return val
}

func clearPhones(phones []string) []string {
	for i, phone := range phones {
		phones[i] = strings.TrimLeftFunc(phone, func(r rune) bool {
			return !unicode.IsDigit(r) && r != '+'
		})
		phones[i] = strings.TrimRightFunc(phone, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
	}

	rez := make([]string, 0)
	for _, phone := range phones {
		if countDigits(phone) > 8 {
			rez = append(rez, phone)
		}
	}

	return rez
}

func CheckIsPartOfDomain(item, domain string) bool {
	parts := strings.Split(item, " ")
	if len(parts) == 0 {
		return false
	}
	p0 := parts[0]
	return ContainsIgnoreCase(p0, domain) || ContainsIgnoreCase(domain, p0)
}

func ContainsIgnoreCase(s, substr string) bool {
	if s == "" || substr == "" {
		return false
	}
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

func countDigits(input string) int {
	count := 0
	for _, char := range input {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func notContainsLetters(phone string) bool {
	for i, r := range phone {
		if unicode.IsDigit(r) {
			// Проверить соседние символы на наличие букв
			if (i > 0 && unicode.IsLetter(rune(phone[i-1]))) || (i < len(phone)-1 && unicode.IsLetter(rune(phone[i+1]))) {
				return false
			}
		}
	}
	return true
}

func NameIsFull(val string) bool {
	names := strings.Split(val, " ")
	return len(names) >= 2
}

func trim(val []string) []string {
	var result []string
	for _, v := range val {
		result = append(result, strings.TrimSpace(v))
	}
	return result
}
