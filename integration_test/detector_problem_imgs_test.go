package integration_test

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------
// Detect problematic images test
// ------------------------------------

const BASE_IMG_PROBLEM_PATH = "/home/olga/projects/card_detector_imgs/problem"

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
				Email:        []string{"mario.mouro@paylivre.com"},
				Site:         []string{"www.paylivre.com"},
				Phone:        []string{"+5511 99222-0597"},
				Name:         "MARIO MOURO",
				Organization: "CSBDO",
				JobTitle:     "",
			},
		},
		{
			"IMG_3623.JPG",
			BASE_IMG_PROBLEM_PATH + "/IMG_3623.jpg",
			&model.Person{
				Email:        []string{"gretta@endorphina.com"},
				Site:         []string{"endorphina.com"},
				Phone:        []string{"+420 222 564 222"},
				Skype:        []string{"gretta@endorphina.com"},
				Name:         "GRETTA KOCHKONYAN",
				Organization: "endorphina",
				JobTitle:     "Head Of Account Management",
				Other:        "",
			},
		},
		{
			"IMG_3587.jpg",
			BASE_IMG_PROBLEM_PATH + "/IMG_3587.jpg",
			&model.Person{
				Email:        []string{"flavio.tamega@upstreamsystems.com"},
				Site:         []string{},
				Phone:        []string{"+55 21 2146 0463", "+55 11 97278 5934"},
				Skype:        []string{"flavio.tamega"},
				Name:         "Flavio Tamega",
				Organization: "", // todo organization
				JobTitle:     "Advertising Commercial Director",
				Other:        "meeting оля pee И РС —",
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
