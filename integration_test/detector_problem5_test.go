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
				JobTitle:     "COO (Chief Operating Officer)",
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
				Phone:        []string{"+371)67772962", "+371)22352883"},
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
			"0_17. kirill.jpg", // 1_17. kirill.jpg
			BASE_IMG_PROBLEM5_PATH + "/0_17. kirill.jpg",
			&model.Person{
				Email:    []string{"d.k@scipiopay.com"},
				Site:     []string{"scipiopay.com"},
				Telegram: []string{"@KlrillMan"},
				Name:     "KIRILL",
				JobTitle: "Business Developer",
				Other:    "com;Cpr;kumo.Ureoio;U126;wwVia.mcteut;81.-12;apura;trust;Telegram;Web;adptrepo",
			},
		},
		{
			"0_21. Alians.jpg",
			BASE_IMG_PROBLEM5_PATH + "/0_21. Alians.jpg",
			&model.Person{
				Email:        []string{"ba.abou@aliansgroup.tech", "contact@aliansgroup.tech"},
				Site:         []string{"www.alianspay.com"},
				Phone:        []string{"+221774387447", "+2250595999760", "+237678456912"},
				Name:         "BA ABOU ALHOUSSEYNI",
				Organization: "AliansPay",
				JobTitle:     "Chief Business Officer",
				Other:        "Oyment;Fournisseur de solutions de paiement;2nd Stage NDEYE KHADY GUEYE Building-Dakar-Senegal",
			},
		},
		{
			"0_22. alina dashaeva.jpg",
			BASE_IMG_PROBLEM5_PATH + "/0_22. alina dashaeva.jpg",
			&model.Person{
				Email:        []string{"alina.d@gt-charge.com"},
				Site:         []string{"gt-charge.com"},
				Phone:        []string{"+79772723695"},
				Name:         "Alina Dashaeva",
				Telegram:     []string{"@GTC_Payment"},
				Organization: "Global Transaction Charge",
				Other:        "u2.femo;Btopcce;Phone;Email;Web",
			},
		},
		{
			"1_3 nathalia mesh.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_3 nathalia mesh.jpg",
			&model.Person{
				Email:        []string{"nathalia@meshconnect.com"},
				Site:         []string{"meshconnect.com"},
				Telegram:     []string{"@nathijuelos"},
				Name:         "Nathalia Hijuelos",
				Organization: "MESH",
				JobTitle:     "Global Account Executive",
				Other:        "crpto",
			},
		},
		{
			"1_6 andressa belvo.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_6 andressa belvo.jpg",
			&model.Person{
				Email:        []string{"andressa.brandao@belvo.com"},
				Site:         []string{"www.belvo.com"},
				Phone:        []string{"+551197122-9898"},
				Name:         "Andressa Brandao",
				Organization: "belvo",
				JobTitle:     "Account Executive",
				Other:        "01X",
			},
		},
		{
			"1_10 delano genial.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_10 delano genial.jpg",
			&model.Person{
				Email: []string{"delano.queiroz@genial.com.vc"},
				Site:  []string{"genialinvestimentos.com.br"},
				Phone: []string{"+55 (21)99945-9191", "+55(11)3206-8267/807"},
				Name:  "Delano Queiroz",
				Other: "tankacgupt;Preee;eli;Sao Paulo-SP",
			},
		},
		{
			"1_20 vladislav globus.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_20 vladislav globus.jpg",
			&model.Person{
				Email:        []string{"info@globuspay.io"},
				Phone:        []string{"+971558318566"},
				Telegram:     []string{"@globuspaid"},
				Name:         "Vladislav Belov",
				Organization: "GLOBUSPAY",
				JobTitle:     "Head of Business Development",
				Other:        "sripe;Vsemc;We provide best payment;solutions for your business",
			},
		},
		{
			"1_28 kolawole wewire.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_28 kolawole wewire.jpg",
			&model.Person{
				Email:        []string{"kofo@wewire.com"},
				Site:         []string{"wewire.com"},
				Phone:        []string{"+234 807444 0966"},
				Name:         "Kofoworola Kolawole",
				Organization: "WeWire",
				JobTitle:     "Lead, BD Nigeria",
				Other:        "Daynents",
			},
		},
		{
			"1_29 dende first bank.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_29 dende first bank.jpg",
			&model.Person{
				Email:        []string{"Zacharie.d.dende@fbnbankrdc.com"},
				Site:         []string{"www.fbnbankrdc.com"},
				Phone:        []string{"+243817002229", "+243829832902"},
				Name:         "DENDE DIA DENDE Zacharie",
				Organization: "FirstBank",
				JobTitle:     "Branch Manager",
				Other:        "DRC;FirstBank DRC SA;07,Av/ de l'Aerodrome,C/Kalamu",
			},
		},
		{
			"1_32 alberta nasno.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_32 alberta nasno.jpg",
			&model.Person{
				Email:    []string{"a.boateng@nsano.com"},
				Site:     []string{"www.nsano.com"},
				Phone:    []string{"+260767527865", "+260767636343"},
				Name:     "Alberta S.Boateng",
				JobTitle: "Country Manager",
				Other:    "ouesu;poyments;Plot.No.6755.Chainama Rd.Olynpia Ext. Lusaka",
			},
		},
		{
			"1_34 cenk lingdom bank.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_34 cenk lingdom bank.jpg",
			&model.Person{
				Email:        []string{"cenk.saltan@thekingdombank.com"},
				Site:         []string{"thekingdombank.com"},
				Name:         "Cenk Saltan",
				Organization: "Kingdom Bank", // todo The Kingdom Bank
				JobTitle:     "Business Development Manager",
				Other:        "seftlement",
			},
		},
		{
			"1_35 jamie orbital.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_35 jamie orbital.jpg",
			&model.Person{
				Email:    []string{"jamie.zammitt@getorbital.com"},
				Site:     []string{"getorbital.com"},
				Phone:    []string{"+35054096000"},
				Name:     "Jamie Zammitt",
				JobTitle: "Business Development Director",
				Other:    "2C2 Leisure lsland;Business Centre;Ocean Village, Gibraltar",
			},
		},
		{
			"1_36 amit netwallet.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_36 amit netwallet.jpg",
			&model.Person{
				Email:        []string{"lehri@thenetwallet.com"},
				Phone:        []string{"+971585935717"},
				Telegram:     []string{"@netwallet1"},
				Name:         "AMIT BANSAL",
				Organization: "NetWallet",
				JobTitle:     "Account Manager Payments-Asia Pacific",
				Other:        "Spend Your Money With Confidence",
			},
		},
		{
			"1_39 glenn revpanda.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_39 glenn revpanda.jpg",
			&model.Person{
				Email:        []string{"glenn@revpanda.com"},
				Site:         []string{"revpanda.com"},
				Phone:        []string{"+35699883592"},
				Telegram:     []string{"@glenn79"},
				Skype:        []string{"glenn_debattista"},
				Name:         "GLENN DEBATTISTA",
				Organization: "REV",
				JobTitle:     "Chief Operating Officer",
				Other:        "ONE STOP DIGITAL;MARKETING AGENCY",
			},
		},
		{
			"1_42 vittor paybrokers.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_42 vittor paybrokers.jpg",
			&model.Person{
				Email:        []string{"vittor.alberti@paybrokers.com.br"},
				Phone:        []string{"+5541992535669"},
				Name:         "VITTOR ALBERTI",
				Organization: "PAYBROKERS",
				JobTitle:     "ACCOUNT MANAGER",
				Other:        "PiX;com",
			},
		},
		{
			"1_43 taras coinspaid.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_43 taras coinspaid.jpg",
			&model.Person{
				Email:    []string{"taras.kolesnikov@coinspaid.com"},
				Telegram: []string{"@Taras_CoinsPaid"},
				Name:     "Taras Kolesnikov",
				JobTitle: "LEAD SALES MANAGER",
				Other:    "etypto",
			},
		},
		{
			"1_45 caroline capitalize.jpg",
			BASE_IMG_PROBLEM5_PATH + "/1_45 caroline capitalize.jpg",
			&model.Person{
				Email:        []string{"cm@capitalixe.com"},
				Site:         []string{"www.capitalixe.com"},
				Phone:        []string{"+4402080888035", "+44 0)7553734915"},
				Name:         "Caroline Moreno",
				Organization: "",
				JobTitle:     "Head of Business Development",
				Other:        "erypto;Forbes;HONOREE;EUROPE2021",
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
