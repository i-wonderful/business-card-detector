package integration_test

import (
	"card_detector/internal/app"
	"card_detector/internal/model"
	"card_detector/internal/repo/inmemory"
	"card_detector/internal/service"
	"card_detector/internal/service/detect/onnx"
	"card_detector/internal/service/field_sort"
	"card_detector/internal/service/img_prepare"
	"card_detector/internal/service/text_recognize/paddleocr"
	manage_file "card_detector/internal/util/file"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

const BASE_IMG_PATH = "/home/olga/projects/card_detector_imgs"

// ---------------------
// Main test pack
// ---------------------
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
				Phone:        []string{"+35699723767"},
				Name:         "OREN COHEN SHWARTZ",
				Organization: "WEDELIVERSPORTS", // todo spaces
				JobTitle:     "CEO",
				Other:        "©Delasport",
			},
		},
		{
			"first IMG_2913.JPG",
			BASE_IMG_PATH + "/first/IMG_2913.jpg",
			&model.Person{
				Email:        []string{"craig@summer-creative.com"},
				Phone:        []string{"+34711017134", "+4402033555336"},
				Name:         "Craig Edwards",
				Organization: "SUMMER CREATIVE",
				JobTitle:     "Business Development Manager",
				Other:        "",
			},
		},
		{
			"first IMG_2914.JPG",
			BASE_IMG_PATH + "/first/IMG_2914.JPG",
			&model.Person{
				Name:         "Areg", // todo Areg Oganesian
				Site:         []string{"www.igtrm.com"},
				Email:        []string{"areg@igtrm.com"},
				Phone:        []string{"+37499452772"},
				Organization: "iGTRM",
				JobTitle:     "CEO",
			},
		},
		{
			"first IMG_2915.JPG",
			BASE_IMG_PATH + "/first/IMG_2915.JPG",
			&model.Person{
				Name:         "Nadezda Tereshchenko",
				Email:        []string{"Info@platinum-expo.com"},
				JobTitle:     "Creative director",
				Site:         []string{"platinum-expo.com"},
				Telegram:     []string{"@Naddiz_T"},
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
				Skype:        []string{"partners@linebet.com"},
				Telegram:     []string{"@linebet_partners_bot"},
				Site:         []string{"linebet.com"},
				Organization: "B2B Department",
			},
		},
		{
			"first IMG_2918.JPG",
			BASE_IMG_PATH + "/first/IMG_2918.JPG",
			&model.Person{
				Name:         "Aron Myerthall",
				Email:        []string{"aron@raventrack.com"},
				Phone:        []string{"07956710535"},
				JobTitle:     "Sales Manager",
				Organization: "RAVENTRACK",
				Other:        "",
			},
		},
		{
			"first IMG_2919.JPG",
			BASE_IMG_PATH + "/first/IMG_2919.JPG",
			&model.Person{
				Name:  "HEADOF CLIENTSUCCESS", // todo Jozef Fabian
				Email: []string{"jf@sportsinnovation.dk"},
				Phone: []string{"4552224150"}, // todo plus
				//JobTitle:     "HEAD OF CLIENT SUCCESS",
				Site:         []string{"www.sportsinnovation.dk"},
				Skype:        []string{"livejof.144"}, // todo live:jof_144
				Organization: "SPORTS INNOVATION",
				Other:        "CONTENTPRODUCTION",
			},
		},
		{
			"first IMG_2920.JPG",
			BASE_IMG_PATH + "/first/IMG_2920.JPG",
			&model.Person{
				Email:    []string{"partner@coins.game"},
				Telegram: []string{"@cg_partners"},
				Site:     []string{"coinsgame.partners"},
				Skype:    []string{"coins.game"},
			},
		},
		{
			"first IMG_2921.JPG",
			BASE_IMG_PATH + "/first/IMG_2921.JPG",
			&model.Person{
				Name:     "DEAN RAYSON",
				Email:    []string{"dean@all-in.global"},
				Phone:    []string{"+351220 991 583", "+356 770408 06", "+44 7921 239 788"},
				JobTitle: "HEAD OF SALES",
				Site:     []string{"all-in.global"},
			},
		},
		{
			"first IMG_2922.JPG",
			BASE_IMG_PATH + "/first/IMG_2922.JPG",
			&model.Person{
				Name:         "Russ Yershon",
				Email:        []string{"russell@connectingbrands.co.uk"},
				Phone:        []string{"+44 (0)7500828120"},
				JobTitle:     "Talent Manager to wide network of Football Legends Director",
				Site:         []string{"Connectingbrands.co.uk"},
				Organization: "CONNECTING", // todo CONNECTING BRANDS .co.ux
				//Skype:        []string{"russ.yershon"}, todo
				//Telegram:     []string{"@connectingt"},
			},
		},
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
				Name:         "Dariya Yeryomenko",
				Email:        []string{"dariya@pay.center"},
				Phone:        []string{"+357963341"}, // todo +35796334118
				JobTitle:     "Key Account Manager",
				Organization: "Payment.Center",
				Telegram:     []string{"@dariya_pc_cy"},
			},
		},
		{
			"3648.jpg",
			BASE_IMG_PATH + "/IMG_3648.jpg",
			&model.Person{
				Name:         "Dariya Yeryomenko",
				Email:        []string{}, // todo
				Phone:        []string{"+35796334118"},
				JobTitle:     "Key Account Manager",
				Organization: "Payment.Center",
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
				JobTitle:     "Chief Executive Officer",
				Organization: "payadmit",
				Other:        "Smart Technology Payment Solution;White Label Gateway;Cashier Service;Middleware",
			},
		},
		{
			name:    "test 3.JPG",
			imgPath: "/home/olga/projects/card_detector_imgs/3.JPG",
			expected: &model.Person{
				Email:        []string{"filip.kisala@zen.com"},
				Phone:        []string{"+48693782997"},
				Name:         "Filip Kisata",
				Organization: "zen",
				JobTitle:     "Enterprise Sales Manager",
			},
		},
		{
			name:    "test 22.JPG",
			imgPath: "/home/olga/projects/card_detector_imgs/22.JPG",
			expected: &model.Person{
				Email:    []string{"erkin@admill.io"},
				Site:     []string{"www.admill.io"},
				Phone:    []string{"+90 536 745 13 03"},
				Telegram: []string{"https://t.me/Nicola_an"},
				Name:     "Erkin Bayrakci",
				Other:    "Sepapaja tn 615551TallinnEstonia",
			},
		},
		{
			name:    "test 53.JPG",
			imgPath: "/home/olga/projects/card_detector_imgs/53.JPG",
			expected: &model.Person{
				Email:    []string{"agustin.perez-vernet@gamelounge.com"},
				Site:     []string{"gamelounge.com"},
				Name:     "Agustin Perez-Vernet",
				JobTitle: "Casino Site Manager",
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
				Phone:        []string{"+44 7885723853"},
				Name:         "Emma Fisher",
				Organization: "ADVERTISING SOLUTIONS",
				JobTitle:     "Account Director",
				Other:        "& INTERNATIONAL;International Media Planning & Buying Services:;° Sponsorships;‚ Radio & Podcasts;e Innovative Solutions;Emma Fisher",
			},
		},
		{
			"4328.JPG",
			BASE_IMG_PATH + "/IMG_4328.jpg",
			&model.Person{
				Email:        []string{"martin@369gaming.media"},
				Site:         []string{"www.369gaming.media"},
				Phone:        []string{"+59895641888"},
				Skype:        []string{"live:cid.9e53d8c1151b4b"},
				Name:         "Martin Buero",
				Organization: "369",
				JobTitle:     "Ceneral Manager", // todo General Manager
				Other:        "CAMING;LATAN;bmomn;6p93;petinry",
			},
		},
		{
			"16.JPG",
			BASE_IMG_PATH + "/16.JPG",
			&model.Person{
				Email:    []string{"shubham.dhamija@deepdivemedia.in"},
				Phone:    []string{"9034901070"},
				Skype:    []string{"live:cid.e53090522ec2bf11"},
				Telegram: []string{"@dshubham26"},
				Name:     "Shubham Dhamija",
				JobTitle: "Strategy Head",
				Other:    "WhatsApp:9034901070",
			},
		},
		{
			"4077.JPG",
			BASE_IMG_PATH + "/IMG_4077.jpg",
			&model.Person{
				//Email:        []string{"siddhartheprimeromediagroup.com"},
				Site:  []string{"siddhartheprimeromediagroup.com"}, // todo www.primeromediagroup.com
				Phone: []string{"+919953414428"},
				Skype: []string{"sidagarwal17"},
				Name:  "", // todo странный шрифт
				// Organization: "PRIMERO MEDIA", // todo
				JobTitle: "FOUNDER",
				Other:    "Sulithrth agarunl;PRIMERO;MEDIA;Official Email: siddhartheprimeromediagroup.com;Official Website:www.primeromediagroup.com;Skype: sidagarwal17;ST.2016;CROUP",
			},
		},
		{
			"4095.JPG",
			BASE_IMG_PATH + "/IMG_4095.jpg",
			&model.Person{
				Email:        []string{"taras.kolesnikov@coinspaid.com"},
				Telegram:     []string{"t.me/Taras_CoinsPaid"},
				Phone:        []string{"+375445165298"},
				Name:         "TARAS KOLESNIKOV",
				Organization: "CoinsPaid",
				JobTitle:     "Sales manager",
			},
		},
		{
			"69.JPG",
			BASE_IMG_PATH + "/69.JPG",
			&model.Person{
				//Email:        []string{"viorel.stan@gshmedia.com"}, // todo
				Site:  []string{"viorel.stanagshmedia.com"}, // todo
				Skype: []string{"viorel.stan87"},
				Name:  "", // todo Viorel Stan
				//Organization: "GSH", // todo
				JobTitle: "CEO",
				Other:    "ViorelStan",
			},
		},
		//{
		// todo russian text
		//	"3606.JPG",
		//	BASE_IMG_PATH + "/IMG_3606.jpg",
		//	&model.Person{
		//		Name:  "ЕЛЕНА СОЛОДУХИНА",
		//		Email: []string{"KVARTA@KVARTA.RU"},
		//		//Site:  []string{"KVARTA.RU"},
		//		Phone: []string{
		//			"+7 (473) 20-20-457",
		//			"+7 (473) 200-0-300",
		//		},
		//		//JobTitle: "МЕНЕДЖЕР ПО РАБОТЕ С КЛИЕНТАМИ", // todo
		//		Other: "КВАРТА». ПЕЧАТАЕМ С 1991 TODA",
		//	},
		//},
	}

	testDetector, config := createDetector2(t)

	manage_file.ClearFolder(config.StorageFolder)
	manage_file.ClearFolder("./tmp")
	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

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

func equalIgnoreCase(t *testing.T, expected, actual []string, field string) {
	for _, e := range expected {
		found := false
		for _, a := range actual {
			if strings.EqualFold(e, a) {
				found = true
				break
			}
		}
		assert.True(t, found, "%s: expected: %s, actual: %s", field, e, actual)
	}
}

func equalIgnoreSpaces(t *testing.T, expected, actual []string, field string) {
	for _, e := range expected {
		found := false
		for _, a := range actual {
			if strings.ReplaceAll(e, " ", "") == strings.ReplaceAll(a, " ", "") {
				found = true
				break
			}
		}
		assert.True(t, found, "%s: expected: %s, actual: %s", field, e, actual)
	}
}

// Detector with paddle ocr
func createDetector2(t *testing.T) (*service.Detector2, *app.Config) {

	os.Setenv("CONFIG_FILE", "./config/config.yml")
	isLogTime := true

	config, err := app.NewConfigFromYml()
	if err != nil {
		t.Fatal(err)
	}

	cardRepo := inmemory.NewCardRepo()
	imgPreparer := img_prepare.NewService(config.StorageFolder)

	textRecognizer, err := paddleocr.NewService(isLogTime,
		config.Paddleocr.RunPath,
		config.Paddleocr.DetPath,
		config.Paddleocr.RecPath,
		config.TmpFolder)
	if err != nil {
		t.Fatal(err)
	}
	cardDetector, err := onnx.NewService(
		config.Onnx.PathRuntime,
		config.Onnx.PathModel,
		isLogTime)
	if err != nil {
		log.Fatal("card detector creation error", err)
	}
	fieldSorter := field_sort.NewService(config.PathProfessionList, config.PathCompanyList, config.PathNamesList, isLogTime)

	// detector
	testDetector := service.NewDetector2(
		imgPreparer,
		textRecognizer,
		cardDetector,
		fieldSorter,
		cardRepo,
		config.StorageFolder,
		isLogTime,
		config.IsDebug)

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
