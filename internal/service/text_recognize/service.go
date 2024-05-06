package text_recognize

import (
	"log"
	"time"

	"github.com/otiai10/gosseract/v2"
)

// TextRecognizeService - recognize text from image
type TextRecognizeService struct {
	client    *gosseract.Client
	isLogTime bool
}

func NewService(isLogTime bool) *TextRecognizeService {
	client := gosseract.NewClient()
	client.SetLanguage("eng+rus")
	client.SetPageSegMode(gosseract.PSM_SINGLE_LINE)
	client.SetBlacklist("№;`^>\\'‘›!¢][¥|")

	//client.SetVariable("user_words", "../config/names")

	return &TextRecognizeService{
		client:    client,
		isLogTime: isLogTime,
	}
}

func (s *TextRecognizeService) RecognizeBatch(contents [][]byte) ([]string, error) {
	if s.isLogTime {
		start := time.Now()
		defer func() {
			log.Printf(">>> Time gosseract : %s", time.Since(start))
		}()
	}

	results := make([]string, len(contents))

	for i, content := range contents {
		text, err := s.recognize(content)
		if err != nil {
			return nil, err
		}
		results[i] = text
	}

	return results, nil

}

func (s *TextRecognizeService) recognize(content []byte) (string, error) {
	err := s.client.SetImageFromBytes(content)
	if err != nil {
		return "", err
	}
	text, err := s.client.Text()
	return text, err
}
