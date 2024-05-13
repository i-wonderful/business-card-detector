package integration_test

import (
	"card_detector/internal/service/text_recognize"
	manage_file "card_detector/internal/util/file"
	"card_detector/internal/util/img"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/otiai10/gosseract/v2"
	"github.com/stretchr/testify/assert"
	"reflect"

	"testing"
)

func TestRecognizeText(t *testing.T) {

	testCases := []struct {
		name     string
		imgPath  string
		expected string
	}{
		{
			name:     "1",
			imgPath:  "./data/vladislav.jpg",
			expected: "Vladyslav Kolodistyi",
		},
		{
			name:     "2", // 10.JPG
			imgPath:  "./data/ilona.jpg",
			expected: "JLONA KROLE",
		},
		{
			name:     "key_acc_manager", // 3.JPG
			imgPath:  "./data/key_acc_manager_2926.jpg",
			expected: "Key account manager",
		},
	}

	service := text_recognize.NewService(true, "../config/tesseract/")
	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			bytes := manage_file.ReadFileBytes(tc.imgPath)

			im, _ := img.ToImage(bytes)

			bytes = img.ToBytes(im)
			content := [][]byte{bytes}

			//img.SaveImg("./tmp/"+fmt.Sprintf("im_gosseract_%s.jpg", tc.name), bytes)

			text, err := service.RecognizeBatch(content)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tc.expected, text[0])
		})
	}
}

// Тестирую предобработку изображения перед отправкой в gosseract
func TestRecognizeTextRawImg(t *testing.T) {

	testCases := []struct {
		name     string
		imgPath  string
		expected string
	}{
		{
			name:     "dariaraw", // IMG_2926.JPG
			imgPath:  "./data/dariya_raw.jpg",
			expected: "Dariya Yeryomenko",
		},
		{
			name:     "email_raw", // IMG_2926.JPG
			imgPath:  "./data/email_raw_2917.jpg",
			expected: "Mail: b2b@Linebet.com Skype: partners@Linebet.com",
		},
		{
			name:     "key_acc_manager_raw", // IMG_2926.JPG
			imgPath:  "./data/key_acc_manager_2926_raw.jpg",
			expected: "Key Account Manager",
		},
		{
			name:     "tg_dariya_raw", // IMG_2926.JPG
			imgPath:  "./data/tg_dariya_2926_raw.jpg",
			expected: "@dariya_pc_cy.",
		},
		{
			name:     "skype_raw",
			imgPath:  "./data/skype_4328_raw.jpg",
			expected: "Skype live: cid. 9e53d8c1 1 51546", //51b4b
		},
		{
			name:     "phone_raw",
			imgPath:  "./data/phone_2919_raw.jpg",
			expected: "+45 52 22 41 50",
		},
		{
			// phone
			name:     "phone 2912",
			imgPath:  "./data/phone_2912_raw2.jpg",
			expected: "+356 99723767",
		},
	}

	service := text_recognize.NewService(true, "../config/tesseract/")
	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			im, _ := img.OpenImg(tc.imgPath)
			fmt.Println("type im:", reflect.TypeOf(im))
			//bytes := manage_file.ReadFileBytes(tc.imgPath)

			//im, _ := img.ToImage(bytes)
			// Серый
			//im = imaging.Grayscale(im)
			// увеличить контраст
			im = imaging.AdjustContrast(im, 20)
			// резкость
			im = imaging.Sharpen(im, 0.35)
			// светлость
			//im = imaging.AdjustGamma(im, 1.2)

			fmt.Println("type im:", reflect.TypeOf(im))
			bytes := img.ToBytes(im)
			content := [][]byte{bytes}
			img.SaveImg("./tmp/"+fmt.Sprintf("im_gosseract_%s.jpg", tc.name), bytes)

			text, err := service.RecognizeBatch(content)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tc.expected, text[0])
		})
	}
}

func TestDetectFromFile(t *testing.T) {

	client := gosseract.NewClient()
	client.SetLanguage("eng+rus")
	client.SetPageSegMode(gosseract.PSM_RAW_LINE | gosseract.PSM_SINGLE_LINE)
	client.SetBlacklist("№;`^>\\'‘")

	err := client.SetImage("./testdata/crop.jpeg")
	if err != nil {
		t.Error(err)
	}
	text, err := client.Text()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(text)
}
