package img

import (
	"card_detector/internal/model"
	manage_file "card_detector/internal/util/file"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"log"
	"time"
)

var greenColor = color.RGBA{G: 128, A: 255}

// DrawBoxes - рисует боксы на изображении
// @return путь к сохраненному изображению
func DrawBoxes(im image.Image, boxes []model.TextArea, pathStorage string) string {
	start := time.Now()
	defer func() {
		log.Printf(">>> Time DrawBoxes: %s", time.Since(start))
	}()

	bounds := im.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, im, bounds.Min, draw.Src)

	for _, box := range boxes {
		rect := image.Rect(box.X, box.Y, box.X+box.Width, box.Y+box.Height)
		DrawBox(rgba, rect, color.RGBA{255, 0, 0, 255}, 2, box.Label)
	}

	outputFilePath := manage_file.GenerateFileName(pathStorage, "boxes", "jpg")
	SaveRGBAJpeg(rgba, outputFilePath)
	return outputFilePath
}

func DrawTextBoxes(im image.Image, worlds []model.DetectWorld, pathStorage string) string {
	start := time.Now()
	defer func() {
		log.Printf(">>> Time DrawTextBoxes: %s", time.Since(start))
	}()

	bounds := im.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, im, bounds.Min, draw.Src)

	for _, w := range worlds {
		boxTop := w.Box.PTop1
		boxBottom := w.Box.PBot1
		rect := image.Rect(boxTop.X, boxTop.Y, boxBottom.X, boxBottom.Y)
		DrawBox(rgba, rect, greenColor, 2, "text")
	}

	outputFilePath := manage_file.GenerateFileName(pathStorage, "text_boxes", "jpg")
	SaveRGBAJpeg(rgba, outputFilePath)
	return outputFilePath
}

func DrawTextAndItemsBoxes(im image.Image, worlds []model.DetectWorld, boxes []model.TextArea, pathStorage string) string {
	start := time.Now()
	defer func() {
		log.Printf(">>> Time DrawTextAndItemsBoxes: %s", time.Since(start))
	}()

	bounds := im.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, im, bounds.Min, draw.Src)

	for _, w := range worlds {
		boxTop := w.Box.PTop1
		boxBottom := w.Box.PBot1
		rect := image.Rect(boxTop.X, boxTop.Y, boxBottom.X, boxBottom.Y)
		DrawBox(rgba, rect, greenColor, 2, "text")
	}

	for _, box := range boxes {
		rect := image.Rect(box.X, box.Y, box.X+box.Width, box.Y+box.Height)
		DrawBox(rgba, rect, color.RGBA{255, 0, 0, 255}, 2, box.Label)
	}

	outputFilePath := manage_file.GenerateFileName(pathStorage, "boxes", "jpg")
	SaveRGBAJpeg(rgba, outputFilePath)
	return outputFilePath
}

func DrawBox(img *image.RGBA, rect image.Rectangle, c color.Color, thickness int, label string) {
	for i := 0; i < thickness; i++ {
		// Верхняя линия
		draw.Draw(img, image.Rect(rect.Min.X-i, rect.Min.Y-i, rect.Max.X+i, rect.Min.Y-i+1), &image.Uniform{c}, image.Point{}, draw.Src)
		// Нижняя линия
		draw.Draw(img, image.Rect(rect.Min.X-i, rect.Max.Y+i-1, rect.Max.X+i, rect.Max.Y+i), &image.Uniform{c}, image.Point{}, draw.Src)
		// Левая линия
		draw.Draw(img, image.Rect(rect.Min.X-i, rect.Min.Y-i, rect.Min.X-i+1, rect.Max.Y+i), &image.Uniform{c}, image.Point{}, draw.Src)
		// Правая линия
		draw.Draw(img, image.Rect(rect.Max.X+i-1, rect.Min.Y-i, rect.Max.X+i, rect.Max.Y+i), &image.Uniform{c}, image.Point{}, draw.Src)
	}

	// Добавляем надпись
	point := fixed.Point26_6{
		X: fixed.Int26_6(rect.Min.X * 64),
		Y: fixed.Int26_6((rect.Min.Y - 5) * 64), // Немного выше бокса
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(c),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
