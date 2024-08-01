package img

import (
	"fmt"
	"image"
)

func Save(img *image.Image, format string, path string) error {
	switch format {
	case "jpg", "jpeg":
		return SaveJpegWithQality(img, path, 89)
	case "png":
		return SavePng(img, path)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}
