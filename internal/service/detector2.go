package service

import (
	"card_detector/internal/model"
	"log"
	"os"
	"path/filepath"
	"time"
)

type RecognizerFull interface {
	RecognizeAll(path string) ([]model.DetectWorld, error)
}

type Detector2 struct {
	imgPreparer          ImgPreparer
	textRecognizeService RecognizerFull
	fieldSorterService   FieldSorter
	cardRepo             CardRepo
	storageFolder        string
	isLogTime            bool
	isDebug              bool
}

func NewDetector2(
	imgPreparer ImgPreparer,
	textRecognizeService RecognizerFull,
	fieldSortService FieldSorter,
	cardRepo CardRepo,
	storageFolder string,
	isLogTime, isDebug bool) *Detector2 {
	return &Detector2{
		imgPreparer:          imgPreparer,
		textRecognizeService: textRecognizeService,
		fieldSorterService:   fieldSortService,
		cardRepo:             cardRepo,
		storageFolder:        storageFolder,
		isLogTime:            isLogTime,
		isDebug:              isDebug,
	}
}

func (d *Detector2) Detect(imgPath string) (*model.Person, error) {
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

	_, currentPath := d.imgPreparer.Prepare(file)

	absPath, _ := filepath.Abs(currentPath)

	detectWorlds, err := d.textRecognizeService.RecognizeAll(absPath)
	if err != nil {
		return nil, err
	}

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
	return person, nil
}

func getOnlyWorlds(detectWorlds []model.DetectWorld) []string {

	worlds := make([]string, len(detectWorlds))
	for i, world := range detectWorlds {
		worlds[i] = world.Text
	}
	return worlds

}
