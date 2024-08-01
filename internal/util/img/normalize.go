package img

import (
	"image"
	"image/color"
)

func NormalizeImage(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	normalizedImg := image.NewRGBA(bounds)

	var min, max uint32 = 65535, 0 // Using the max possible value for 16-bit color depth

	// First pass to find min and max pixel values
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r < min {
				min = r
			}
			if g < min {
				min = g
			}
			if b < min {
				min = b
			}
			if r > max {
				max = r
			}
			if g > max {
				max = g
			}
			if b > max {
				max = b
			}
		}
	}

	// Second pass to normalize pixel values
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			nr := uint8((r - min) * 255 / (max - min))
			ng := uint8((g - min) * 255 / (max - min))
			nb := uint8((b - min) * 255 / (max - min))

			normalizedImg.Set(x, y, color.RGBA{R: nr, G: ng, B: nb, A: uint8(a >> 8)})
		}
	}

	return normalizedImg
}
