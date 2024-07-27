package img

import (
	r "github.com/nfnt/resize"
	"image"
	"math"
)

const LONG_SIDE = 1000

func Resize(img image.Image) image.Image {

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	//longSide := 1000
	var newWidth, newHeight uint
	if width > height {
		newWidth = uint(LONG_SIDE)
		newHeight = 0
	} else {
		newWidth = 0
		newHeight = uint(LONG_SIDE)
	}

	resizedImg := r.Resize(newWidth, newHeight, img, r.Lanczos3)

	return resizedImg
}

// ResizeImages scaled and resizes 'im' return a new `image.Image` instance whose maximum dimension is 'max'.
func ResizeWith(img image.Image, max int) image.Image {

	bounds := img.Bounds()
	dims := bounds.Max

	width := dims.X
	height := dims.Y

	ratio_w := float64(max) / float64(width)
	ratio_h := float64(max) / float64(height)

	ratio := math.Min(ratio_w, ratio_h)

	w := uint(float64(width) * ratio)
	h := uint(float64(height) * ratio)

	sm := r.Thumbnail(w, h, img, r.Lanczos3)

	return sm
}

func ResizeImageByHeight(img image.Image, h int) image.Image {
	bounds := img.Bounds()
	w, hImg := bounds.Max.X, bounds.Max.Y

	// Если изображение уже меньше или равно заданной высоте, ничего не делаем
	if hImg <= h {
		return img
	}

	// Вычисляем новую ширину, чтобы сохранить пропорции
	newW := int(math.Round(float64(w) * float64(h) / float64(hImg)))

	// Создаем новое изображение с рассчитанными размерами
	//newImg := imaging.Resize(img, newW, h, imaging.Lanczos)

	newImg := r.Thumbnail(uint(newW), uint(h), img, r.Lanczos3) //  r.Bilinear

	// Заполняем фон белым цветом, чтобы избежать прозрачного фона
	//newImg = imaging.Fill(newImg, bounds.Max.X, bounds.Max.Y, color.White, imaging.Top, imaging.Left)

	return newImg
}

// todo try https://github.com/aaronland/go-image/blob/main/resize/resize.go
