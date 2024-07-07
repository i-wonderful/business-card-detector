package service

import (
	"image"
	"log"
	"os"
	"time"

	"card_detector/internal/model"
	"card_detector/internal/service/box_merge"
	boxes_card "card_detector/internal/util/boxes"
	"card_detector/internal/util/img"
)

type ImgPreparer interface {
	Rotage(imgFile *os.File) (image.Image, string)
	CropCard(img image.Image, boxes []model.TextArea) (image.Image, error)
}

type RecognizerFull interface {
	RecognizeImg(img *image.Image) ([]model.DetectWorld, error)
}

type CardDetector interface {
	Detect(img image.Image) ([]model.TextArea, error)
}

type Detector2 struct {
	imgPreparer          ImgPreparer
	textRecognizeService RecognizerFull
	cardDetector         CardDetector
	fieldSorterService   FieldSorter
	cardRepo             CardRepo
	storageFolder        string
	isLogTime            bool
	isDebug              bool
}

func NewDetector2(
	imgPreparer ImgPreparer,
	textRecognizeService RecognizerFull,
	cardDetector CardDetector,
	fieldSortService FieldSorter,
	cardRepo CardRepo,
	storageFolder string,
	isLogTime, isDebug bool) *Detector2 {
	return &Detector2{
		imgPreparer:          imgPreparer,
		textRecognizeService: textRecognizeService,
		cardDetector:         cardDetector,
		fieldSorterService:   fieldSortService,
		cardRepo:             cardRepo,
		storageFolder:        storageFolder,
		isLogTime:            isLogTime,
		isDebug:              isDebug,
	}
}

func (d *Detector2) Detect(imgPath string) (*model.Person, string, error) {
	if d.isLogTime {
		start := time.Now()
		defer func() {
			log.Printf(">>> Time full detection: %s", time.Since(start))
		}()
	}

	// ----------------------
	file, err := os.Open(imgPath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()
	// ----------------------

	// 1. возможный поворот
	im, _ := d.imgPreparer.Rotage(file)

	// 2. find card items
	boxes, err := d.cardDetector.Detect(im)
	if err != nil {
		return nil, "", err
	}
	boxes = boxes_card.MergeCardBoxes(boxes)

	if d.isDebug {
		log.Println("Detected: ")
		for _, box := range boxes {
			log.Println(box)
		}
	}

	// 3.
	im, _ = d.imgPreparer.CropCard(im, boxes)

	detectWorlds, err := d.textRecognizeService.RecognizeImg(&im)
	if err != nil {
		return nil, "", err
	}

	detectWorlds = box_merge.MergeBoxes(detectWorlds)
	if d.isDebug {
		log.Println("Recognized: ")
		for _, world := range detectWorlds {
			log.Printf("%s - %f", world.Text, world.Prob)
		}
	}

	worlds := getOnlyWorlds(detectWorlds)

	p := d.fieldSorterService.Sort(worlds)

	//manage_file.ClearFolder(d.storageFolder)

	person := model.NewPerson(p)
	card := mapCard(*person, "")
	if err := d.cardRepo.Save(card); err != nil {
		log.Println("Error saving card:", err)
	}

	//boxesPath := img.DrawBoxes(im, boxes, d.storageFolder)
	boxesPath := img.DrawTextAndItemsBoxes(im, detectWorlds, boxes, d.storageFolder)

	return person, boxesPath, nil
}

func getOnlyWorlds(detectWorlds []model.DetectWorld) []string {

	worlds := make([]string, len(detectWorlds))
	for i, world := range detectWorlds {
		worlds[i] = world.Text
	}
	return worlds

}
