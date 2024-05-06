package img

import (
	"image"
	"image/color"
)

// BinarizeImage - Бинаризация это преобразование изображения в черно-белый формат,
// где каждый пиксель имеет значение 0 (черный) или 255 (белый).
func BinarizeImage(src image.Image, threshold uint8) *image.Gray {
	bounds := src.Bounds()
	gray := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := src.At(x, y).RGBA()
			grayValue := uint8((0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 256)

			if grayValue < threshold {
				gray.Set(x, y, color.Gray{0})
			} else {
				gray.Set(x, y, color.Gray{255})
			}
		}
	}

	return gray
}

// Увеличение изображения в 2 раза
//subImage = imaging.Resize(subImage, subImage.Bounds().Dx()*2, subImage.Bounds().Dy()*2, imaging.Lanczos)
