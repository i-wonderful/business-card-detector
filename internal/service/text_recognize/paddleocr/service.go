package paddleocr

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"fmt"
	"log"
	"math"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"time"
)

const MIN_WORD_LEN = 3

type TextRecognizeService struct {
	pathToPythonRun string
	pathRecOnnx     string
	pathDetOnnx     string
	isLog           bool
}

func NewService(isLog bool, pathToRun string, pathDetOnnx string, pathRecOnnx string) (*TextRecognizeService, error) {
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
		pathToPythonRun: pathToRun,
		pathDetOnnx:     pathDetOnnx, // "./lib/paddleocr/onnx/en_PP-OCRv3_det_infer.onnx",
		pathRecOnnx:     pathRecOnnx, // "./lib/paddleocr/onnx/en_PP-OCRv4_rec_infer.onnx",
		isLog:           isLog,
	}, nil
}

func (s *TextRecognizeService) RecognizeAll(path string) ([]model.DetectWorld, error) {
	if s.isLog {
		start := time.Now()
		defer func() {
			log.Printf(">>> Time paddle recognize: %s", time.Since(start))
		}()
	}

	// "/home/olga/env/bin/python"
	//
	// "python"
	cmd := exec.Command("python", s.pathToPythonRun, path, "stdout")

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

	worlds = sortByProbAndY(worlds)
	return worlds, nil
}

func parseString(input string) ([]model.DetectWorld, error) {
	re := regexp.MustCompile(`\[\[\[(\d+\.\d+), (\d+\.\d+)\], \[(\d+\.\d+), (\d+\.\d+)\], \[(\d+\.\d+), (\d+\.\d+)\], \[(\d+\.\d+), (\d+\.\d+)\]\], '(.+?)', (\d+\.\d+)\]`)

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
			Rect: model.Rectangle{
				PTop1: model.Point{X: int(p1X), Y: int(p1Y)},
				PTop2: model.Point{X: int(p2X), Y: int(p2Y)},
				PBot1: model.Point{X: int(p3X), Y: int(p3Y)},
				PBot2: model.Point{X: int(p4X), Y: int(p4Y)},
			},
			Prob: float32(prob),
		}
		results = append(results, detectWorld)
	}

	return results, nil
}

func sortByProbAndY(worlds []model.DetectWorld) []model.DetectWorld {
	sort.Slice(worlds, func(i, j int) bool {
		w1 := worlds[i]
		w2 := worlds[j]
		if math.Abs(float64(w1.Prob-w2.Prob)) > 0.04 {
			return w1.Prob > w2.Prob
		}
		return w1.Rect.PTop1.Y < w2.Rect.PTop1.Y
	})
	return worlds
}
