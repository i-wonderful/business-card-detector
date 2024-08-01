package img

import (
	"golang.org/x/image/draw"
	"image"
)

func ResizeImage(img image.Image, minSize int) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	if width >= minSize && height >= minSize {
		return img
	}

	var newWidth, newHeight int
	if width < height {
		newWidth = minSize
		newHeight = height * minSize / width
	} else {
		newHeight = minSize
		newWidth = width * minSize / height
	}

	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.ApproxBiLinear.Scale(newImg, newImg.Bounds(), img, bounds, draw.Over, nil)
	//draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, bounds, draw.Over, nil)

	return newImg
}

func ScaleToSquare(img image.Image, size int) image.Image {
	bounds := img.Bounds()
	//width, height := bounds.Dx(), bounds.Dy()

	newImg := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.ApproxBiLinear.Scale(newImg, newImg.Bounds(), img, bounds, draw.Over, nil)

	return newImg

}
