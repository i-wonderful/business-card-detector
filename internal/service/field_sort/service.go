package field_sort

import (
	manage_file "card_detector/internal/util/file"
	manage_str2 "card_detector/internal/util/str"
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

var emailRegex = regexp.MustCompile(`(?i)(?:Mail:\s*)?([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`)
var phoneRegex = regexp.MustCompile(`\+?[\d\s\(\)-]{6,16}\d`)
var phoneRegexExtract = regexp.MustCompile(`[\+\(]?[0-9 .\(\)-]{7,}`)
var nameRegex = regexp.MustCompile(`^[A-ZА-Я][A-ZА-Яa-zа-я-]+ [A-ZА-Я][A-ZА-Яa-zа-я-]+([ \-][A-ZА-Я][A-ZА-Яa-zа-я-]+)?$`)
var telegramRegex = regexp.MustCompile(`(?:https?://)?(t\.me/|@)[A-Za-z][A-Za-z0-9_]{4,31}(?:\s[A-Za-z0-9_]+)*`)
var urlRegex = regexp.MustCompile(`(?:https?://)?([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)
var skypeRegex = regexp.MustCompile(`^[a-z]{4,31}\.[a-z]{4,31}`)
var skypeRegexLive = regexp.MustCompile(`(?i)live\s*[:.]{0,2}\s*cid\.([\da-f]+)`)
var skypeRegexLiveUser = regexp.MustCompile(`(?i)live:([\w\-\.]+)`)
var emailCheckRegex = regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`)
var emailExtractRegex = regexp.MustCompile(`[\S\s]+@[\S\s]+\.[\S\s]+`)

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

	var name, company, jobTitle, telegram, skype, address string
	var mailName, domain, zone string
	phones := []string{}
	emails := []string{}
	websites := []string{}

	notDetectItems := []string{}

	recognized, data := s.evidentSort(data)

	for _, line := range data {
		line = clearTrashSymbols(line)

		if match := emailRegex.FindStringSubmatch(line); len(match) > 1 {
			findEmail := strings.Replace(match[1], " ", "", -1)
			if !isContains(findEmail, emails) && !ContainsIgnoreCase(line, "skype") {
				emails = append(emails, findEmail)
			} else if skype == "" { // скайп может совпадать с email
				skype = findEmail
			}

			// domain may be company
			// mailName may be Name
			mailName, domain, zone = extractMainNameDomainAndZone(findEmail)

			line = strings.ReplaceAll(line, match[1], "")
			if len(line) > 5 {
				notDetectItems = append(notDetectItems, line)
			}
		} else if match := telegramRegex.FindString(line); match != "" && telegram == "" {
			telegram = strings.ReplaceAll(match, " ", "_")
		} else if w := extractURL(line); len(w) > 0 && len(websites) == 0 {
			websites = w
		} else if name == "" && isContainsWithSpace(line, s.names) {
			name = line
			/* else if len(phones) == 0 && phoneRegex.MatchString(line) && notContainsLetters(line) {
				p := phoneRegex.FindString(line)
				phones = append(phones, p)
			} */
		} else if sk := s.extractSimpleSkype(skype, line); sk != "" {
			skype = sk
		} else if isContainsWithSpace(line, s.professions) {
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
			if s := extractBrokenUrl(item, domain, zone); s != "" {
				websites = append(websites, s)
				continue
			}
		}
		if skype == "" {
			if s := extractSkype(item); s != "" {
				skype = s
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

		if telegram == "" {
			if s := extractTelegram(item); s != "" {
				telegram = s
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
	person["email"] = emails
	person["phone"] = append(person["phone"].([]string), clearPhones(phones)...)
	person["name"] = strings.TrimSpace(name)
	person["skype"] = skype
	person["company"] = strings.TrimSpace(company)
	person["telegram"] = telegram
	person["site"] = websites
	person["jobTitle"] = strings.TrimSpace(jobTitle)
	person["other"] = strings.Join(other, ";") + address

	return person
}

func (s *Service) extractSimpleSkype(skype string, line string) string {
	if skype != "" {
		return ""
	}

	if skypeRegex.MatchString(line) {
		return line
	}

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

func extractEmail(item string) string {
	match := emailExtractRegex.FindStringSubmatch(item)
	if len(match) > 0 {
		return strings.ReplaceAll(match[0], " ", "")
	}
	return ""
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

func extractSkype(text string) string {
	// Определяем шаблон регулярного выражения
	pattern := `(?i)(skype\s*[:\.]?\s*|s:)([a-zA-Z0-9\.\-_]+(?:@\w+\.[\w.]+)?)`

	re := regexp.MustCompile(pattern)
	subs := re.FindStringSubmatch(text)
	if len(subs) > 2 {
		return subs[2]
	}
	return ""
}
func extractTelegram(text string) string {
	pattern := `(?i)(telegram\s*:?\s*)(\w+)`
	regex := regexp.MustCompile(pattern)

	matches := regex.FindStringSubmatch(text)
	if len(matches) > 2 {
		return matches[2]
	}
	return ""
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
