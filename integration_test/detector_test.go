package integration_test

import (
	"card_detector/internal/app"
	"card_detector/internal/model"
	"card_detector/internal/repo/inmemory"
	"card_detector/internal/service"
	"card_detector/internal/service/field_sort"
	"card_detector/internal/service/img_prepare"
	"card_detector/internal/service/text_find/onnx"
	"card_detector/internal/service/text_recognize"
	manage_file "card_detector/internal/util/file"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const BASE_IMG_PATH = "/home/olga/projects/card_detector_imgs"

/*
22.JPG ok
3.JPG ok
2.JPG ok
53.JPG ok
9.JPG ok
first IMG_2912.JPG ok
first IMG_2913.JPG ok
first IMG_2926.JPG ok
18.JPG ok
first IMG_2915.JPG ok not all fields
first IMG_2918.JPG ok
*/
func TestDetect(t *testing.T) {
	testCases := []struct {
		name     string
		imgPath  string
		expected *model.Person
	}{
		{
			"first IMG_2912.JPG",
			BASE_IMG_PATH + "/first/IMG_2912.jpg",
			&model.Person{
				Email:        []string{"oren@delasport.com"},
				Site:         []string{"www.delasport.com"},
				Phone:        []string{"+556 99725"}, // todo "+356 99723767"
				Name:         "OREN COHEN SHWARTZ",
				Organization: "WE DELIVER SPORTS",
				JobTitle:     "", // todo CEO
				Other:        "©Delasport",
			},
		},
		{
			"first IMG_2913.JPG",
			BASE_IMG_PATH + "/first/IMG_2913.jpg",
			&model.Person{
				Email:        []string{"craig@summer-creative.com"},
				Phone:        []string{"+34 711 017 134", "+44 (0)20 3355 5336"},
				Name:         "Craig Edwards",
				Organization: "SUMMER",
				JobTitle:     "Business Development Manager",
				Other:        "",
			},
		},
		{
			"first IMG_2914.JPG",
			BASE_IMG_PATH + "/first/IMG_2914.JPG",
			&model.Person{
				//Name: "Areg Oganesian", // todo
				Site:         []string{"www.igtrm.com"},
				Email:        []string{"areg@igtrm.com"},
				Phone:        []string{"+374 99 452772"},
				Organization: "",
				JobTitle:     "", //todo "CEO",
			},
		},
		{
			"first IMG_2915.JPG",
			BASE_IMG_PATH + "/first/IMG_2915.JPG",
			&model.Person{
				Name:     "Nadezda Tereshchenko",
				Email:    []string{"info@platinum-expo.com"},
				JobTitle: "Creative director",
				Site:     []string{"platinum-expo.com"},
				//Telegram:     []string{"@Naddiz_T"}, todo
				Organization: "PLATINUM",
			},
		},
		//	{
		//		"./data/IMG_2916.JPG",
		//		&model.Person{
		//			Email: "info@revpanda.com",
		//		},
		//	},
		{
			"first IMG_2917.JPG",
			BASE_IMG_PATH + "/first/IMG_2917.JPG",
			&model.Person{
				Email:        []string{"b2b@linebet.com"},
				Skype:        []string{"partners@Linebet.com"},
				Telegram:     []string{"@linebet"}, // todo wrong @linebet_partners_bot
				Site:         []string{"Linebet.com"},
				Organization: "B2B Department",
			},
		},
		{
			"first IMG_2918.JPG",
			BASE_IMG_PATH + "/first/IMG_2918.JPG",
			&model.Person{
				Name:     "Aron Myerthall",
				Email:    []string{"aron@raventrack.com"},
				Phone:    []string{"07956 710535"},
				JobTitle: "Sales Manager",
				//Organization: "-R/VENTRACK",
				Other: "—R/VENTRACK —",
			},
		},
		{
			"first IMG_2919.JPG",
			BASE_IMG_PATH + "/first/IMG_2919.JPG",
			&model.Person{
				Name:         "Jozef Fabian",
				Email:        []string{"jf@sportsinnovation.dk"},
				Phone:        []string{"+45 52 22 41 50"},
				JobTitle:     "HEAD OF CLIENT SUCCESS",
				Site:         []string{"www.sportsinnovation.dk"},
				Skype:        []string{"livejof_144"}, // todo wrong
				Organization: "SPORTS",
			},
		},
		//	{
		//		"./data/IMG_2920.JPG",
		//		&model.Person{
		//			Email:    "partner@coins.game",
		//			Telegram: "@cg_partners",
		//			Site:     "coins.game",
		//			Skype:    "coinsgame.partners",
		//		},
		//	},
		//	{
		//		"./data/IMG_2921.JPG",
		//		&model.Person{
		//			Name:     "DEAN RAYSON",
		//			Email:    "dean@all-in.global",
		//			Phone:    "+351 220 991 583",
		//			JobTitle: "HEAD OF SALES",
		//			Site:     "all-in.global",
		//			Other:    "MT Office +356 770 408 06;UK Mobile +44 7921 239 788",
		//		},
		//	},
		//	{
		//		"./data/IMG_2922.JPG",
		//		&model.Person{
		//			Name:     "Russ Yershon",
		//			Email:    "russell@connectingbrands.co.uk",
		//			Phone:    "+44 (0) 7500828120",
		//			JobTitle: "Director",
		//			Site:     "Connectingbrands.co.uk",
		//			Organization:  "CONNECTING", // todo CONNECTING BRANDS .co.ux
		//			Skype:    "russ.yershon",
		//			Telegram: "@connectingt",
		//		},
		//	},
		//	{
		//		"./data/IMG_2923.JPG",
		//		&model.Person{
		//			Name:     "Milda Grigaliunaite",
		//			Email:    "milda@world-card.com",
		//			Phone:    "+ 44 20 3835 6450",
		//			JobTitle: "Backoffice Supervisor",
		//			Organization:  "WsrldCard",
		//		},
		//	},
		//	{
		//		"./data/IMG_2924.JPG",
		//		&model.Person{
		//			Name:     "Naman Sharma",
		//			Email:    "naman.sharma@phoenixtech.consulting",
		//			Phone:    "+91-8607860787",
		//			JobTitle: "(Business Development Manager)",
		//			Site:     "https://phoenixtech.consulting",
		//			Organization:  "Phoenix Tech Consulting",
		//			Other:    "Haryana 122001;~hoenix;147 Vipul Trade Center, Sohna -;Pvt.Ltd;Road, Sector 48, Gurugram,",
		//		},
		//	},
		//	{
		//		"./data/IMG_2925.JPG",
		//		&model.Person{
		//			Name:     "Karina",
		//			Email:    "karen.s@nowpayments.io",
		//			JobTitle: "Business Development Manager",
		//			Site:     "nowpayments.io",
		//			Organization:  "Payments",
		//			Other:    "Seven Oo FAY WENT G;Йо PAYMENT GATE WR",
		//		},
		//	},
		{
			"first IMG_2926.JPG",
			BASE_IMG_PATH + "/first/IMG_2926.JPG",
			&model.Person{
				//Name:         "Dariya Yeryomenko",  // todo
				Email:    []string{"dariya@pay.center"},
				Phone:    []string{"+357 963 341 18"},
				JobTitle: "Key Account Manager",
				//Organization: "Payment.Center",  // todo
				Telegram: []string{"@dariya_pc_cy"},
			},
		},
		{
			"3648.jpg",
			BASE_IMG_PATH + "/IMG_3648.jpg",
			&model.Person{
				Name:         "Dariya Yeryomenko",
				Email:        []string{}, // todo
				Phone:        []string{"+357 963 341 18"},
				JobTitle:     "Key Account Manager",
				Organization: "dariya@pay-center",
				Telegram:     []string{"@dariya_pc_cy"},
			},
		},
		//	{
		//		"./data/IMG_2927.JPG",
		//		&model.Person{
		//			Name:    "Contact", // todo wrong
		//			Email:   "business@libernetix.com",
		//			Site:    "libernetix.com",
		//			Organization: "libernetix team",
		//			Other:   "and partnerships.;t us for collaborations",
		//		},
		//	},
		//	{
		//		"./data/IMG_2928.JPG",
		//		&model.Person{
		//			Name:     "Kate Elterova",
		//			Email:    "kate@oddsdigger.com",
		//			Telegram: "@KateElterova",
		//			JobTitle: "AFFILIATE MANAGER",
		//			Skype:    "elterova.ekaterina",
		//			Organization:  "DEV", // todo wrong
		//		},
		//	},
		//	{
		//		"./data/rus_example.jpg",
		//		&model.Person{
		//			Email: "unstroy444@mail.ru",
		//			Site:  "gk-grishkino.ru",
		//			Phone: "8-910-539-84-44",
		//			Name:  "Валерий Викторови Некрылов",
		//			Other: "из ЛСТК;||роектирование;производство;Строительство;< |",
		//		},
		//	},
		{
			name:    "test 2.JPG",
			imgPath: "/home/olga/projects/card_detector_imgs/2.JPG",
			expected: &model.Person{
				Name:         "Vladyslav Kolodistyi",
				Email:        []string{},
				Phone:        []string{},
				JobTitle:     "Chief Executive Officer",
				Organization: "Y payadmit",
				Other:        "Smart Technology Payment Solution;» White Label Gateway;Cashier Service;Middleware",
			},
		},
		{
			name:    "test 3.JPG",
			imgPath: "/home/olga/projects/card_detector_imgs/3.JPG",
			expected: &model.Person{
				Email:        []string{"filip.kisala@zen.com"},
				Site:         []string{},
				Phone:        []string{"+48 693 782 997"},
				Skype:        []string{},
				Telegram:     []string{},
				Name:         "Filip Kisata",
				Organization: "",
				JobTitle:     "Enterprise Sales Manager",
				Other:        "",
			},
		},
		{
			name:    "test 22.JPG",
			imgPath: "/home/olga/projects/card_detector_imgs/22.JPG",
			expected: &model.Person{
				Email:        []string{"erkin@admill.io"},
				Site:         []string{"www.admill.io"},
				Phone:        []string{"+90 536 745 13 03", "15551"}, // todo
				Skype:        []string{},
				Telegram:     []string{"https://t.me/Nicola_an"},
				Name:         "Erkin Bayrakcl",
				Organization: "",
				JobTitle:     "",
				Other:        "Sepapaja tn 6, 15551, Tallinn, Estonia",
			},
		},
		{
			name:    "test 53.JPG",
			imgPath: "/home/olga/projects/card_detector_imgs/53.JPG",
			expected: &model.Person{
				Email:        []string{"agustin.perez-vernet@gamelounge.com"},
				Site:         []string{"gamelounge.com"},
				Phone:        []string{},
				Skype:        []string{},
				Telegram:     []string{},
				Name:         "Agustin",
				Organization: "",
				JobTitle:     "Casino Site Manager",
				Other:        "Perez-Vernet",
			},
		},
		{
			name:    "test 9.JPG",
			imgPath: "/home/olga/projects/card_detector_imgs/9.JPG",

			expected: &model.Person{
				Email:        []string{"ivk@colibrix.io"},
				Site:         []string{"www.colibrix.io"},
				Phone:        []string{},
				Skype:        []string{},
				Telegram:     []string{"@lveta_IVK"},
				Name:         "IVETA KRUMINA",
				Organization: "Payment Solutions",
				JobTitle:     "Head of Sales",
				Other:        "",
			},
		},

		{
			"test 18.JPG",
			"/home/olga/projects/card_detector_imgs/18.JPG",
			&model.Person{
				Email:        []string{"emma@internationaladvertisingsolutions.com"},
				Site:         []string{},
				Phone:        []string{"+44 7885 723 853"},
				Skype:        []string{},
				Telegram:     []string{},
				Name:         "Emma Fisher",
				Organization: "Jip ADVERTISING SOLUTIONS",
				JobTitle:     "Account Director",
				Other:        "& INTERNATIONAL;International Media Planning & Buying Services:;° Sponsorships;‚ Radio & Podcasts;e Innovative Solutions;Emma Fisher",
			},
		},
		{
			"4328.JPG",
			BASE_IMG_PATH + "/IMG_4328.jpg",
			&model.Person{
				Email: []string{"Martin@369gaming.media"},
				Site:  []string{"369gaming.media"},
				Phone: []string{"+598 95 641 888"},
				//	Skype:        []string{"cid"}, // todo
				Telegram:     []string{},
				Name:         "Martin Buero",
				Organization: "GAMING",
				JobTitle:     "General Manager",
				Other:        "",
			},
		},
	}

	testDetector, config := createDetector(t)

	manage_file.ClearFolder(config.StorageFolder)
	manage_file.ClearFolder("./tmp")
	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			actual, err := testDetector.Detect(tc.imgPath)

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

func createDetector(t *testing.T) (*service.Detector, *app.AppConfig) {

	os.Setenv("CONFIG_FILE", "./config/config.yml")
	isLogTime := true

	config, err := app.NewConfigFromYml()
	if err != nil {
		t.Fatal(err)
	}

	cardRepo := inmemory.NewCardRepo()
	imgPreparer := img_prepare.NewService(config.StorageFolder)
	findTextService, err := onnx.NewService(config.Onnx.PathRuntime, config.Onnx.PathModel, isLogTime)
	if err != nil {
		t.Fatal(err)
	}
	//findTextService := remote.NewFindTextService()
	textRecognizer := text_recognize.NewService(isLogTime)
	fieldSorter := field_sort.NewService(config.PathProfessionList, config.PathCompanyList, config.PathNamesList, isLogTime)

	// detector
	testDetector := service.NewDetector(
		imgPreparer,
		findTextService,
		textRecognizer,
		fieldSorter,
		cardRepo,
		config.StorageFolder,
		isLogTime)

	return testDetector, config
}

func fillEmpty(p *model.Person) {
	if p.Skype == nil {
		p.Skype = []string{}
	}
	if p.Telegram == nil {
		p.Telegram = []string{}
	}
	if p.Site == nil {
		p.Site = []string{}
	}
	if p.Email == nil {
		p.Email = []string{}
	}
	if p.Phone == nil {
		p.Phone = []string{}
	}
}
