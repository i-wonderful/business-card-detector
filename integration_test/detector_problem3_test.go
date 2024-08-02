package integration_test

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

const BASE_IMG_PROBLEM3_PATH = "/home/olga/projects/card_detector_imgs/problem3"

// --------------------------------------
// Test pack with some problem images (3)
// --------------------------------------
func TestDetectProblem3(t *testing.T) {
	testCases := []struct {
		name     string
		imgPath  string
		expected *model.Person
	}{
		{
			"1.jpg",
			BASE_IMG_PROBLEM3_PATH + "/1.jpg",
			&model.Person{
				Email:        []string{"u.sarper@tvbet.tv"},
				Phone:        []string{"+353870984819"}, // todo +38066 225 2081
				Skype:        []string{"live:cid.639e35052e7e9fe1"},
				Name:         "UTKUSARPER", // todo UTKU SARPER
				Organization: "TVBET",
				JobTitle:     "Business Development Manager",
				Other:        "D-daryna;PayA+las/co;Uuoponnatea",
			},
		},
		{
			"2.jpg",
			BASE_IMG_PROBLEM3_PATH + "/2.jpg",
			&model.Person{
				Email:        []string{"RG@entourage-global.com"},
				Phone:        []string{"+34699 86 1970"},
				Name:         "Ruben Guillem",
				Organization: "ENTOURAGE",
				JobTitle:     "Head of Sales",
				Other:        "SPORT &ENTERTAINMENT;Milan-London-Madrid;gosmmp",
			},
		},
		{
			"3.jpg",
			BASE_IMG_PROBLEM3_PATH + "/3.jpg",
			&model.Person{
				Email:    []string{"ADVERTISING@ADSCOMPASS.COM"},
				Skype:    []string{"ADVERTISING@ADSCOMPASS.COM"},
				Name:     "NATALIYA",
				Telegram: []string{"ADV_ADSCOMPASS"},
				JobTitle: "business development",
				Other:    "EMAIL;SKYPE;CPMC",
			},
		},
		{
			"4.jpg",
			BASE_IMG_PROBLEM3_PATH + "/4_cropped.jpg",
			&model.Person{
				Email:        []string{"info@ameegoclick.com"},
				Site:         []string{"www.ameegoclick.com"},
				Phone:        []string{"+918510904040"},
				Skype:        []string{"live:rajyousp"},
				Telegram:     []string{"@rajyousp"},
				Name:         "Rajiv Kumar",
				Organization: "AMEEGO CLICK",
				JobTitle:     "Growth Manager-Mobile and Web",
			},
		},
		{
			"5.jpg",
			BASE_IMG_PROBLEM3_PATH + "/5.jpg",
			&model.Person{
				Email:    []string{"marketing@myeventplanner.com.mt"},
				Phone:    []string{"+356 9903 6659"},
				Name:     "Anna Lubojatzka",
				JobTitle: "head of marketing",
				Other:    "Malta;60,The Offices;Triq Caravaggio;GharghurGHI602;Malta;Cyprus;6,Agias Marinas street;4044,Yermasogoia.Limossol;Cyprus;Mohammed Abdulla Bin Demaithan;Blldg No.ll;304.Damascus Street;Industrial Areo2 Al Quasis.Dubai;UAE"},
		},
	}

	detector, config := createDetector2(t)

	manage_file.ClearFolder(config.StorageFolder)
	manage_file.ClearFolder(config.TmpFolder)

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			t.Parallel()

			actual, _, err := detector.Detect(tc.imgPath)

			assert.NoError(t, err, "could not detect person")

			fillEmpty(tc.expected)
			assert.Equal(t, tc.expected.Name, actual.Name, "Name")
			assert.Equal(t, tc.expected.Email, actual.Email, "Email")
			equalIgnoreSpaces(t, tc.expected.Phone, actual.Phone, "Phone")
			assert.Equal(t, tc.expected.JobTitle, actual.JobTitle, "JobTitle")
			assert.Equal(t, tc.expected.Telegram, actual.Telegram, "Telegram")
			assert.Equal(t, tc.expected.Site, actual.Site, "Site")
			assert.Equal(t, tc.expected.Skype, actual.Skype, "Skype")
			assert.Equal(t, tc.expected.Organization, actual.Organization, "Organization")
		})
	}
}
