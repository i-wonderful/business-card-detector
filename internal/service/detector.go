package service

import (
	"card_detector/internal/service/box_merge"
	manage_file "card_detector/internal/util/file"
	"fmt"
	"image"
	"path/filepath"
	"time"

	"card_detector/internal/model"
	boxes_card "card_detector/internal/util/boxes"
	"card_detector/internal/util/img"
	"card_detector/pkg/log"
)

type ImgPreparer interface {
	Rotage(imgPath string) (image.Image, string, error)
	CropCard(img image.Image, boxes []model.TextArea) image.Image
	ResizeAndSaveForPaddle(im *image.Image, boxes []model.TextArea) (image.Image, string, error)
	FillIcons(im image.Image, boxes []model.TextArea) image.Image
}

type RecognizerFull interface {
	RecognizeImg(img *image.Image) ([]model.DetectWorld, error)
	RecognizeImgByPath(pathImg string) ([]model.DetectWorld, error)
}

type CardDetector interface {
	Detect(img image.Image) ([]model.TextArea, error)
}

type FieldSorter interface {
	Sort(data []model.DetectWorld, boxes []model.TextArea) map[string]interface{}
}

type (
	TextFinder interface {
		PredictTextCoord(img image.Image) ([]model.TextArea, error)
	}

	TextRecognizer interface {
		DetectLang(path string)
		RecognizeBatch(contents [][]byte) ([]string, error)
		//	RecognizeByPath(paths []string) ([]string, error)
	}

	CardRepo interface {
		Save(c model.Card) error
	}
)

type Detector2 struct {
	imgPreparer          ImgPreparer
	textRecognizeService RecognizerFull
	cardDetector         CardDetector
	fieldSorterService   FieldSorter
	cardRepo             CardRepo
	storageFolder        string
	tmpFolder            string
	isLogTime            bool
	logger               *log.Logger
}

func NewDetector2(
	imgPreparer ImgPreparer,
	textRecognizeService RecognizerFull,
	cardDetector CardDetector,
	fieldSortService FieldSorter,
	cardRepo CardRepo,
	storageFolder string,
	tmpFolder string,
	isLogTime bool,
	logger *log.Logger) *Detector2 {
	return &Detector2{
		imgPreparer:          imgPreparer,
		textRecognizeService: textRecognizeService,
		cardDetector:         cardDetector,
		fieldSorterService:   fieldSortService,
		cardRepo:             cardRepo,
		storageFolder:        storageFolder,
		tmpFolder:            tmpFolder,
		isLogTime:            isLogTime,
		logger:               logger,
	}
}

func (d *Detector2) Detect(imgPath string) (*model.Person, string, error) {
	if d.isLogTime {
		start := time.Now()
		defer func() {
			d.logger.Info(fmt.Sprintf("Time full detection: %s", time.Since(start)))
		}()
	}
	// 1. возможный поворот
	im, imgPath, err := d.imgPreparer.Rotage(imgPath)
	if err != nil {
		return nil, "", err
	}
	// 2. find card items
	boxes, err := d.cardDetector.Detect(im)
	if err != nil {
		return nil, "", err
	}
	boxes = boxes_card.MergeCardBoxes(boxes)

	d.logger.Debug("Detected boxes: ")
	for _, box := range boxes {
		d.logger.Debug(box.ToString())
	}

	im2, _ := img.OpenImg(imgPath)
	// 3. Prepare image for recognize text: crop card and resize to optimal square for paddle
	im2 = d.imgPreparer.CropCard(im2, boxes)
	boxes, _ = d.cardDetector.Detect(im2)
	im2 = d.imgPreparer.FillIcons(im2, boxes)
	im2, absPath, _ := d.imgPreparer.ResizeAndSaveForPaddle(&im2, boxes)

	// 4. recognize text
	detectWorlds, err := d.textRecognizeService.RecognizeImgByPath(absPath)
	if err != nil {
		return nil, "", err
	}

	// 5. merge text blocks
	detectWorlds = box_merge.MergeBoxesVertical(detectWorlds)

	d.logger.Debug("Recognized worlds: ")
	for _, world := range detectWorlds {
		d.logger.Debug(world.Text)
	}

	// 6. sort text to person item
	p := d.fieldSorterService.Sort(detectWorlds, boxes)

	// 6. save
	person := model.NewPerson(p)

	boxesPath := imgPath //img.DrawTextAndItemsBoxes(im2, detectWorlds, boxes, d.storageFolder)
	card := mapCard(*person, boxesPath, "", filepath.Base(imgPath))
	if err := d.cardRepo.Save(card); err != nil {
		d.logger.Error("Error saving card:", err)
	}

	manage_file.ClearFolder(d.storageFolder)
	manage_file.ClearFolder(d.tmpFolder)

	return person, boxesPath, nil
}

func mapCard(p model.Person, photoUrl, logoUrl, originalName string) model.Card {
	return model.Card{
		Owner:        "admin",
		UploadedAt:   time.Now(),
		PhotoUrl:     photoUrl,
		LogoUrl:      logoUrl,
		OriginalName: originalName,
		Email:        p.Email,
		Site:         p.Site,
		Phone:        p.Phone,
		Name:         p.Name,
		JobTitle:     p.JobTitle,
		Company:      p.Organization,
		Telegram:     p.Telegram,
		Skype:        p.Skype,
		Other:        p.Other,
	}
}
