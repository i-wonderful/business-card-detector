package field_sort

import (
	. "card_detector/internal/service/field_sort/helper"
	"strings"
)

// checkAndFixUrls checks and fixes URLs based on emails.
//
// Parameters:
// - urls []string: slice of URLs to check and fix
// - emails []string: slice of emails used for checking and fixing URLs
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

// checkAndFixOrganization checks and fixes the organization name based on domain names.
//
// org *string - pointer to the organization name
// domainNames []string - slice of domain names
func (s *Service) checkAndFixOrganization(org *string, domainNames []string) {
	if len(*org) == 0 || len(domainNames) == 0 {
		return
	}
	domainName := domainNames[0]

	clearOrg := strings.Replace(strings.ToLower(*org), " ", "", -1)
	clearDomainName := strings.Replace(strings.ToLower(domainName), "-", "", -1)
	diff := LevenshteinDistance(clearOrg, clearDomainName)

	if diff == 0 || diff > 2 {
		return
	} else {
		*org = domainName
	}

	if diff < 3 {
		*org = domainName
	}
}
