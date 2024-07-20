package field_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
				"site":     []string{"endorphina.com"},
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
				"site":  []string{"www.igtrm.com"},
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
				"site":     []string{"www.369gaming.media"},
				//"skype":    "cid. 9e53d8c1 1 51546",
				"phone": []string{"+598 95 641 888"},
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
			"Doubled name",
			[]string{
				"Aiga Bunkse", "Aiga Bunkse",
			},
			map[string]interface{}{
				"name": "Aiga Bunkse",
			},
		},
		{
			"Contacted name",
			[]string{
				"GLENNDEBATTISTA",
			},
			map[string]interface{}{
				"name": "GLENN DEBATTISTA",
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
