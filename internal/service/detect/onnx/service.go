package onnx

import (
	"fmt"
	"github.com/nfnt/resize"
	ort "github.com/yalue/onnxruntime_go"
	"image"
	"math"
	"sort"
	"time"

	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"card_detector/pkg/log"
)

const IMG_SIZE = 640
const COUNT_CLASSES = 9
const CARD_CLASS = "card"

var yolo_classes = []string{"card", "location", "logo", "mail", "phone", "skype", "telegram", "web", "whatsapp"}

const IOU_LIMIT = 0.8 // 0.8
const PROB_MIN = 0.3  // 0.5

type box struct {
	x1, y1, x2, y2 float64
	label          string
	prob           float32
}

type CardDetectService struct {
	pathToOnnxRuntime string
	pathToModel       string
	isLogTime         bool
	logger            *log.Logger
}

func NewService(pathToOnnxRuntime string, pathToModel string, isLogTime bool, logger *log.Logger) (*CardDetectService, error) {
	if manage_file.FileExists(pathToOnnxRuntime) == false {
		return nil, fmt.Errorf("file onnxruntime not found: %s", pathToOnnxRuntime)
	}

	if manage_file.FileExists(pathToModel) == false {
		return nil, fmt.Errorf("file model not found: %s", pathToModel)
	}

	ort.SetSharedLibraryPath(pathToOnnxRuntime)
	err := ort.InitializeEnvironment()
	if err != nil {
		return nil, err
	}

	return &CardDetectService{
		pathToOnnxRuntime: pathToOnnxRuntime,
		pathToModel:       pathToModel,
		isLogTime:         isLogTime,
		logger:            logger,
	}, nil
}

func (s *CardDetectService) Detect(img image.Image) ([]model.TextArea, error) {
	if s.isLogTime {
		start := time.Now()
		defer func() {
			s.logger.Debug(fmt.Sprintf(">>> Time onnx: %s", time.Since(start)))
		}()
	}

	rawPrediction, err := detect_objects_on_image(s.pathToModel, img)
	if err != nil {
		s.logger.Error("Onnx error: ", err)
		return nil, err
	}

	result := []model.TextArea{}
	for _, p := range rawPrediction {

		x1 := p.x1
		y1 := p.y1
		x2 := p.x2
		y2 := p.y2

		h := y2 - y1
		w := x2 - x1

		prediction := model.TextArea{X: int(x1), Y: int(y1), Width: int(w), Height: int(h), Label: p.label, Prob: p.prob}
		result = append(result, prediction)
	}

	return result, nil
}

// Returns Array of bounding boxes in format [[x1,y1,x2,y2,object_type,probability],..]
func detect_objects_on_image(pathModel string, img image.Image) ([]box, error) {
	input, img_width, img_height := prepare_input(img)
	output, err := runModel(pathModel, input)
	if err != nil {
		return nil, err
	}
	return process_output(output, img_width, img_height), nil
}

//	the ONNX library for Go, requires you to provide tensor RGB as a flat array,
//
// e.g. to concat these three colors in one after one
func prepare_input(src image.Image) ([]float32, int64, int64) {
	size := src.Bounds().Size()
	img_width, img_height := int64(size.X), int64(size.Y)

	//img = img2.ScaleToSquare(img, IMG_SIZE)
	//img := imaging.Clone(src)
	//img := img2.NormalizeImage(src)
	//img = imaging.Resize(img, IMG_SIZE, IMG_SIZE, imaging.Lanczos)
	//imaging.Invert(&img)

	//img := imaging.Resize(src, IMG_SIZE, IMG_SIZE, imaging.Linear)
	//src = imaging.AdjustSigmoid(src, 0.6, -7.0)
	//img = imaging.AdjustBrightness(img, -7)
	//src = imaging.AdjustGamma(src, 0.6)
	img := resize.Resize(IMG_SIZE, IMG_SIZE, src, resize.Lanczos3)

	//path := manage_file.GenerateFileName("./tmp", "for_onnx", "jpg")
	//img2.SaveJpeg(&img, path)

	// collect the colors of pixels to different arrays
	red := []float32{}
	green := []float32{}
	blue := []float32{}

	// Extract array of color components of each pixel and destruct them to r, g and b variables.
	// Then it scales these components and appends them to appropriate arrays.
	for y := 0; y < IMG_SIZE; y++ {
		for x := 0; x < IMG_SIZE; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			red = append(red, float32(r/257)/255.0)
			green = append(green, float32(g/257)/255.0)
			blue = append(blue, float32(b/257)/255.0)
		}
	}

	input := append(red, green...)
	input = append(input, blue...)

	return input, img_width, img_height

}

func runModel(pathToModel string, input []float32) ([]float32, error) {

	inputShape := ort.NewShape(1, 3, IMG_SIZE, IMG_SIZE)
	inputTensor, _ := ort.NewTensor(inputShape, input)

	outputShape := ort.NewShape(1, 4+COUNT_CLASSES, 8400) // todo 84 ?
	outputTensor, _ := ort.NewEmptyTensor[float32](outputShape)

	model, _ := ort.NewSession[float32](pathToModel,
		[]string{"images"}, []string{"output0"},
		[]*ort.Tensor[float32]{inputTensor}, []*ort.Tensor[float32]{outputTensor})

	err := model.Run()
	if err != nil {
		return nil, err
	}
	return outputTensor.GetData(), nil
}

// Returns Array of bounding boxes in format [[x1,y1,x2,y2,object_type,probability],..]
func process_output(output []float32, img_width, img_height int64) []box {
	boxes := []box{}
	for index := 0; index < 8400; index++ {
		class_id, prob := 0, float32(0.0)
		for col := 0; col < COUNT_CLASSES; col++ {
			if output[8400*(col+4)+index] > prob {
				prob = output[8400*(col+4)+index]
				class_id = col
			}
		}
		if prob < PROB_MIN {
			continue
		}

		xc := output[index] // center
		yc := output[8400+index]
		w := output[2*8400+index]
		h := output[3*8400+index]
		x1 := (xc - w/2) / IMG_SIZE * float32(img_width)
		y1 := (yc - h/2) / IMG_SIZE * float32(img_height)
		x2 := (xc + w/2) / IMG_SIZE * float32(img_width)
		y2 := (yc + h/2) / IMG_SIZE * float32(img_height)
		label := yolo_classes[class_id]

		b := box{float64(x1), float64(y1), float64(x2), float64(y2), label, prob}
		boxes = append(boxes, b)
	}

	result := nmsFilter(boxes)
	return result

}

//	non-maximum suppression (NMS)

// удаляем лишние боксы
func nmsFilter(boxes []box) []box {
	// сортируем боксы по убыванию вероятности
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i].prob > boxes[j].prob
	})

	result := []box{}
	for _, b := range boxes {
		keepBox := true
		for _, selectedBox := range result {
			iouBox := iou(b, selectedBox)
			if iouBox >= IOU_LIMIT {
				keepBox = false
				break
			}
		}
		if keepBox {
			result = append(result, b)
		}
	}
	return result
}

// -------------- IoU ----------------

func get10Percent(v float64) float64 {
	return v * 0.1
}

func get20Percent(v float64) float64 {
	return v * 0.2
}

func normalizeChannel(channel []float32) []float32 {
	var sum, sumSquared float64
	length := float64(len(channel))

	for _, value := range channel {
		sum += float64(value)
		sumSquared += float64(value) * float64(value)
	}

	mean := sum / length
	variance := sumSquared/length - mean*mean
	stddev := float32(math.Sqrt(variance))

	normalized := make([]float32, len(channel))
	for i, value := range channel {
		normalized[i] = (value - float32(mean)) / stddev
	}

	return normalized
}
