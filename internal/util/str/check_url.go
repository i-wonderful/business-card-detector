package manage_str

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var tlds = []string{
	"com", "org", "net", "edu", "gov", "mil",
	"ac", "ad", "ae", "af", "ag", "ai" /*"al",*/, "am", "ao", "aq", "ar", "as", "at", "au", "aw", "ax", "az",
	"ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bm", "bn", "bo", "bq", "br", "bs", "bt", "bv", "bw", "by", "bz",
	"ca", "cc", "cd", "cf", "cg", "ch", "ci", "ck", "cl", "cm", "cn", "co", "consulting", "cr", "cu", "cv", "cw", "cx", "cy", "cz",
	"de", "dj", "dk", "dm", "do", "dz",
	"ec", "ee", "eg", "eh", "er", "es", "et", "eu",
	"fi", "fj", "fk", "fm", "fo", "fr",
	"ga", "game", "gb", "gd", "ge", "gf", "gg", "gh", "gi", "gl", "global", "gm", "gn", "gp", "gq", "gr", "gs", "gt", "gu", "gw", "gy",
	"hk", "hm", "hn", "hr", "ht", "hu",
	"id", "ie", "il", "im", "in", "io", "iq", "ir", "is", "it",
	"je", "jm", "jo", "jp",
	"ke", "kg", "kh", "ki", "km", "kn", "kp", "kr", "kw", "ky", "kz",
	"la", "lb", "lc", "li", "lk", "lr", "ls", "lt", "lu", "lv", "ly",
	"ma" /*"mc",*/, "md" /*"me",*/, "mf", "mg", "mh", "mk", "ml", "mm", "mn", "mo", "mp", "mq", "mr", "ms", "mt", "mu", "mv", "mw", "mx", "my", "mz", "media",
	"na", "nc", "ne", "nf", "ng", "ni", "nl", "no", "np", "nr", "nu", "nz",
	"om",
	"pa", "partners", "pe", "pf", "pg", "ph", "pk", "pl", "pm", "pn", "pr", "ps", "pt", "pw", "py",
	"qa",
	"re", "ro", "rs", "ru", "rw",
	"sa", "sb", "sc", "sd", "se", "sg", "sh", "si", "sj", "sk", "sl", "sm", "sn", "so", "sr", "ss", "st", "su", "sv", "sx", "sy", "sz",
	"tc", "td", "tf", "tg", "th", "tj", "tk", "tl", "tm", "tn", "to", "tr", "tt", "tv", "tw", "tz",
	"ua", "ug", "uk", "us", "uy", "uz",
	"va", "vc", "ve", "vg", "vi", "vn", "vu",
	"wf", "ws",
	"ye", //"yt",
	"za", "zm", "zw",
	"arpa",
}

func IsValidURL(u string) bool {
	// Add scheme if not provided for URL parser
	if !strings.HasPrefix(u, "http://") && !strings.HasPrefix(u, "https://") {
		u = "http://" + u
	}

	parsedURL, err := url.Parse(u)
	if err != nil || parsedURL.Host == "" {
		return false
	}
	const urlRegex = `^([\w-\.]+:\/\/)?(www\.)?([\w-\.]+)\.([\w-\.]+)\/?.+$`

	if matched, _ := regexp.MatchString(urlRegex, u); !matched {
		return false
	} else {
		return isValidDomain(parsedURL.Hostname())
	}

	return true
}

//
//func IsValidURL(url string) bool {
//	const urlRegex = `^([\w-\.]+:\/\/)?(www\.)?([\w-\.]+)\.([\w-\.+]+)\/?(\w+(-\w+)*)*`
//	matched, _ := regexp.MatchString(urlRegex, url)
//
//	if matched {
//		domain := FindDomain(url)
//		return domain != ""
//	}
//
//	return false
//}

func isValidDomain(domain string) bool {
	parts := strings.Split(domain, ".")
	tld := parts[len(parts)-1]
	tld = strings.ToLower(tld)

	for _, validTld := range tlds {
		if tld == validTld {
			return true
		}
	}

	return false
}

func CheckURLAvailability(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Ошибка при попытке подключиться к %s: %v\n", url, err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("URL доступен")
		return true
	} else {
		fmt.Printf("URL %s недоступен (код статуса: %d)\n", url, resp.StatusCode)
		return false
	}
}

// Check doest string is domain
//func IsDomain(domain string) bool {
//	regex := regexp.MustCompile(`(?:https?://)?([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)
//
//	return regex.MatchString(domain)
//
//}
