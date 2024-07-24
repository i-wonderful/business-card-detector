package integration_test

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectTelegram(t *testing.T) {
	testCases := []struct {
		name     string
		imgPath  string
		expected *model.Person
	}{
		{
			"0_13. hanna omelchak.jpg",
			BASE_IMG_PROBLEM5_PATH + "/0_13. hanna omelchak.jpg",
			&model.Person{
				Email:        []string{"hanna.m@paycord.com"},
				Phone:        []string{"+351934021737"},
				Telegram:     []string{"hanna_biz"},
				Name:         "Hanna Omelchak",
				Organization: "paycord",
				JobTitle:     "Business Development Manager",
				Other:        "Tansfe",
			},
		},
		{
			"1_23 german magnus.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_23 german magnus.jpg",
			&model.Person{
				Email:    []string{"german_cvo@expay.cash"},
				Telegram: []string{"magnus_exp"},
				Name:     "GermanS",
				JobTitle: "CVO Expay",
				Other:    "GermanS;NFC))",
			},
		},
		{
			"1_40 maks paycord.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_40 maks paycord.jpg",
			&model.Person{
				Email: []string{"ms@paycord.com"},
				//Telegram: []string{"m_pkr"}, //todo
				//Name:         "", // todo
				Organization: "paycord",
				JobTitle:     "Chief Operating Officer",
				Other:        "Maks S",
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
