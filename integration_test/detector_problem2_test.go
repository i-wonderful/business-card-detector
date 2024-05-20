package integration_test

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

const BASE_IMG_PROBLEM2_PATH = "/home/olga/projects/card_detector_imgs/problem2"

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
				Email:        []string{"u.sarper@tvbet.ty"},
				Site:         []string{},
				Phone:        []string{"42306088", "+35 387 098 48 19"}, // todo
				Skype:        []string{},
				Telegram:     []string{},
				Name:         "UTKU SARPER",
				Organization: "Business Developmen",
				JobTitle:     "Manager",
				Other:        "",
			},
		},

		{
			"3.jpg",
			BASE_IMG_PROBLEM2_PATH + "/3.jpg",
			&model.Person{
				Email:    []string{"kam@jeton.com"},
				Site:     []string{"www.jeton.com"},
				JobTitle: "Key Account Manager",
				Other:    "Jeton",
			},
		},
		{
			"4.jpg",
			BASE_IMG_PROBLEM2_PATH + "/4.jpg",
			&model.Person{
				Skype: []string{
					//"live:.cid.a53b3a75cdo63b4a", // todo
				},
				Telegram: []string{
					"@Eskas_", // todo @Eskas_Aff
				},
			},
		},
		{
			"5.jpg",
			BASE_IMG_PROBLEM2_PATH + "/5.jpg",
			&model.Person{
				Email:        []string{"ponyango@pio.ke"},
				Site:         []string{"www.pio.ke"},
				Phone:        []string{"+254 720 961 738", "+254 113 804 990", " 773 - 00606"},
				Skype:        []string{},
				Telegram:     []string{},
				Name:         "Ring Road Westlands",
				Organization: "PHILLIP",
				JobTitle:     "",
				Other:        "4% Floor The Westwood;ONYANGO;Nairobi. Kenya",
			},
		},
		{ // lang detect rus todo
			"6.jpg",
			BASE_IMG_PROBLEM2_PATH + "/6.jpg",
			&model.Person{
				Email: []string{"SUPPORT@HUGE.PARTNERS"},
				//Site:         []string{"huge.partners"},
				Phone:        []string{},
				Skype:        []string{},
				Telegram:     []string{"@HUGEPARTNERS"},
				Name:         "FOR ADVERTISERS",
				Organization: "GAMBLING",
				JobTitle:     "",
				Other:        "cor ADVERTISERS",
			},
		},
		{
			"7.jpg",
			BASE_IMG_PROBLEM2_PATH + "/7.jpg",
			&model.Person{
				Email:        []string{"alex@softgamings.com"},
				Site:         []string{"alex.softgamings.com"},
				Phone:        []string{"+371 25 155112"},
				Skype:        []string{},
				Telegram:     []string{},
				Name:         "Alexander Yerin",
				Organization: "NWW.SOftGamings.com",
				JobTitle:     "Head of Sales Department —>Direct: +371 25 371 708",
				Other:        "",
			},
		},
		{
			"8.jpg",
			BASE_IMG_PROBLEM2_PATH + "/8.jpg",
			&model.Person{
				Email:        []string{"james.singer@3oaks.com"},
				Site:         []string{"www.3oaks.com"},
				Phone:        []string{},
				Skype:        []string{},
				Telegram:     []string{},
				Name:         "",
				Organization: "JAMES",
				JobTitle:     "Senior Business Development",
				Other:        "SINGER;Manager",
			},
		},
		{
			"9.jpg",
			BASE_IMG_PROBLEM2_PATH + "/9.jpg",
			&model.Person{
				Email:        []string{"slava@monotech.group"},
				Site:         []string{},
				Phone:        []string{},
				Skype:        []string{},
				Telegram:     []string{},
				Name:         "Slava Chernenko",
				Organization: "f MONOTECH",
				JobTitle:     "and Accounts Manager",
				Other:        "Senior Partnerships",
			},
		},
		{
			"10.jpg",
			BASE_IMG_PROBLEM2_PATH + "/10.jpg",
			&model.Person{
				Email:        []string{"p@smartx.consutting"},
				Site:         []string{},
				Phone:        []string{"+598 9 970 0086"}, // todo +598 9 970 0085
				Skype:        []string{},
				Telegram:     []string{},
				Name:         "Laura Porto",
				Organization: "",
				JobTitle:     "ead Consultant - Brazil",
				Other:        "i laura.",
			},
		},
	}

	detector, config := createDetector(t)

	manage_file.ClearFolder(config.StorageFolder)
	manage_file.ClearFolder("./tmp")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			actual, err := detector.Detect(tc.imgPath)

			assert.NoError(t, err, "could not detect person")

			fillEmpty(tc.expected)
			assert.Equal(t, tc.expected.Name, actual.Name, "Name")
			assert.Equal(t, tc.expected.Email, actual.Email, "Email")
			assert.Equal(t, tc.expected.Phone, actual.Phone, "Phone")
			assert.Equal(t, tc.expected.JobTitle, actual.JobTitle, "JobTitle")
			assert.Equal(t, tc.expected.Telegram, actual.Telegram, "Telegram")
			assert.Equal(t, tc.expected.Site, actual.Site, "Site")
			assert.Equal(t, tc.expected.Skype, actual.Skype, "Skype")
			assert.Equal(t, tc.expected.Organization, actual.Organization, "Organization")
		})
	}
}
