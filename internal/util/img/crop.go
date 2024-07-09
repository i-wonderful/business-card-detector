package img

import (
	"image"
	"image/draw"
)

const MIN_SIZE = 800

// CropSquare crops image to square with center at rectangles center.
// @param img - image to crop
// @param x - left top rectangle corner x
// @param y - left top rectangle corner y
// @param width - rectangle width
// @param height - rectangle height
// @return cropped image, diff X, diff Y
func CropSquare(img image.Image, x, y, width, height int) (image.Image, int, int) {
	// Вычисляем центр квадрата
	centerCropX := x + (width / 2)
	centerCropY := y + (height / 2)

	// Размер квадрата выбирается как максимальный из ширины и высоты
	squareSize := max(width, height)

	// Убеждаемся, что размер квадрата не превышает размеры изображения
	squareSizeX := min(squareSize, img.Bounds().Dx())
	squareSizeY := min(squareSize, img.Bounds().Dy())

	// Вычисляем координаты верхнего левого угла квадрата
	xCrop := centerCropX - (squareSizeX / 2)
	yCrop := centerCropY - (squareSizeY / 2)

	// Убеждаемся, что координаты внутри границ изображения
	if xCrop < 0 {
		xCrop = 0
	}
	if yCrop < 0 {
		yCrop = 0
	}
	if xCrop+squareSize > img.Bounds().Dx() {
		xCrop = img.Bounds().Dx() - squareSizeX
	}
	if yCrop+squareSize > img.Bounds().Dy() {
		yCrop = img.Bounds().Dy() - squareSizeY
	}
	//return getSubImageWithZeroOrigin(img, xCrop, yCrop, squareSizeX, squareSizeY), xCrop, yCrop
	return getSubImageWithZeroOrigin(img, xCrop, yCrop, squareSizeX, squareSizeY), xCrop, yCrop
}

func getSubImageWithZeroOrigin(img image.Image, x, y, width, height int) *image.RGBA {
	// Создаем новое изображение с нулевыми координатами и нужными размерами
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// Копируем нужную часть исходного изображения прямо в новое изображение
	draw.Draw(dst, dst.Bounds(), img, image.Point{X: x, Y: y}, draw.Src)

	return dst
}

func getSubImage(img image.Image, x, y, width, height int) image.Image {
	rect := image.Rect(x, y, x+width, y+height)
	subImg := image.NewRGBA(rect)
	draw.Draw(subImg, rect, img, image.Point{X: x, Y: y}, draw.Src)
	return subImg
}

// CropToSquareCenter crops image to center square
func CropToSquareCenter(img image.Image, percent float64) image.Image {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
