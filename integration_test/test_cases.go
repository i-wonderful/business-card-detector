package integration

import (
	. "card_detector/internal/model"
)

type testCase struct {
	name     string
	filePath string
	//expectedHTTPCode int
	expected *Person
}

var testCases = []testCase{
	//{
	//	name:     "IMG_2912.JPG",
	//	filePath: "./data/IMG_2912.JPG",
	//	//expectedHTTPCode: http.StatusOK,
	//	expected: &Person{
	//		Email:        "oren@delasport.com",
	//		Site:         "www.delasport.com",
	//		Phone:        "+356 99723767",
	//		Name:         "OREN COHEN SHWARTZ",
	//		Organization: "DELIVER SPORTS",
	//		JobTitle:     "CEO",
	//		Other:        "Г Delasport;OREN COHEN SHWARTZ", //  todo doubles
	//	},
	//}, {
	//	name:     "IMG_2913.JPG",
	//	filePath: "./data/IMG_2913.JPG",
	//	expected: &Person{
	//		Email:        "craig@summer-creative.com",
	//		Phone:        "+34 711 017 134",
	//		Name:         "Be CREATIVE Edwards",
	//		Organization: "Business Developmen",
	//		JobTitle:     "Manager",
	//		Other:        "+44 (0)20 3355 5336;SUMMER;Craig;Be CREATIVE",
	//	},
	//}, {
	//	name:     "IMG_2914.JPG",
	//	filePath: "./data/IMG_2914.JPG",
	//	expected: &Person{
	//		Name:     "Areg Oganesian",
	//		Email:    "areg@igtrm.com",
	//		Phone:    "+374 99 452772",
	//		JobTitle: "CEO",
	//		//Site: "www.igtrm.com", // todo как разные блоки распознает
	//	},
	//}, {
	//	name:     "IMG_2915.JPG",
	//	filePath: "./data/IMG_2915.JPG",
	//	expected: &Person{
	//		Name:         "Nadezda Tereshchenko",
	//		Email:        "info@platinum-expo.com",
	//		JobTitle:     "Creative director",
	//		Site:         "platinum-expo.com",
	//		Telegram:     "@Naddiz_T",
	//		Organization: "PLATINUM",
	//	},
	//},
	//{
	//	name:     "IMG_2916.JPG",
	//	filePath: "./data/IMG_2916.JPG",
	//	expected: &Person{
	//		Email: "info@revpanda.com",
	//	},
	//},
	//{
	//	"IMG_2917.JPG",
	//	"./data/IMG_2917.JPG",
	//	&Person{
	//		Email:        "b2b@linebet.com",
	//		Skype:        "partners@linebet.com",
	//		Telegram:     "@linebet",
	//		Site:         "linebet.com",
	//		Organization: "B2B Department",
	//	},
	//},
	//{
	//	"IMG_2918.JPG",
	//	"./data/IMG_2918.JPG",
	//	&Person{
	//		Name:         "Aron Myerthall",
	//		Email:        "aron@raventrack.com",
	//		Phone:        "07956 710535",
	//		JobTitle:     "Sales Manager",
	//		Organization: "-R/VENTRACK",
	//	},
	//},

	//{
	//	"IMG_2920.JPG",
	//	"./data/IMG_2920.JPG",
	//	&Person{
	//		Email:    "partner@coins.game",
	//		Telegram: "@cg_partners",
	//		Site:     "coins.game",
	//		Skype:    "coinsgame.partners",
	//	},
	//},
	//{
	//	"IMG_2921.JPG",
	//	"./data/IMG_2921.JPG",
	//	&Person{
	//		Name:     "DEAN RAYSON",
	//		Email:    "dean@all-in.global",
	//		Phone:    "+351 220 991 583",
	//		JobTitle: "HEAD OF SALES",
	//		Site:     "all-in.global",
	//		Other:    "MT Office +356 770 408 06;UK Mobile +44 7921 239 788",
	//	},
	//},
	//{
	//	"IMG_2922.JPG",
	//	"./data/IMG_2922.JPG",
	//	&Person{
	//		Name:         "Russ Yershon",
	//		Email:        "russell@connectingbrands.co.uk",
	//		Phone:        "+44 (0) 7500828120",
	//		JobTitle:     "Director",
	//		Site:         "Connectingbrands.co.uk",
	//		Organization: "CONNECTING", // todo CONNECTING BRANDS .co.ux
	//		Skype:        "russ.yershon",
	//		Telegram:     "@connectingt",
	//	},
	//},
	//{
	//	"IMG_2923.JPG",
	//	"./data/IMG_2923.JPG",
	//	&Person{
	//		Name:         "Milda Grigaliunaite",
	//		Email:        "milda@world-card.com",
	//		Phone:        "+ 44 20 3835 6450",
	//		JobTitle:     "Backoffice Supervisor",
	//		Organization: "WsrldCard",
	//	},
	//},
	//{
	//	"IMG_2924.JPG",
	//	"./data/IMG_2924.JPG",
	//	&Person{
	//		Name:         "Naman Sharma",
	//		Email:        "naman.sharma@phoenixtech.consulting",
	//		Phone:        "+91-8607860787",
	//		JobTitle:     "(Business Development Manager)",
	//		Site:         "https://phoenixtech.consulting",
	//		Organization: "Phoenix Tech Consulting",
	//		Other:        "Haryana 122001;~hoenix;147 Vipul Trade Center, Sohna -;Pvt.Ltd;Road, Sector 48, Gurugram,",
	//	},
	//},
	//{
	//	"IMG_2925.JPG",
	//	"./data/IMG_2925.JPG",
	//	&Person{
	//		Name:         "Karina",
	//		Email:        "karen.s@nowpayments.io",
	//		JobTitle:     "Business Development Manager",
	//		Site:         "nowpayments.io",
	//		Organization: "Payments",
	//		Other:        "Seven Oo FAY WENT G;Йо PAYMENT GATE WR",
	//	},
	//},
	//{
	//	"IMG_2926.JPG",
	//	"./data/IMG_2926.JPG",
	//	&Person{
	//		Name:         "Dariya Yeryomenko",
	//		Email:        "dariya@pay.center",
	//		Phone:        "+357 963 341 18",
	//		JobTitle:     "Key Account Manager",
	//		Organization: "Payment.Center",
	//		Telegram:     "@dariya_pc_ey",
	//	},
	//},
	//{
	//	"IMG_2927.JPG",
	//	"./data/IMG_2927.JPG",
	//	&Person{
	//		Name:         "Contact", // todo wrong
	//		Email:        "business@libernetix.com",
	//		Site:         "libernetix.com",
	//		Organization: "libernetix team",
	//		Other:        "and partnerships.;t us for collaborations",
	//	},
	//},
	//{
	//	"IMG_2928.JPG",
	//	"./data/IMG_2928.JPG",
	//	&Person{
	//		Name:         "Kate Elterova",
	//		Email:        "kate@oddsdigger.com",
	//		Telegram:     "@KateElterova",
	//		JobTitle:     "AFFILIATE MANAGER",
	//		Skype:        "elterova.ekaterina",
	//		Organization: "DEV", // todo wrong
	//	},
	//},
	//{
	//	"rus_example.jpg",
	//	"./data/rus_example.jpg",
	//	&Person{
	//		Email: "unstroy444@mail.ru",
	//		Site:  "gk-grishkino.ru",
	//		Phone: "8-910-539-84-44",
	//		Name:  "Валерий Викторови Некрылов",
	//		Other: "из ЛСТК;||роектирование;производство;Строительство;< |",
	//	},
	//},
	// todo change to arrays
}
