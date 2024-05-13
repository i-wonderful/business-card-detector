package text_recognize

import (
	"log"
	"regexp"
	"time"

	"github.com/otiai10/gosseract/v2"
)

// TextRecognizeService - recognize text from image
type TextRecognizeService struct {
	lang         string
	pathSettings string
	isLogTime    bool
}

func NewService(isLogTime bool, pathSettings string) *TextRecognizeService {
	return &TextRecognizeService{
		lang:         ENG,
		pathSettings: pathSettings,
		isLogTime:    isLogTime,
	}
}

func (s *TextRecognizeService) RecognizeBatch(contents [][]byte) ([]string, error) {
	if s.isLogTime {
		start := time.Now()
		defer func() {
			log.Printf(">>> Time gosseract : %s", time.Since(start))
		}()
	}

	clientBase, err := s.getClient()
	if err != nil {
		return nil, err
	}
	defer clientBase.Close()

	clientPhone, err := s.getClientForPhone()
	if err != nil {
		return nil, err
	}
	defer clientPhone.Close()

	results := make([]string, len(contents))

	for i, content := range contents {
		text, err := recognize(content, clientBase)
		if err != nil {
			return nil, err
		}
		if isPhone(text) {
			if text, err = recognize(content, clientPhone); err != nil {
				return nil, err
			}
		}
		results[i] = text
	}

	return results, nil
}

//func (s *TextRecognizeService) RecognizeByPath(paths []string) ([]string, error) {
//	if s.isLogTime {
//		start := time.Now()
//		defer func() {
//			log.Printf(">>> Time gosseract : %s", time.Since(start))
//		}()
//	}
//
//	clientBase, err := s.getClient()
//	if err != nil {
//		return nil, err
//	}
//
//	results := make([]string, len(paths))
//
//	for i, path := range paths {
//		text, err := s.recognizeByPath(path, clientBase)
//		if err != nil {
//			return nil, err
//		}
//		results[i] = text
//	}
//
//	return results, nil
//}

func (s *TextRecognizeService) DetectLang(path string) {
	lang := DetectLang(path)
	if lang == RUS {
		log.Printf("Language: %s", RUS)
		s.lang = "eng+rus"
	} else {
		s.lang = "eng"
	}
}

func (s *TextRecognizeService) getClient() (*gosseract.Client, error) {
	client := gosseract.NewClient()
	client.SetLanguage(s.lang)
	err := client.SetConfigFile(s.pathSettings + "base.config")
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *TextRecognizeService) getClientForPhone() (*gosseract.Client, error) {
	client := gosseract.NewClient()
	client.SetLanguage("rus")
	err := client.SetConfigFile(s.pathSettings + "detect-phone.config")
	if err != nil {
		return nil, err
	}
	return client, nil
}

func recognize(content []byte, cl *gosseract.Client) (string, error) {
	err := cl.SetImageFromBytes(content)
	if err != nil {
		return "", err
	}
	return cl.Text()
}

func (s *TextRecognizeService) recognizeByPath(path string, cl *gosseract.Client) (string, error) {
	err := cl.SetImage(path)
	if err != nil {
		return "", err
	}
	return cl.Text()
}

func isPhone(val string) bool {
	digitRegex := regexp.MustCompile(`\d+`)
	letterRegex := regexp.MustCompile(`[a-zA-Z]+`)

	digits := digitRegex.FindAllString(val, -1)
	letters := letterRegex.FindAllString(val, -1)

	digitCount := 0
	for _, d := range digits {
		digitCount += len(d)
	}

	if digitCount < 4 {
		return false
	}

	letterCount := 0
	for _, l := range letters {
		letterCount += len(l)
	}

	if letterCount > 2 {
		return false
	}

	return true
}
