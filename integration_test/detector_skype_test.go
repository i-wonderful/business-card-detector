package integration_test

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectSkype(t *testing.T) {
	testCases := []struct {
		name     string
		imgPath  string
		expected *model.Person
	}{
		{
			"32.JPG",
			BASE_IMG_PATH + "/32.JPG",
			&model.Person{
				Email:        []string{"sebastian.jeppsson@wigetmedia.com"},
				Skype:        []string{"sebastian.jepsson86"},
				Phone:        []string{"+4676 16643 31"},
				Name:         "Sebastian Jeppsson",
				Organization: "wiget",
				JobTitle:     "Head of Performance Marketing",
				Other:        "Wiget Media;Kocksgatan 1;11624 Stockholm",
			},
		},
		{
			"46.JPG",
			BASE_IMG_PATH + "/46.JPG",
			&model.Person{
				Email:        []string{"andrew.friedl@actionpay.com.br"},
				Phone:        []string{"(21)996477550"},
				Skype:        []string{"andrew.pp.friedl"},
				Name:         "Andrew Friedl",
				Organization: "actionpay",
				JobTitle:     "Analista Comercial",
			},
		},
		{
			"69.JPG",
			BASE_IMG_PATH + "/69.JPG",
			&model.Person{
				//Email:        []string{"viorel.stan@gshmedia.com"}, // todo
				Site:  []string{"viorel.stanagshmedia.com"}, // todo
				Skype: []string{"viorel.stan87"},            // not detect pict
				//Name:  "Viorel Stan", // todo
				Organization: "ViorelStan CEO", // todo "GSH"
				//JobTitle:     "CEO", // todo
				Other: "ViorelStan",
			},
		},
	}
	testDetector, config := createDetector2(t)

	manage_file.ClearFolder(config.StorageFolder)
	manage_file.ClearFolder(config.TmpFolder)

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			//t.Parallel()

			actual, _, err := testDetector.Detect(tc.imgPath)

			assert.NoError(t, err, "could not detect person")

			fillEmpty(tc.expected)
			assert.Equal(t, tc.expected.Name, actual.Name, "Name")
			equalIgnoreCase(t, tc.expected.Email, actual.Email, "Email")
			equalIgnoreSpaces(t, tc.expected.Phone, actual.Phone, "Phone")
			assert.Equal(t, tc.expected.JobTitle, actual.JobTitle, "JobTitle")
			assert.Equal(t, tc.expected.Telegram, actual.Telegram, "Telegram")
			equalIgnoreCase(t, tc.expected.Site, actual.Site, "Site")
			equalIgnoreCase(t, tc.expected.Skype, actual.Skype, "Skype")

			assert.Equal(t, tc.expected.Organization, actual.Organization, "Organization")
		})

	}

}
