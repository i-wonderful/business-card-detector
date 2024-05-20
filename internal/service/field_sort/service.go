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

func (d *Service) Sort(data []string) map[string]interface{} {
	if d.isLogTime {
		start := time.Now()
		defer func() {
			log.Printf(">>> Time sort fields: %s", time.Since(start))
		}()
	}

	phonePattern := `[\+\(]?[0-9 .\(\)-]{7,}`
	//websitePattern := `^[a-zA-Z0-9\.-]+\.[a-zA-Z]{2,}$`
	//addressPattern := `^[\w\s\p{L},0-9()/\\-]*$`

	//emaiNewRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,11}`)
	//emailNewRegex := regexp.MustCompile(`(?:Mail:\s*)?([\S\s]+?@[\S\s]+?\.\S+)\b`)
	emailNewRegex := regexp.MustCompile(`(?i)(?:Mail:\s*)?([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`)

	//regexp.MustCompile(`[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+`) // email
	phoneRegex := regexp.MustCompile(phonePattern)
	//websiteRegex := regexp.MustCompile(websitePattern)
	//addressRegex := regexp.MustCompile(addressPattern)
	skypeRegex := regexp.MustCompile(`^[a-z]{4,31}\.[a-z]{4,31}`)
	companyRegex := regexp.MustCompile(`^[A-Z]{3,}[ \.]*$|^[A-Z]+$`)
	nameRegex := regexp.MustCompile(`^[A-ZА-Я][A-ZА-Яa-zа-я-]+ [A-ZА-Я][A-ZА-Яa-zа-я-]+([ \-][A-ZА-Я][A-ZА-Яa-zа-я-]+)?$`)
	//singleName := regexp.MustCompile("^[А-ЯA-Z][а-яa-z]{1,}$")
	telegramRegex := regexp.MustCompile(`(?:https?://t\.me/|@)[A-Za-z][A-Za-z0-9_]{4,31}`)

	var email, name, company, jobTitle, telegram, website, skype, domain, zone string
	phones := []string{}

	person := map[string]interface{}{}
	notDetectItems := []string{}

	for _, line := range data {
		line = clearTrashSymbols(line)

		if match := emailNewRegex.FindStringSubmatch(line); len(match) > 1 {
			findEmail := strings.Replace(match[1], " ", "", -1)
			if email == "" {
				email = findEmail
			} else if skype == "" { // скайп может совпадать с email
				skype = findEmail
			}

			// domain may be company
			domain, zone = extractDomainAndZone(findEmail)
			d.companies = append(d.companies, domain)

			line = strings.ReplaceAll(line, match[1], "")
			if len(line) > 5 {
				notDetectItems = append(notDetectItems, line)
			}
		} else if jobTitle == "" && isContains(line, d.professions) {
			jobTitle = line
		} else if match := telegramRegex.FindString(line); match != "" {
			telegram = match
			line = strings.ReplaceAll(line, match, "")
			if len(line) > 5 {
				notDetectItems = append(notDetectItems, line)
			}
		} else if w := extractURL(line); w != "" && website == "" { // todo уменьшить вызовы
			website = w
		} else if company == "" && isContains(line, d.companies) {
			company = line
		} else if name == "" && isContainsWithSpace(line, d.names) {
			name = line
		} else if len(phones) == 0 && phoneRegex.MatchString(line) {
			phones = phoneRegex.FindAllString(line, -1)
		} else if skype == "" && skypeRegex.MatchString(line) {
			skype = line
		} else {
			// If no pattern matched, consider it as part of the name or other descriptions
			//other = fmt.Sprintf("%s %s", name, line)
			notDetectItems = append(notDetectItems, line)
		}
	}

	other := []string{}
	for _, item := range notDetectItems {
		item = strings.TrimSpace(item)

		if email == "" {
			if s := extractEmail(item); s != "" {
				email = s
				continue
			}
		}
		if website == "" {
			if s := extractBrokenUrl(item, domain, zone); s != "" {
				website = s
				continue
			}
		}
		if company == "" {
			if domain != "" && strings.Contains(item, domain) {
				company = domain
				continue
			}
			if math := companyRegex.MatchString(item); math {
				company += " " + item
				continue
			}
		}
		if skype == "" {
			if s := extractSkype(item); s != "" {
				skype = s
				continue
			}
		}

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

		if m := phoneRegex.MatchString(item); m {
			phNew := phoneRegex.FindAllString(item, -1)
			phones = append(phones, phNew...)
			continue
		}

		other = append(other, item)
	}

	person["email"] = email
	person["phone"] = clearPhones(phones)
	person["name"] = strings.TrimSpace(name)
	person["skype"] = skype
	person["company"] = strings.TrimSpace(company)
	person["telegram"] = telegram
	person["site"] = website
	person["jobTitle"] = jobTitle
	person["other"] = strings.Join(other, ";")

	return person
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
			}
			return false
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

// ео\nSkype live:.cid.9e53d8c1 151646
func extractSkype(text string) string {
	pattern := `(?i)(?:skype\s*(?:live\s*)?[:\.]\s*|skype:\s*)(\w+(?:@\w+\.[\w.]+)?)|(?i)skype\s+(\w+\.\w+)`
	re := regexp.MustCompile(pattern)
	subs := re.FindStringSubmatch(text)
	if len(subs) > 1 {
		if subs[1] != "" {
			return subs[1]
		} else if subs[2] != "" {
			return subs[2]
		}
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

func extractDomainAndZone(email string) (string, string) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", "" // Неверный формат адреса электронной почты
	}
	domainParts := strings.Split(parts[1], ".")
	if len(domainParts) < 2 {
		return "", "" // Неверный формат домена
	}
	return domainParts[0], domainParts[1]
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

	return phones
}

func trim(val []string) []string {
	var result []string
	for _, v := range val {
		result = append(result, strings.TrimSpace(v))
	}
	return result
}
