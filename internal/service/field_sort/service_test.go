package field_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractURL(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid URL",
			input:    "Visit https://example.com for more information.",
			expected: "example.com",
		},
		{
			name:     "URL without scheme",
			input:    "The site is example.org.",
			expected: "example.org",
		},
		{
			name:     "URL with path",
			input:    "Check out https://www.example.net/path/to/page",
			expected: "www.example.net",
		},
		{
			name:     "URL with query parameters",
			input:    "https://search.example.com?q=query&page=2",
			expected: "search.example.com",
		},
		{
			name:     "Invalid URL",
			input:    "This is not a URL.",
			expected: "",
		}, {
			name:     "without http or https",
			input:    "www.admill.io",
			expected: "www.admill.io",
		},
		{
			name:     "with site world",
			input:    "Site: Linebet.com Telegram:",
			expected: "Linebet.com",
		},
		{
			name:     "email not URL",
			input:    "Mail:  Skype: partners@Linebet.com",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := extractURL(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestExtractTelegram(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid telegram handle",
			input:    "My telegram handle is: Telegram: johndoe",
			expected: "johndoe",
		},
		{
			name:     "Valid telegram handle with whitespace",
			input:    "My telegram handle is: Telegram:   johndoe123",
			expected: "johndoe123",
		},
		{
			name:     "Case insensitive",
			input:    "My Telegram handle is: tELEGRAM: janedoe",
			expected: "janedoe",
		},
		{
			name:     "No telegram handle",
			input:    "This string does not contain a telegram handle",
			expected: "",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractTelegram(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, actual)
			}
		})
	}
}

func TestExtractSkype(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid Skype ID",
			input:    "My Skype: skype.sample_id",
			expected: "skype.sample_id",
		},
		//{
		//	name:     "Valid Skype ID with spaces",
		//	input:    "My Skype ID is: skype.live: sample_id",
		//	expected: "sample_id",
		//},
		{
			name:     "Valid Skype ID with case-insensitive pattern",
			input:    "My Skype: SAMPLE_ID",
			expected: "SAMPLE_ID",
		},
		{
			name:     "No Skype ID",
			input:    "This text does not contain",
			expected: "",
		},
		//{
		//	name:     "Multiple Skype IDs",
		//	input:    "My Skype IDs are: skype:id1 and skype:id2",
		//	expected: "id1",
		//},
		{
			name:     "Extract with @",
			input:    "Mail: b2b@lLinebet.com Skype: partners@Linebet.com",
			expected: "partners@Linebet.com",
		},
		{
			"Without colon",
			"Skype flavio.tamega",
			"flavio.tamega",
		},
		{
			"With two points",
			"Skype:alex.softgamings.com",
			"alex.softgamings.com",
		},
		//{
		//	"With point and colon",
		//	"live:.cid.e53090522ec2bf11",
		//	"live:.cid.e53090522ec2bf11",
		//},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := extractSkype(test.input)
			if result != test.expected {
				t.Errorf("Expected %q, but got %q", test.expected, result)
			}
		})
	}
}

func TestExtractEmail(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "With spaces",
			input:    "Martin@369gaming. media",
			expected: "Martin@369gaming.media",
		},
		//{
		//	"With 'E-mail'",
		//	"E-mail:christoffer.froberg@qpgames.se",
		//	"christoffer.froberg@qpgames.se",
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := extractEmail(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestService_Sort(t *testing.T) {

	testCases := []struct {
		name     string
		input    []string
		expected map[string]interface{}
	}{
		//{
		//	// len:3, cap:3
		//	name: "Find names",
		//	input: []string{
		//		"Site: Linebet.com Telegram: @linebet partners bot",
		//		"Mail: b2b@lLinebet.com Skype: partners@Linebet.com",
		//		"B2B Department",
		//	},
		//	expected: map[string]interface{}{
		//		//"site":     "Linebet.com", todo
		//		"telegram": "@linebet_partners_bot",
		//		"email":    []string{"b2b@lLinebet.com"},
		//		"skype":    "partners@Linebet.com",
		//		"company":  "B2B Department",
		//	},
		//},
		{
			name: "Email with space",
			input: []string{
				"Martin@369gaming. media",
			},
			expected: map[string]interface{}{
				"email": []string{"Martin@369gaming.media"},
			},
		},
		{
			name: "Email with 'E-mail'",
			input: []string{
				"E-mail: christoffer.froberg@qpgames.se",
			},
			expected: map[string]interface{}{
				"email": []string{"christoffer.froberg@qpgames.se"},
			},
		},
		{
			name: "Simple Email",
			input: []string{
				"ivk@colibrix.io",
			},
			expected: map[string]interface{}{
				"email": []string{"ivk@colibrix.io"},
			},
		},
		{
			"Simple site",
			[]string{
				"www.tornado-games.com",
			},
			map[string]interface{}{
				"site": "www.tornado-games.com",
			},
		}, {
			name: "Complex",
			input: []string{
				"-endorphina",
				"GRETTA KOCHKONYAN",
				"Head Of Account Management",
				"5 gretta@endorphina.com",
				"5) gretta@endorphina.com",
				"+420 222 564 222",
				"‚ endorphina.com",
			},
			expected: map[string]interface{}{
				"name":     "GRETTA KOCHKONYAN",
				"jobTitle": "Head Of Account Management",
				"company":  "endorphina",
				"email":    []string{"gretta@endorphina.com"},
				"skype":    "gretta@endorphina.com",
				"phone":    []string{"+420 222 564 222"},
				"site":     "endorphina.com",
			},
		},
		{
			"Complex2",
			[]string{
				"Areg Oganesian",
				"www.igtrm.com",
				"+374 99 452772",
			},
			map[string]interface{}{
				"name":  "Areg Oganesian",
				"site":  "www.igtrm.com",
				"phone": []string{"+374 99 452772"},
			},
		},
		{
			"Company",
			[]string{
				"Payment.Center",
			},
			map[string]interface{}{
				"company": "Payment.Center",
			},
		},
		{
			"Complex3",
			[]string{
				"GAMING",
				", www.369gaming media",
				"VRID",
				"Martin Buero",
				"General Manager",
				": +598 95 641 888",
				"Martin@369gaming.media",
				"Skype live: cid. 9e53d8c1 1 51546",
			},
			map[string]interface{}{
				"name":     "Martin Buero",
				"jobTitle": "General Manager",
				"company":  "GAMING",
				"email":    []string{"Martin@369gaming.media"},
				"site":     "www.369gaming.media",
				//"skype":    "cid. 9e53d8c1 1 51546",
				"phone": []string{"+598 95 641 888"},
			},
		},
		{
			"Phones",
			[]string{
				"+7 (473) 200-0-300 © ",
				"+7 (473) 20-20-457 (©",
				"773-00606",
			},
			map[string]interface{}{
				"phone": []string{
					"+7 (473) 200-0-300",
					"+7 (473) 20-20-457",
				},
			},
		},
		{
			"Skype",
			[]string{
				"Skype live.cid.9e53d8c1151b4b",
			},
			map[string]interface{}{
				"skype": "live.cid.9e53d8c1151b4b",
			},
		},
		{
			"Skype2",
			[]string{
				"s:live:cid.639e35052e7e9fe1 bla",
			},
			map[string]interface{}{
				"skype": "live:cid.639e35052e7e9fe1",
			},
		},
		{
			"Skype with two points",
			[]string{
				"Skype:alex.softgamings.com",
			},
			map[string]interface{}{
				"skype": "alex.softgamings.com",
			},
		},
		{
			"Skype with point and semicolon",
			[]string{"live:.cid.e53090522ec2bf11"},
			map[string]interface{}{
				"skype": "live:.cid.e53090522ec2bf11",
			},
		},
		{
			"Organization",
			[]string{
				"Manager", "TVBET", "+353870984819", "live:cid.639e35052e7e9fe1", "Business Development", "u.sarper@tvbet.tv",
			},
			map[string]interface{}{
				"company": "TVBET",
			},
		},
		{
			"Simple Name",
			[]string{
				"Jeton", "KAM", "kam@jeton.com", "www.jeton.com",
			},
			map[string]interface{}{
				"name":    "KAM",
				"company": "Jeton",
			},
		},
		{
			"Telegram with spaces",
			[]string{"@Eska8 Aff"},
			map[string]interface{}{
				"telegram": "@Eska8_Aff",
			},
		},
		{
			"Many emails",
			[]string{
				"SUPPORT@HUGE.PARTNERS",
				"INFO@HUGE.PARTNERS",
			},
			map[string]interface{}{
				"email": []string{
					"SUPPORT@HUGE.PARTNERS",
					"INFO@HUGE.PARTNERS",
				},
			},
		},
		{
			"Long job title",
			[]string{
				"Slava Chernenko",
				"Senior Partnerships",
				"and Accounts Manager",
			},
			map[string]interface{}{
				"jobTitle": "Senior Partnerships and Accounts Manager",
			},
		},
		{
			"Job title with defines",
			[]string{"Lead Consultant-Brazil"},
			map[string]interface{}{
				"jobTitle": "Lead Consultant-Brazil",
			},
		},
		{
			"Telegram by url",
			[]string{
				"https://t.me/Nicola_an",
				"www.admillio",
				"erkin@admillio",
			},
			map[string]interface{}{
				"telegram": "https://t.me/Nicola_an",
			},
		},
	}

	service := NewService(
		"../../../config/professions.txt",
		"../../../config/companies.txt",
		"../../../config/names.txt",
		true)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := service.Sort(tc.input)
			for k, v := range tc.expected {
				assert.Equal(t, v, result[k])
			}
		})
	}
}
