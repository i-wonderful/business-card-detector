package service

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	. "card_detector/internal/util/img"
	"card_detector/internal/util/rectangle"
	manage_str2 "card_detector/internal/util/str"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"log"
	"os"
	"reflect"
	"time"
)

type (
	ImgPreparer interface {
		Prepare(imgFile *os.File) (image.Image, string)
	}

	TextFinder interface {
		PredictTextCoord(img image.Image) ([]model.TextArea, error)
	}

	TextRecognizer interface {
		DetectLang(path string)
		RecognizeBatch(contents [][]byte) ([]string, error)
		//	RecognizeByPath(paths []string) ([]string, error)
	}

	FieldSorter interface {
		Sort(data []string) map[string]interface{}
	}

	CardRepo interface {
		Save(c model.Card) error
	}
)

type Detector struct {
	imgPreparer          ImgPreparer
	textRecognizeService TextRecognizer
	findTextService      TextFinder
	fieldSorterService   FieldSorter
	cardRepo             CardRepo
	storageFolder        string
	isLogTime            bool
	isDebug              bool
}

func NewDetector(
	imgPreparer ImgPreparer,
	textFindService TextFinder,
	textRecognizeService TextRecognizer,
	fieldSortService FieldSorter,
	cardRepo CardRepo,
	storageFolder string,
	isLogTime bool) *Detector {
	return &Detector{
		imgPreparer:          imgPreparer,
		textRecognizeService: textRecognizeService,
		findTextService:      textFindService,
		fieldSorterService:   fieldSortService,
		cardRepo:             cardRepo,
		storageFolder:        storageFolder,
		isLogTime:            isLogTime,
		isDebug:              false,
	}
}

func (d *Detector) Detect(imgPath string) (*model.Person, error) {
	if d.isLogTime {
		start := time.Now()
		defer func() {
			log.Printf(">>> Time full detection: %s", time.Since(start))
		}()
	}

	// ----------------------
	file, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// ----------------------

	// 1. Prepare image for text recognition
	_, currentFilePath := d.imgPreparer.Prepare(file)

	// 2. Find text area fields
	currentImg, _ := OpenImg(currentFilePath) // OpenJPEGAsNRGBA(currentFilePath) //
	fmt.Println("type current from file:", reflect.TypeOf(currentImg))

	coords, err := d.findTextService.PredictTextCoord(currentImg)
	if err != nil {
		fmt.Println("Error getting prediction:", err)
		return nil, err
	}

	// 2.1 Merge overlapping text areas
	coords = rectangle.MergeOverlappingRectangles(coords)

	// 3. Create subimages with text area
	var imagesWithText [][]byte
	var paths []string
	for i, coord := range coords {
		subImage := GetSubImage(currentImg, coord.X, coord.Y, coord.Width, coord.Height)
		subImage, _ = ToTiff(subImage)

		//fmt.Println("type subImage:", reflect.TypeOf(subImage))
		//subImage = BinarizeImage(subImage, 128)
		//if reflect.TypeOf(subImage) == reflect.TypeOf(&image.YCbCr{}) {
		//	subImage = YCbCrToRGBA(subImage.(*image.YCbCr))
		//}

		// увеличить контраст
		subImage = imaging.AdjustContrast(subImage, 20)
		//fmt.Println("type subImage:", reflect.TypeOf(subImage))
		// резкость (???)
		subImage = imaging.Sharpen(subImage, 0.36)
		// светлость
		subImage = imaging.AdjustGamma(subImage, 1.6)

		//subImage = imaging.AdjustBrightness(subImage, -10)

		subImageBytes, _ := ToTiffBytes(subImage)
		if d.isDebug {
			//SaveImg("./tmp/"+fmt.Sprintf("subimage%d.jpg", i), subImageBytes)
			path := "./tmp/" + fmt.Sprintf("subimage%d.tiff", i)
			SaveTiff(subImage, path)
			paths = append(paths, path)
		}
		imagesWithText = append(imagesWithText, subImageBytes)
	}

	// 4. Recognize text
	//worlds, err := d.textRecognizeService.RecognizeByPath(paths)
	d.textRecognizeService.DetectLang(currentFilePath)
	worlds, err := d.textRecognizeService.RecognizeBatch(imagesWithText)
	if err != nil {
		return nil, err
	}

	// 5. Process text
	worlds = manage_str2.RemoveSubstrings(worlds)

	p := d.fieldSorterService.Sort(worlds)
	person := model.NewPerson(p)

	manage_file.ClearFolder("./tmp")
	manage_file.ClearFolder(d.storageFolder)

	card := mapCard(*person, "")
	if err := d.cardRepo.Save(card); err != nil {
		fmt.Println("Error saving card:", err)
	}
	return person, nil
}

func mapCard(p model.Person, photoUrl string) model.Card {
	return model.Card{
		Owner:      "admin",
		UploadedAt: time.Now(),
		PhotoUrl:   photoUrl,
		Email:      p.Email,
		Site:       p.Site,
		Phone:      p.Phone,
		Name:       p.Name,
		JobTitle:   p.JobTitle,
		Company:    p.Organization,
		Telegram:   p.Telegram,
		Skype:      p.Skype,
		Other:      p.Other,
	}
}
