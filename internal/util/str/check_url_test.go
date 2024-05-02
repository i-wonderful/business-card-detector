package manage_str

import "testing"

func TestIsValidURL(t *testing.T) {

	validUrls := []string{
		"https://www.google.com",
		"http://golang.org",
		"https://leetcode.com/problems",
		"google.com",
	}

	invalidUrls := []string{
		"http://",
		"ftp://invalid.com",
		": russell@connectingbrands.co.uk",
		"russell@connectingt",
	}

	for _, url := range validUrls {
		if !IsValidURL(url) {
			t.Errorf("Expected %s to be valid", url)
		}
	}

	for _, url := range invalidUrls {
		if IsValidURL(url) {
			t.Errorf("Expected %s to be invalid", url)
		}
	}

}

//func TestFindDomain(t *testing.T) {
//
//	urls := map[string]string{
//		"https://www.google.com":         "google.com",
//		"http://golang.org":              "golang.org",
//		"https://leetcode.com/problems/": "leetcode.com",
//	}
//
//	for url, expected := range urls {
//		actual := FindDomain(url)
//		if actual != expected {
//			t.Errorf("FindDomain(%s) == %s, expected %s", url, actual, expected)
//		}
//	}
//
//}
