package integration_test

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

const BASE_IMG_PROBLEM5_PATH = "/home/olga/projects/card_detector_imgs/problem5"

func TestDetectProblem5(t *testing.T) {
	testCases := []struct {
		name     string
		imgPath  string
		expected *model.Person
	}{
		{
			// todo
			"0_3. Yummypay Danil.jpg",
			BASE_IMG_PROBLEM5_PATH + "/0_3. Yummypay Danil.jpg",
			&model.Person{
				Email:        []string{"contacts@yummypay.tech"},
				Skype:        []string{"YummyPsp@gmail.com"},
				Telegram:     []string{"@YummyPayDon"},
				Site:         []string{"yummypay.tech"},
				Name:         "Danil",
				Organization: "YummyPay",
				JobTitle:     "COO (Chief Operating Officer",
				Other:        "Ae-406sy.ygt.pap;5-59;mptp;skype:;email:",
			},
		},

		{
			"0_5. Newton CCO.jpg",
			BASE_IMG_PROBLEM5_PATH + "/0_5. Newton CCO.jpg",
			&model.Person{
				Email:    []string{"newton@pixtopay.com.br"},
				Phone:    []string{"+5541999197980"},
				Name:     "NEWTON AQUINO",
				JobTitle: "CCO",
				Other:    "Perl;Nioeria;Santa Catarina,Brazil",
			},
		},
		{
			"0_10. LPB Arkadijs.jpg",
			BASE_IMG_PROBLEM5_PATH + "/0_10. LPB Arkadijs.jpg",
			&model.Person{
				Email:        []string{"arkadijs.narcuks@lpb.lv"},
				Site:         []string{"www.lpb.lv"},
				Phone:        []string{"+37167772962", "+371)22352883"},
				Name:         "Arkadijs Narcuks",
				Organization: "LPB",
				JobTitle:     "E-komercijas departaments", // todo  E-komercijas departaments Klientu vaditajs
				Other:        "Klientu vaditajs;Brivibas iela 54.Riga.LV-1011",
			},
		},
		{
			"0_14. paytiqo.jpg",
			BASE_IMG_PROBLEM5_PATH + "/0_14. paytiqo.jpg",
			&model.Person{
				Email:    []string{"PAYTIQO@GMAIL.COM"},
				Phone:    []string{"+34628770939"},
				Site:     []string{"PAYFINANS.COM"},
				Name:     "PAYTIQO",
				Telegram: []string{"@DMYTRO1112", "@A_WSD"},
			},
		},
		{
			"0_18. gate2way.jpg",
			BASE_IMG_PROBLEM5_PATH + "/0_18. gate2way.jpg",
			&model.Person{
				Email:        []string{"john@gate2way.com"},
				Phone:        []string{"+35795 600889"},
				Name:         "lonut Paulenco", // todo Ionut Paulenco
				Organization: "gate2way",
				JobTitle:     "BDM",
				Other:        "Tnelil;",
			},
		},
		{
			"0_17. kirill.jpg",
			BASE_IMG_PROBLEM5_PATH + "/0_17. kirill.jpg",
			&model.Person{
				Email:        []string{"d.k@scipiopay.com"},
				Site:         []string{"scipiopay.com"},
				Phone:        []string{},
				Skype:        []string{},
				Telegram:     []string{"@KlrillMan"},
				Name:         "KIRILL",
				Organization: "",
				JobTitle:     "Business Developereyxcru",
				Other:        "4124.l;Cer;pumoUreoio;81.-2;hfpura;Web;adpTrepo",
			},
		},
	}

	detector, config := createDetector2(t)

	manage_file.ClearFolder(config.StorageFolder)
	manage_file.ClearFolder(config.TmpFolder)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//t.Parallel()
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
