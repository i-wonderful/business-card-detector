package field_sort

import (
	"card_detector/internal/model"
	. "card_detector/internal/service/field_sort/helper"
	manage_file "card_detector/internal/util/file"
	manage_str "card_detector/internal/util/str"
	"log"
	"net/http"
	"sort"
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

func (s *Service) Sort(data []model.DetectWorld, boxes []model.TextArea) map[string]interface{} {
	if s.isLogTime {
		start := time.Now()
		defer func() {
			log.Printf(">>> Time sort fields: %s", time.Since(start))
		}()
	}

	var name, company, jobTitle, skype, address string
	var mailName, zone string
	domain := []string{}
	phones := []string{}
	emails := []string{}
	websites := []string{}
	telegram := []string{}

	notDetectItems := []model.DetectWorld{}

	recognized, indexesNotDetect, nameWorld := s.categorizeEvidentFields(data)
	data = keepIndexes(data, indexesNotDetect)

	data = s.categorizeByIcons(recognized, data, boxes)

	if _, ok := recognized[FIELD_TELEGRAM]; ok {
		telegram = recognized[FIELD_TELEGRAM].([]string)
	}
	if _, ok := recognized[FIELD_EMAIL]; ok {
		emails = recognized[FIELD_EMAIL].([]string)
		if len(emails) > 0 {
			var d string
			mailName, d, zone = ExtractMainNameDomainAndZone(emails[0])
			domain = append(domain, d)
		}
	}
	if _, ok := recognized[FIELD_SKYPE]; ok {
		skype = recognized[FIELD_SKYPE].(string)
	}
	if _, ok := recognized[FIELD_NAME]; ok {
		name = recognized[FIELD_NAME].(string)
	}

	if _, ok := recognized[FIELD_URLS]; ok {
		websites = recognized[FIELD_URLS].([]string)
		for _, site := range websites {
			d := ExtractDomainFromUrl(site)
			domain = append(domain, d)
		}
	}

	for _, word := range data {
		line := manage_str.ClearTrashSymbols(word.Text)

		if company == "" && s.processCompanyByDomain(line, domain, &company) {
			continue
		}

		if jobTitle == "" && s.processJobTitle(line, &jobTitle) {
			continue
		}

		if company == "" && manage_str.IsContains(line, s.companies) {
			company = line
		} else if address == "" && ContainsIgnoreCase(line, "address") {
			address = line
		} else {
			notDetectItems = append(notDetectItems, word)
		}
	}

	notDetectItems2 := []model.DetectWorld{}
	for _, word := range notDetectItems {
		item := strings.TrimSpace(word.Text)

		if len(emails) == 0 {
			if s := extractEmail(item); s != "" {
				emails = append(emails, s)
				continue
			}
		}
		if len(websites) == 0 {
			if IsSimpleUrlAndCheck(item) {
				websites = append(websites, item)
				continue
			}
			if s := ExtractBrokenUrl(item, domain, zone); s != "" {
				websites = append(websites, s)
				continue
			}
		}

		if len(telegram) == 0 {
			if s := extractTelegram(item); s != "" {
				telegram = append(telegram, s)
				continue
			}
		}
		if sk := ExtractSimpleSkype(skype, item); sk != "" {
			skype = sk
			continue
		}
		if name == "" {
			if m := nameRegex.MatchString(item); m {
				name = item
				nameWorld = word
				continue
			}
		}

		if s.processNameByMailName(item, mailName, &name) {
			nameWorld = word
			continue
		}

		if m := phoneRegexExtract.MatchString(item); m {
			phNew := phoneRegexExtract.FindAllString(item, -1)
			phones = append(phones, phNew...)
			continue
		}

		notDetectItems2 = append(notDetectItems2, word)
	}

	s.processSurnameIfSingleName(&name, &nameWorld, notDetectItems2)
	s.checkAndFixUrls(websites, emails)
	s.checkAndFixOrganization(&company, domain)

	person := recognized
	person[FIELD_EMAIL] = emails
	person[FIELD_PHONE] = append(person[FIELD_PHONE].([]string), clearPhones(phones)...)
	person[FIELD_NAME] = strings.TrimSpace(name)
	person[FIELD_SKYPE] = skype
	person["company"] = strings.TrimSpace(company)
	person[FIELD_TELEGRAM] = telegram
	person[FIELD_URLS] = websites
	person["jobTitle"] = strings.TrimSpace(jobTitle)
	person["other"] = strings.Join(GetOnlyWorlds(notDetectItems2), ";") + address

	return person
}

func (s *Service) processCompanyByDomain(line string, domains []string, company *string) bool {
	if CheckIsPartOfDomain(line, domains) && !IsSimpleUrl(line) {
		*company = line
		return true
	}
	return false
}

func (s *Service) processJobTitle(line string, jobTitle *string) bool {
	isFind, namesFind := isContainsManyWith(line, s.professions)
	if isFind {
		*jobTitle = line
		for _, n := range namesFind { // вставляем пробелы, если в должности они пропущены
			*jobTitle = InsertSpaceIfNeeded(*jobTitle, n)
		}
		return true
	}
	return false
}

func (s *Service) checkAndFixUrls(urls []string, emails []string) {
	if len(urls) == 0 || len(emails) == 0 {
		return
	}
	url := ExtractDomainAndZoneUrlFromUrl(urls[0])
	partsEmail := strings.Split(emails[0], "@")
	emailUrl := partsEmail[1]

	diff := LevenshteinDistance(url, emailUrl)
	if diff == 0 {
		return
	}
	if diff < 3 {
		urlOk := isSiteReachable(url)
		if urlOk {
			emails[0] = strings.Replace(emails[0], emailUrl, url, 1)
		} else if isSiteReachable(emailUrl) {
			urls[0] = strings.Replace(urls[0], url, emailUrl, 1)
		}
	}
}

func (s *Service) checkAndFixOrganization(org *string, domainNames []string) {
	if len(*org) == 0 || len(domainNames) == 0 {
		return
	}
	domainName := domainNames[0]

	clearOrg := strings.Replace(strings.ToLower(*org), " ", "", -1)
	clearDomainName := strings.Replace(strings.ToLower(domainName), "-", "", -1)
	diff := LevenshteinDistance(clearOrg, clearDomainName)

	if diff == 0 || diff > 3 {
		return
	} else {
		*org = domainName
	}

	if diff < 3 {
		*org = domainName
	}

}

func isSiteReachable(url string) bool {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func isContainsManyWith(s string, list []string) (bool, []string) {
	s = strings.ToLower(s)
	contains := make([]string, 0)
	isFind := false
	for _, p := range list {
		if strings.Contains(s, p) {
			contains = append(contains, p)
			isFind = true
		}
	}
	return isFind, contains
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
		if CountDigits(phone) > 8 {
			rez = append(rez, phone)
		}
	}

	return rez
}

func CheckIsPartOfDomain(item string, domains []string) bool {
	parts := strings.Split(item, " ")
	if len(parts) == 0 {
		return false
	}
	p0 := parts[0]

	for _, domain := range domains {
		if ContainsIgnoreCase(p0, domain) || ContainsIgnoreCase(domain, p0) {
			return true
		}
	}
	return false
}

func trim(val []string) []string {
	var result []string
	for _, v := range val {
		result = append(result, strings.TrimSpace(v))
	}
	return result
}

func keepIndexes(slice []model.DetectWorld, indexes []int) []model.DetectWorld {
	sort.Ints(indexes) // sort indexes in ascending order
	newSlice := make([]model.DetectWorld, 0)
	index := 0
	for i, v := range slice {
		if index < len(indexes) && indexes[index] == i {
			newSlice = append(newSlice, v)
			index++
		}
	}
	return newSlice
}

func remove(slice []model.DetectWorld, index int) []model.DetectWorld {
	return append(slice[:index], slice[index+1:]...)
}
