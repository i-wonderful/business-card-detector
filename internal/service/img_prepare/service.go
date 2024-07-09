package img_prepare

import (
	"image"
	"log"
	"os"

	"card_detector/internal/model"
	"card_detector/internal/service/detect/onnx"
	boxes_util "card_detector/internal/util/boxes"
	"card_detector/internal/util/img"
	aa_imaging "github.com/aaronland/go-image/imaging"
	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

// ------------------------
// Rotage image for OCR
// ------------------------

const ORIENTATION_NONE = "1"

type Service struct {
	storageFolder string
}

func NewService(storageFolder string) *Service {
	return &Service{storageFolder: storageFolder}
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

	// увеличить контраст
	//im = imaging.AdjustContrast(im, 20)

	// резкость (???)
	//im = imaging.Sharpen(im, 0.36)
	// светлость
	im = imaging.AdjustGamma(im, 1.6)
	im = imaging.AdjustBrightness(im, -10)
	//im = img.BinarizeImage(im, 80)

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
func (s *Service) CropCard(im image.Image, boxes []model.TextArea) (image.Image, error) {
	for _, box := range boxes {
		if box.Label == onnx.CARD_CLASS {
			subImg, offsetX, offsetY := img.CropSquare(im, box.X, box.Y, box.Width, box.Height)

			boxes_util.Transpose(boxes, offsetX, offsetY)
			return subImg, nil
		}
	}
	return im, nil
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
