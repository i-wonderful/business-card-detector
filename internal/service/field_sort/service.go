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

	//websitePattern := `^[a-zA-Z0-9\.-]+\.[a-zA-Z]{2,}$`
	//addressPattern := `^[\w\s\p{L},0-9()/\\-]*$`

	//emaiNewRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,11}`)
	//emailNewRegex := regexp.MustCompile(`(?:Mail:\s*)?([\S\s]+?@[\S\s]+?\.\S+)\b`)
	emailNewRegex := regexp.MustCompile(`(?i)(?:Mail:\s*)?([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`)

	//regexp.MustCompile(`[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+`) // email
	//phoneRegex := regexp.MustCompile(`^[\+]?[(]?[0-9\s-]{6,}[)]?$`)
	phoneRegex := regexp.MustCompile(`^[\+]?[(]?\d{1,3}[)]?[\s-]?\d+(?:[\s-]?\d+)*$`)
	//phoneRegexExtract := regexp.MustCompile(`[\+]?[(]?\d{1,3}[)]?[\s-]?\d+(?:[\s-]?\d+)*`)
	phoneRegexExtract := regexp.MustCompile(`[\+\(]?[0-9 .\(\)-]{7,}`)

	//websiteRegex := regexp.MustCompile(websitePattern)
	//addressRegex := regexp.MustCompile(addressPattern)

	//companyRegex := regexp.MustCompile(`^[A-Z]{3,}[ \.]*$|^[A-Z]+$`)
	nameRegex := regexp.MustCompile(`^[A-ZА-Я][A-ZА-Яa-zа-я-]+ [A-ZА-Я][A-ZА-Яa-zа-я-]+([ \-][A-ZА-Я][A-ZА-Яa-zа-я-]+)?$`)
	//singleName := regexp.MustCompile("^[А-ЯA-Z][а-яa-z]{1,}$")
	//telegramRegex := regexp.MustCompile(`(?:https?://t\.me/|@)[A-Za-z][A-Za-z0-9_]{4,31}`)
	//	telegramRegex := regexp.MustCompile(`(?:https?://t\.me/|@)([A-Za-z][A-Za-z0-9_]{4,31}(?:\s[A-Za-z0-9_]+)*)`)
	telegramRegex := regexp.MustCompile(`(?:https?://t\.me/|@)[A-Za-z][A-Za-z0-9_]{4,31}(?:\s[A-Za-z0-9_]+)*`)

	var name, company, jobTitle, telegram, website, skype string
	var mailName, domain, zone string
	phones := []string{}
	emails := []string{}

	person := map[string]interface{}{}
	notDetectItems := []string{}

	for _, line := range data {
		line = clearTrashSymbols(line)

		if match := emailNewRegex.FindStringSubmatch(line); len(match) > 1 {
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
		} else if isContainsWithSpace(line, s.professions) {
			jobTitle += " " + line
		} else if match := telegramRegex.FindString(line); match != "" && telegram == "" {
			telegram = strings.ReplaceAll(match, " ", "_")
		} else if w := extractURL(line); w != "" && website == "" { // todo уменьшить вызовы
			website = w
		} else if name == "" && isContainsWithSpace(line, s.names) {
			name = line
		} else if len(phones) == 0 && phoneRegex.MatchString(line) {
			phones = phoneRegex.FindAllString(line, -1)
		} else if sk := s.extractSimpleSkype(skype, line); sk != "" {
			skype = sk
		} else if company == "" && isContains(line, s.companies) {
			company = line
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
		if website == "" {
			if s := extractBrokenUrl(item, domain, zone); s != "" {
				website = s
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
		if ContainsIgnoreCase(item, domain) || ContainsIgnoreCase(domain, item) {
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
		if ContainsIgnoreCase(item, mailName) || ContainsIgnoreCase(mailName, item) {
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

	person["email"] = emails
	person["phone"] = clearPhones(phones)
	person["name"] = strings.TrimSpace(name)
	person["skype"] = skype
	person["company"] = strings.TrimSpace(company)
	person["telegram"] = telegram
	person["site"] = website
	person["jobTitle"] = strings.TrimSpace(jobTitle)
	person["other"] = strings.Join(other, ";")

	return person
}

func (s *Service) extractSimpleSkype(skype string, line string) string {
	if skype != "" {
		return ""
	}
	skypeRegex := regexp.MustCompile(`^[a-z]{4,31}\.[a-z]{4,31}`)
	//skypeRegexLive := regexp.MustCompile(`live:([.:])?cid\.[0-9a-f]+`)
	skypeRegexLive := regexp.MustCompile(`live:([.:]?cid\.[0-9a-f]+)`)

	if skypeRegex.MatchString(line) {
		return line
	}

	matches := skypeRegexLive.FindStringSubmatch(line)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

func extractEmail(item string) string {
	emailRegex := regexp.MustCompile(`[\S\s]+@[\S\s]+\.[\S\s]+`)
	match := emailRegex.FindStringSubmatch(item)
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

func extractURL(text string) string {
	// Проверяем, содержит ли строка email
	emailRegex := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`)
	if emailRegex.MatchString(text) {
		return ""
	}
	if ContainsIgnoreCase(text, "skype") {
		return ""
	}

	// Если email не найден, извлекаем URL
	return simpleGetUrl(text)
}

func simpleGetUrl(val string) string {
	urlRegex := regexp.MustCompile(`(?:https?://)?([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)
	match := urlRegex.FindStringSubmatch(val)
	if len(match) > 1 && manage_str2.IsValidURL(match[1]) {

		return match[1]
	}
	return ""
}

func extractBrokenUrl(text string, domain, zone string) string {
	url := extractURL(text)
	if url != "" {
		return url
	}
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
	//pattern := `(?i)skype\s*[:\.]?\s*([a-zA-Z0-9\.\-_]+(?:@\w+\.[\w.]+)?)`
	pattern := `(?i)skype\s*[:\.]?\s*([a-zA-Z0-9\.\-_]+(?:@\w+\.[\w.]+)?)`
	re := regexp.MustCompile(pattern)
	subs := re.FindStringSubmatch(text)
	if len(subs) > 1 {
		return subs[1]
	}
	return ""
}
func extractTelegram(text string) string {
	pattern := `(?i)telegram:\s*(\w+)`
	regex := regexp.MustCompile(pattern)

	matches := regex.FindStringSubmatch(text)
	if len(matches) > 1 {
		return matches[1]
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

func trim(val []string) []string {
	var result []string
	for _, v := range val {
		result = append(result, strings.TrimSpace(v))
	}
	return result
}
