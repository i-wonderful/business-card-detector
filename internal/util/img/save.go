package img

import (
	"bytes"
	"golang.org/x/image/tiff"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func SaveImg(outputPath string, content []byte) error {
	file, err := os.Create(outputPath)
	if err != nil {
		log.Printf("Error creating image file: %v", err)
		return err
	}
	defer file.Close()

	if _, err = file.Write(content); err != nil {
		log.Printf("Error writing image to file: %v", err)
		return err
	}
	log.Println("Image has been saved:", outputPath)
	return nil
}

func SaveNRGBA(img *image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, *img)
	if err != nil {
		return err
	}

	return nil
}

func SaveRGBAJpeg(img *image.RGBA, outputFilePath string) error {
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatalf("Error encoding image: %v", err)
	}

	return SaveImg(outputFilePath, buf.Bytes())
}

func SaveJpeg(img *image.Image, outputFilePath string) error {
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, *img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatalf("Error encoding image: %v", err)
	}

	return SaveImg(outputFilePath, buf.Bytes())
}

func SaveTiff(im image.Image, filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Кодируем изображение в TIFF формат
	err = tiff.Encode(outFile, im, nil)
	if err != nil {
		return err
	}

	return nil
}

// Save resized image as JPG file
//resizedImgFile, err := os.Create(outputPath)
//if err != nil {
//	log.Fatalf("Error creating resized image file: %v", err)
//}
//defer resizedImgFile.Close()
//
//if err := jpeg.Encode(resizedImgFile, resizedImg, nil); err != nil {
//	log.Fatalf("Error encoding resized image: %v", err)
//}
// fmt.Println("Resized image has been saved:", outputPath)
