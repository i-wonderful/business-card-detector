package img_prepare

import (
	"card_detector/internal/util/img"
	aa_imaging "github.com/aaronland/go-image/imaging"
	"github.com/google/uuid"
	"github.com/rwcarlsen/goexif/exif"
	"image"
	"log"
	"os"
	"reflect"
)

// ------------------------
// Prepare image for OCR
// ------------------------

const ORIENTATION_NONE = "1"

type Service struct {
	storageFolder string
}

func NewService(storageFolder string) *Service {
	return &Service{storageFolder: storageFolder}
}

func (s *Service) Prepare(imgFile *os.File) (image.Image, string) {
	orientation := getOrientation(imgFile)

	im, format, err := image.Decode(imgFile)
	if err != nil {
		log.Printf("Error decoding image: %v", err)
		return nil, ""
	}
	log.Println("Image format:", format)

	//im, _ = img.OpenJPEGAsNRGBA(imgFile.Name())
	im, _ = RotateImageWithOrientation(im, orientation)

	// вырезать центр // уберу, иногда картинки внизу фото (
	// im = img.CropToSquare(im, 0.98)

	log.Println("Image type:", reflect.TypeOf(im))
	//im = img.MakeRectMinZero(im) // to NRGBA
	bytes := img.ToBytes(im)
	//im, _ = img.ToImage(bytes) // save image type
	currentFilePath := s.storageFolder + "/" + uuid.New().String() + "." + format
	//img.SaveNRGBA(&im, currentFilePath)
	if err = img.SaveImg(currentFilePath, bytes); err != nil {
		log.Printf("Error saving image: %v", err)
		return nil, ""
	}

	return im, currentFilePath
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
		log.Printf("Error decoding image metadata: %v", err)
		imgFile.Seek(0, 0)
		return ORIENTATION_NONE
	}

	_, err = imgFile.Seek(0, 0) // Устанавливаем курсор в начало файла
	if err != nil {
		log.Printf("Error resetting file cursor: %v", err)
	}

	orientation, err := metaData.Get(exif.Orientation)
	if err != nil {
		log.Printf("Error getting orientation: %v", err)
		return ORIENTATION_NONE
	} else {
		return orientation.String()
	}
}
