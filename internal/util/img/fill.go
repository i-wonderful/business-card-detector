package img

import (
	"image"
	"image/draw"
)

// FillRectangle - fills a rectangle on an image with the color at (x, y)
func FillRectangle(img image.Image, x, y, width, height int) image.Image {
	// Get the color at (x, y)
	rectColor := img.At(x, y)

	// Create a new image to draw on
	bounds := img.Bounds()
	rgbaImg := image.NewRGBA(bounds)

	// Copy the original image to the new image
	draw.Draw(rgbaImg, bounds, img, bounds.Min, draw.Src)

	// Define the rectangle to fill
	rect := image.Rect(x, y, x+width, y+height)

	// Fill the rectangle with the color
	draw.Draw(rgbaImg, rect, &image.Uniform{rectColor}, image.Point{}, draw.Src)

	return rgbaImg
}
