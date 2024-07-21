package img_prepare

import (
	"card_detector/internal/service/detect/onnx"
	manage_file "card_detector/internal/util/file"
	"github.com/disintegration/imaging"
	"golang.org/x/image/draw"
	"image"
	"log"
	"os"
	"path/filepath"

	"card_detector/internal/model"
	boxes_util "card_detector/internal/util/boxes"
	"card_detector/internal/util/img"
	aa_imaging "github.com/aaronland/go-image/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

// ------------------------
// Rotage image for OCR
// ------------------------

const ORIENTATION_NONE = "1"

type Service struct {
	storageFolder string
	tmpFolder     string
}

func NewService(storageFolder string, tmpFolder string) *Service {
	return &Service{
		storageFolder: storageFolder,
		tmpFolder:     tmpFolder,
	}
}

func (s *Service) Rotage(imgFile *os.File) (image.Image, string) {
	orientation := getOrientation(imgFile)

	im, _, err := image.Decode(imgFile)
	if err != nil {
		log.Printf("Error decoding image: %v", err)
		return nil, ""
	}

	//im, _ = img.OpenJPEGAsNRGBA(imgFile.Name())
	im, _ = RotateImageWithOrientation(im, orientation)

	// резкость (???)
	//im = imaging.Sharpen(im, 0.36)
	// светлость
	im = imaging.AdjustGamma(im, 1.6)
	// яркость
	im = imaging.AdjustBrightness(im, -10)
	//im = img.BinarizeImage(im, 80)
	// увеличить контраст
	//im = imaging.AdjustContrast(im, 20)

	//----- save rotated image
	//bytes := img.ToBytes(im)
	//currentFilePath := s.storageFolder + "/" + uuid.New().String() + "." + format
	////img.SaveNRGBA(&im, currentFilePath)
	//if err = img.SaveImg(currentFilePath, bytes); err != nil {
	//	log.Printf("Error saving image: %v", err)
	//	return nil, ""
	//}
	// -----

	return im, ""
}

// CropCard - crop card by square from image and transpose boxes
func (s *Service) CropCard(im image.Image, boxes []model.TextArea) image.Image {
	for _, box := range boxes {
		if box.Label == onnx.CARD_CLASS {
			var padding int //15
			if box.IsVertical() {
				padding = 0
			} else {
				padding = 10
			}

			subImg, offsetX, offsetY := img.CropSquare(im, box.X-padding, box.Y-padding, box.Width+padding, box.Height+padding)
			boxes_util.Transpose(boxes, offsetX, offsetY)

			return subImg
		}
	}
	return im
}

func (s *Service) ResizeAndSaveForPaddle(im *image.Image, boxes []model.TextArea) (image.Image, string, error) {
	paddleSize := 940 //1800
	oldWidth := (*im).Bounds().Max.X
	oldHeight := (*im).Bounds().Max.Y
	resized := img.ResizeImageByHeight(*im, paddleSize)

	resized = resizeImage(resized, paddleSize)
	resized = imaging.AdjustContrast(resized, 2)
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
		log.Printf("Error resetting file cursor: %v", err)
	}

	orientation, err := metaData.Get(exif.Orientation)
	if err != nil {
		//log.Printf("Error getting orientation: %v", err)
		return ORIENTATION_NONE
	} else {
		return orientation.String()
	}
}

func resizeImage(img image.Image, minSize int) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	if width >= minSize && height >= minSize {
		return img
	}

	var newWidth, newHeight int
	if width < height {
		newWidth = minSize
		newHeight = height * minSize / width
	} else {
		newHeight = minSize
		newWidth = width * minSize / height
	}

	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.ApproxBiLinear.Scale(newImg, newImg.Bounds(), img, bounds, draw.Over, nil)

	return newImg
}
