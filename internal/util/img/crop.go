package img

import (
	"image"
)

const MIN_SIZE = 800

// CropToSquare crops image to center square
func CropToSquare(img image.Image, percent float64) image.Image {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	//if w < size || h < size {
	//	// Если изображение меньше требуемого размера, не обрезаем его
	//	var buf bytes.Buffer
	//	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100}); err != nil {
	//		log.Fatalf("Error encoding image: %v", err)
	//	}
	//	return buf.Bytes()
	//}

	// Определяем меньшую сторону изображения
	minSide := min(w, h)

	// Уменьшаем меньшую сторону на percent %
	size := int(float64(minSide) * percent)

	if size < MIN_SIZE {
		size = MIN_SIZE
	}

	// Вычисляем координаты центральной части изображения
	x0 := (w - size) / 2
	y0 := (h - size) / 2

	subImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(x0, y0, x0+size, y0+size))

	return subImg
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
