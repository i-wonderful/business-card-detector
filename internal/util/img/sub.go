package img

import (
	"image"
)

func GetSubImage(img image.Image, x, y, width, height int) image.Image {
	subImage := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(x, y, x+width, y+height))

	return subImage
}
