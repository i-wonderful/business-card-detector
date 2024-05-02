package img

import (
	"image"
	"log"
	"os"
)

func OpenImg(path string) (image.Image, error) {
	imgFile, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening image file: %v", err)
		return nil, err
	}
	defer imgFile.Close()
	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Printf("Error decoding image: %v", err)
		return nil, err
	}
	return img, nil
}
