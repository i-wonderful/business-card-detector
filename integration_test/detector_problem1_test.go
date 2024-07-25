package integration_test

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

const BASE_IMG_PROBLEM_PATH = "/home/olga/projects/card_detector_imgs/problem"

// --------------------------------------
// Test pack with some problem images (1)
// --------------------------------------

func TestDetectProblem(t *testing.T) {
	testCases := []struct {
		name     string
		imgPath  string
		expected *model.Person
	}{
		{
			"74.JPG",
			BASE_IMG_PROBLEM_PATH + "/74.JPG",
			&model.Person{
				Email:    []string{"mario.mouro@paylivre.com"},
				Site:     []string{"www.paylivre.com"},
				Phone:    []string{"+5511 99222-0597"},
				Name:     "MARIO MOURO",
				JobTitle: "CSBDO",
			},
		},
		{
			"IMG_3623.JPG",
			BASE_IMG_PROBLEM_PATH + "/IMG_3623.jpg",
			&model.Person{
				Email:        []string{"gretta@endorphina.com"},
				Site:         []string{"endorphina.com"},
				Phone:        []string{"+420222564222"},
				Skype:        []string{"gretta@endorphina.com"},
				Name:         "GRETTA KOCHKONYAN",
				Organization: "endorphina",
				JobTitle:     "Head Of Account Management",
			},
		},
		{
			"IMG_3587.jpg",
			BASE_IMG_PROBLEM_PATH + "/IMG_3587.jpg",
			&model.Person{
				Email:        []string{"flavio.tamega@upstreamsystems.com"},
				Phone:        []string{"+5511972785934", "+55212146 0463"},
				Skype:        []string{"flavio.tamega"},
				Name:         "Flavio Tamega",
				Organization: "upstream ads",
				JobTitle:     "Advertising Commercial Director",
				Other:        "Sws Spy;Heuys",
			},
		},
		{
			"55.jpg",
			BASE_IMG_PATH + "/55.JPG",
			&model.Person{
				Email:    []string{"guilherme.rezende@nobrebet.com"},
				Site:     []string{"www.nobrebet.com"},
				Name:     "GUILHERME REZENDE",
				JobTitle: "LEAD ANALYST",
			},
		},
		{
			"75.JPG",
			BASE_IMG_PATH + "/75.JPG",
			&model.Person{
				Email:        []string{"OMgonzalez@sis.tv"},
				Site:         []string{"www.sis.tv"},
				Phone:        []string{"+34 658777836"},
				Name:         "Oscar Murueta",
				Organization: "SIS",
				JobTitle:     "",
				Other:        "Sports;Information;Services;Office;2 Whitehall Avenue;Kingston;Milton Keynes;MK10 ОАХ",
			},
		},
	}

	detector, config := createDetector2(t)

	manage_file.ClearFolder(config.StorageFolder)
	manage_file.ClearFolder(config.TmpFolder)

	for _, tc := range testCases {
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
