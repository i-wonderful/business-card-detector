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

// CutSquareFromCenter вырезает квадратное изображение из центра прямоугольного изображения.
func CutSquareFromCenter(srcImg image.Image) (image.Image, error) {
	// Определение размеров исходного изображения
	srcBounds := srcImg.Bounds()
	width, height := srcBounds.Dx(), srcBounds.Dy()

	// Вычисление размера стороны квадрата
	squareSize := min(width, height)

	// Вычисление координат верхнего левого угла квадрата
	x := (width - squareSize) / 2
	y := (height - squareSize) / 2

	// Создание нового квадратного изображения
	dstImg := image.NewRGBA(image.Rect(0, 0, squareSize, squareSize))

	// Копирование части исходного изображения в новое квадратное изображение
	draw.Draw(dstImg, dstImg.Bounds(), srcImg, srcImg.Bounds().Min.Add(image.Pt(x, y)), draw.Src)

	return dstImg, nil
}

// CutShapeFromCenter вырезает квадратное или прямоугольное изображение из центра прямоугольного изображения,
// в зависимости от соотношения сторон исходного изображения.
func CutShapeFromCenter(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	aspectRatio := float64(width) / float64(height)
	var newRect image.Rectangle

	if aspectRatio < 1.5 && aspectRatio > (1/1.5) {
		// Если изображение близко к квадрату
		size := width
		if height < width {
			size = height
		}
		x0 := (width - size) / 2
		y0 := (height - size) / 2
		newRect = image.Rect(x0, y0, x0+size, y0+size)
	} else {
		// Если изображение прямоугольное
		if width > height {
			newWidth := int(float64(width) * 0.9)
			x0 := (width - newWidth) / 2
			newRect = image.Rect(x0, 0, x0+newWidth, height)
		} else {
			newHeight := int(float64(height) * 0.9)
			y0 := (height - newHeight) / 2
			newRect = image.Rect(0, y0, width, y0+newHeight)
		}
	}

	newImg := image.NewRGBA(newRect)
	draw.Draw(newImg, newRect, img, newRect.Min, draw.Src)

	return newImg
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
