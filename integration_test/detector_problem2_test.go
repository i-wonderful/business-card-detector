package integration_test

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

const BASE_IMG_PROBLEM2_PATH = "/home/olga/projects/card_detector_imgs/problem2"

// --------------------------------------
// Test pack with some problem images (2)
// --------------------------------------
func TestDetectProblem2(t *testing.T) {
	testCases := []struct {
		name     string
		imgPath  string
		expected *model.Person
	}{
		{
			"1.jpg",
			BASE_IMG_PROBLEM2_PATH + "/1.jpg",
			&model.Person{
				Email:        []string{"u.sarper@tvbet.tv"},
				Phone:        []string{"+353870984819", "+380668352081"}, // todo "+380662252081"
				Skype:        []string{"live:.cid"},                      // todo live:cid.639e35052e7e9fe1
				Name:         "UTKU SARPER",
				Organization: "TVBET",
				JobTitle:     "Business Development Manager",
				Other:        "D-daryno;PayA+las;Onrcyes;uuop",
			},
		},
		{
			"3.jpg",
			BASE_IMG_PROBLEM2_PATH + "/3.jpg",
			&model.Person{
				Email:        []string{"kam@jeton.com"},
				Site:         []string{"www.jeton.com"},
				JobTitle:     "Key Account Manager",
				Organization: "Jeton",
				Name:         "KAM",
			},
		},
		{
			"4.jpg",
			BASE_IMG_PROBLEM2_PATH + "/4.jpg",
			&model.Person{
				Skype: []string{
					"live:cid.a53b3a75cd063b4a",
				},
				Telegram: []string{
					"@Eska8_Aff",
				},
			},
		},
		{
			"5.jpg",
			BASE_IMG_PROBLEM2_PATH + "/5.jpg",
			&model.Person{
				Email:    []string{"ponyango@pio.ke"},
				Site:     []string{"www.pio.ke"},
				Phone:    []string{"+254720961738", "+254113804990"},
				Name:     "PHILLIP ONYANGO",
				JobTitle: "MANAGING PARTNER",
				Other:    "4th Floor,The Westwood;Ring Road, Westlands;Nairobi, Kenya;licenges",
			},
		},
		{
			"6.jpg",
			BASE_IMG_PROBLEM2_PATH + "/6.jpg",
			&model.Person{
				Email:        []string{"INFO@HUGE.PARTNERS", "SUPPORT@HUGE.PARTNERS"},
				Organization: "HUGE",
				Name:         "GAMBLING BETTING",
				Other:        "FOR ADVERTISERS;FOR AFFILIATES",
			},
		},
		{
			"7.jpg",
			BASE_IMG_PROBLEM2_PATH + "/7.jpg",
			&model.Person{
				Email:        []string{"alex@softgamings.com"},
				Site:         []string{"www.SoftGamings.com"},
				Phone:        []string{"+37125371708", "+37125 155 112"},
				Skype:        []string{"alex.softgamings.com"},
				Name:         "Alexander Yerin",
				Organization: "SoftGamings",
				JobTitle:     "Head of Sales Department",
				Other:        "Brivibas 151,RigaLatvia,LV-1012",
			},
		},
		{
			"8.jpg",
			BASE_IMG_PROBLEM2_PATH + "/8.jpg",
			&model.Person{
				Email:    []string{"james.singer@3oaks.com"},
				Site:     []string{"www.3oaks.com"},
				Name:     "JAMES SINGER",
				JobTitle: "Senior Business Development Manager",
			},
		},
		{
			"9.jpg",
			BASE_IMG_PROBLEM2_PATH + "/9.jpg",
			&model.Person{
				Email:        []string{"slava@monotech.group"},
				Name:         "Slava Chernenko",
				Organization: "MMONOTECH",
				JobTitle:     "Senior Partnerships and Accounts Manager",
			},
		},
		{
			"10.jpg",
			BASE_IMG_PROBLEM2_PATH + "/10.jpg",
			&model.Person{
				Email:    []string{"laura.p@smartx.consulting"},
				Phone:    []string{"+59899700086"},
				Name:     "Laura Porto",
				JobTitle: "Lead Consultant-Brazil",
			},
		},
	}

	detector, config := createDetector2(t)

	manage_file.ClearFolder(config.StorageFolder)
	manage_file.ClearFolder(config.TmpFolder)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

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

func equalsListIgnoreOrder(expected []string, actual []string) bool {

	// implement todo

	return false
}
