package paddleocr

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"card_detector/internal/util/img"
	"fmt"
	"image"
	"log"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

const MIN_WORD_LEN = 3

type TextRecognizeService struct {
	pathToRun   string
	pathRecOnnx string
	pathDetOnnx string
	pathPython  string
	isLog       bool
	pathTmp     string
}

func NewService(isLog bool, pathToRun, pathDetOnnx, pathRecOnnx, pathTmp string) (*TextRecognizeService, error) {
	if !manage_file.FileExists(pathToRun) {
		return nil, fmt.Errorf("file not found: %s", pathToRun)
	}

	if !manage_file.FileExists(pathDetOnnx) {
		return nil, fmt.Errorf("ONNX detection not found: %s", pathDetOnnx)
	}

	if !manage_file.FileExists(pathRecOnnx) {
		return nil, fmt.Errorf("ONNX recognition not found: %s", pathRecOnnx)
	}

	return &TextRecognizeService{
		pathToRun:   pathToRun,
		pathDetOnnx: pathDetOnnx, // "./lib/paddleocr/onnx/en_PP-OCRv3_det_infer.onnx",
		pathRecOnnx: pathRecOnnx, // "./lib/paddleocr/onnx/en_PP-OCRv4_rec_infer.onnx",
		pathPython:  "python",    // "python"
		pathTmp:     pathTmp,
		isLog:       isLog,
	}, nil
}

// RecognizeImg - recognize text from image, save image in tmp folder
// @param im - image
// @return - list of words
func (s *TextRecognizeService) RecognizeImg(im *image.Image) ([]model.DetectWorld, error) {
	if s.isLog {
		start := time.Now()
		defer func() {
			log.Printf(">>> Time paddle recognize: %s", time.Since(start))
		}()
	}

	resized := img.ResizeImageByHeight(*im, 1800)
	filePath := manage_file.GenerateFileName(s.pathTmp, "for_paddle", "jpg")
	img.SaveJpegWithQality(&resized, filePath, 87)
	absPath, _ := filepath.Abs(filePath)

	return s.RecognizeImgByPath(absPath)
}

func (s *TextRecognizeService) RecognizeImgByPath(pathImg string) ([]model.DetectWorld, error) {
	if s.isLog {
		start := time.Now()
		defer func() {
			log.Printf(">>> Time paddle recognize: %s", time.Since(start))
		}()
	}

	cmd := exec.Command(s.pathPython, s.pathToRun, pathImg, s.pathDetOnnx, s.pathRecOnnx, "stdout")

	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error : %v", err)
		return nil, err
	}

	raw := string(output)

	worlds, err := parseString(raw)
	if err != nil {
		return nil, err
	}

	return worlds, nil
}

func parseString(input string) ([]model.DetectWorld, error) {
	re := regexp.MustCompile(`\[\[\[(\d+\.\d+), (\d+\.\d+)\], \[(\d+\.\d+), (\d+\.\d+)\], \[(\d+\.\d+), (\d+\.\d+)\], \[(\d+\.\d+), (\d+\.\d+)\]\], ["'](.+?)["'], (\d+\.\d+)\]`)

	matches := re.FindAllStringSubmatch(input, -1)
	if matches == nil {
		return nil, fmt.Errorf("no matches found")
	}

	var results []model.DetectWorld
	for _, match := range matches {
		text := match[9]
		if len(text) < MIN_WORD_LEN {
			continue
		}
		p1X, _ := strconv.ParseFloat(match[1], 64)
		p1Y, _ := strconv.ParseFloat(match[2], 64)
		p2X, _ := strconv.ParseFloat(match[3], 64)
		p2Y, _ := strconv.ParseFloat(match[4], 64)
		p3X, _ := strconv.ParseFloat(match[5], 64)
		p3Y, _ := strconv.ParseFloat(match[6], 64)
		p4X, _ := strconv.ParseFloat(match[7], 64)
		p4Y, _ := strconv.ParseFloat(match[8], 64)
		prob, _ := strconv.ParseFloat(match[10], 64)
		detectWorld := model.DetectWorld{
			Text: text,
			Box: model.NewBoxFromPoints(
				model.Point{X: int(p1X), Y: int(p1Y)},
				model.Point{X: int(p2X), Y: int(p2Y)},
				model.Point{X: int(p3X), Y: int(p3Y)},
				model.Point{X: int(p4X), Y: int(p4Y)},
			),
			Prob: float32(prob),
		}
		results = append(results, detectWorld)
	}

	return results, nil
}
