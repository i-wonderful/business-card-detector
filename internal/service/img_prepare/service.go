package img_prepare

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"

	"card_detector/internal/model"
	"card_detector/internal/service/detect/onnx"
	boxes_util "card_detector/internal/util/boxes"
	manage_file "card_detector/internal/util/file"
	"card_detector/internal/util/img"
	"card_detector/pkg/log"
	aa_imaging "github.com/aaronland/go-image/imaging"
)

// ------------------------
// Rotage image for OCR
// ------------------------

const ORIENTATION_NONE = "1"

type Service struct {
	storageFolder string
	tmpFolder     string
	log           *log.Logger
}

func NewService(storageFolder string, tmpFolder string, logger *log.Logger) *Service {
	return &Service{
		storageFolder: storageFolder,
		tmpFolder:     tmpFolder,
		log:           logger,
	}
}

func (s *Service) Rotage(imgPath string) (image.Image, string, error) {
	// ----------------------
	imgFile, err := os.Open(imgPath)
	if err != nil {
		return nil, "", err
	}
	defer imgFile.Close()
	// ----------------------

	orientation := getOrientation(imgFile)

	im, format, err := image.Decode(imgFile)
	if err != nil {
		s.log.Error("Error decoding image: %v", err)
		return nil, "", err
	}

	//im, _ = img.OpenJPEGAsNRGBA(imgFile.Name())
	im, _ = RotateImageWithOrientation(im, orientation)

	if orientation != ORIENTATION_NONE {
		path := manage_file.GenerateFileName(s.tmpFolder, "rotated", format)
		img.Save(&im, format, path)
		imgPath = path
	}

	// резкость (???)
	//im = imaging.Sharpen(im, 0.36)
	// светлость
	//im = imaging.AdjustGamma(im, 1.6)
	// яркость
	//im = imaging.AdjustBrightness(im, -5)
	//im = img.BinarizeImage(im, 80)
	// увеличить контраст
	//im = imaging.AdjustContrast(im, 20)

	//im = img.CutShapeFromCenter(im)
	//path := manage_file.GenerateFileName(s.tmpFolder, "cut", "jpg")
	//img.SaveJpeg(&im, path)

	//----- save rotated image
	//bytes := img.ToBytes(im)
	//currentFilePath := s.storageFolder + "/" + uuid.New().String() + "." + format
	////img.SaveNRGBA(&im, currentFilePath)
	//if err = img.SaveImg(currentFilePath, bytes); err != nil {
	//	log.Printf("Error saving image: %v", err)
	//	return nil, ""
	//}
	// -----

	return im, imgPath, nil
}

// CropCard - crop card by square from image and transpose boxes
func (s *Service) CropCard(im image.Image, boxes []model.TextArea) image.Image {
	for _, box := range boxes {
		if box.Label == onnx.CARD_CLASS {
			if box.IsSquare() {
				im, offsetX, offsetY := img.CutSquareFromCenter(im)
				boxes_util.Transpose(boxes, offsetX, offsetY)
				s.logDebug("Cropped center square", im.Bounds())
				return im
			}

			var padding int //15
			if box.IsVertical() {
				padding = 0
			} else {
				padding = 10
			}
			x := box.X - padding
			y := box.Y - padding
			w := box.Width + padding // ??? calculate
			h := box.Height + padding

			subImg, offsetX, offsetY := img.CropSquare(im, x, y, w, h)
			boxes_util.Transpose(boxes, offsetX, offsetY)

			s.logDebug("Cropped card", subImg.Bounds())
			return subImg
		}
	}
	return im
}

func (s *Service) FillIcons(im image.Image, boxes []model.TextArea) image.Image {
	for _, box := range boxes {
		if box.Label == "skype" || box.Label == "telegram" {
			im = img.FillRectangle(im, box.X, box.Y, box.Width, box.Height)
		}
	}
	return im
}

func (s *Service) ResizeAndSaveForPaddle(im *image.Image, boxes []model.TextArea) (image.Image, string, error) {
	paddleSize := 940 //1800
	oldWidth := (*im).Bounds().Max.X
	oldHeight := (*im).Bounds().Max.Y
	resized := img.ResizeImageByHeight(*im, paddleSize)

	s.logDebug("Resized To H", resized.Bounds())

	resized = img.ResizeImage(resized, paddleSize) // ?? растянуть до нужного размера

	s.logDebug("Scaled", resized.Bounds())
	// светлость
	resized = imaging.AdjustGamma(resized, 1.6)
	// яркость
	resized = imaging.AdjustBrightness(resized, -10)

	resized = imaging.AdjustContrast(resized, 15)
	resized = imaging.AdjustSigmoid(resized, 0.5, -3.0)

	newWidth := resized.Bounds().Max.X
	newHeight := resized.Bounds().Max.Y
	if newWidth != oldWidth || newHeight != oldHeight {
		scaleX := float64(newWidth) / float64(oldWidth)
		scaleY := float64(newHeight) / float64(oldHeight)
		boxes_util.Scaling(boxes, scaleX, scaleY)
	}

	// Сохранение
	format := "jpg"
	if resized.Bounds().Max.X < 600 || resized.Bounds().Max.Y < 600 {
		format = "tiff"
	} else // if resized.Bounds().Max.X < 900 || resized.Bounds().Max.Y < 900
	{
		format = "png"
	}

	filePath := manage_file.GenerateFileName(s.tmpFolder, "for_paddle", format)

	if format == "jpg" {
		img.SaveJpegWithQality(&resized, filePath, 87) // 87
	} else if format == "png" {
		img.SavePng(&resized, filePath)
	} else {
		img.SaveTiff(resized, filePath)
	}
	absPath, _ := filepath.Abs(filePath)

	return resized, absPath, nil
}
func (s *Service) logDebug(m string, r image.Rectangle) {
	s.log.Debug(m,
		log.Field{Key: "x", Value: r.Max.X},
		log.Field{Key: "y", Value: r.Max.Y})
}

// RotateImageWithOrientation will rotate 'im' based on EXIF orientation value defined in 'orientation'.
func RotateImageWithOrientation(im image.Image, orientation string) (image.Image, error) {

	switch orientation {
	case ORIENTATION_NONE:
		// pass
	case "2":
		im = aa_imaging.FlipV(im)
	case "3":
		im = aa_imaging.Rotate180(im)
	case "4":
		im = aa_imaging.Rotate180(aa_imaging.FlipV(im))
	case "5":
		im = aa_imaging.Rotate270(aa_imaging.FlipV(im))
	case "6":
		im = aa_imaging.Rotate270(im)
	case "7":
		im = aa_imaging.Rotate90(aa_imaging.FlipV(im))
	case "8":
		im = aa_imaging.Rotate90(im)
	}

	return im, nil
}

func getOrientation(imgFile *os.File) string {
	metaData, err := exif.Decode(imgFile)

	if err != nil {
		//log.Printf("Error decoding image metadata: %v", err)
		imgFile.Seek(0, 0)
		return ORIENTATION_NONE
	}

	_, err = imgFile.Seek(0, 0) // Устанавливаем курсор в начало файла
	if err != nil {
		fmt.Printf("Error resetting file cursor: %v", err)
	}

	orientation, err := metaData.Get(exif.Orientation)
	if err != nil {
		return ORIENTATION_NONE
	} else {
		return orientation.String()
	}
}
