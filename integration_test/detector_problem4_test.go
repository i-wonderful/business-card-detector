package integration_test

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

// --------------------------------------
// Test pack with some problem images (4)
// --------------------------------------

const BASE_IMG_PROBLEM4_PATH = "/home/olga/projects/card_detector_imgs/problem4"

func TestDetectProblem4(t *testing.T) {
	testCases := []struct {
		name     string
		imgPath  string
		expected *model.Person
	}{
		{
			"1.jpg",
			BASE_IMG_PROBLEM4_PATH + "/1.jpg",
			&model.Person{
				Email:        []string{"aiga@entez.com"},
				Site:         []string{"entez.com"},
				Phone:        []string{"+37129 146 960"},
				Name:         "Aiga Bunkse",
				Telegram:     []string{"@aigabunkse"},
				Organization: "", // todo vertical text entez
				JobTitle:     "Account Manager",
				Other:        "ZETNN;B10N",
			},
		},
	}

	detector, config := createDetector2(t)

	manage_file.ClearFolder(config.StorageFolder)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			actual, err := detector.Detect(tc.imgPath)

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
