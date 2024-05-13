package img

import (
	"bytes"
	"golang.org/x/image/tiff"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
)

func ToImage(b []byte) (image.Image, error) {
	reader := bytes.NewReader(b)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func ToBytes(img image.Image) []byte {
	var buf bytes.Buffer
	// todo add formats
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatalf("Error encoding image: %v", err)
	}
	return buf.Bytes()
}

func MakeRectMinZero(src image.Image) *image.NRGBA {
	// Создаем новое изображение с нулевыми координатами
	bounds := src.Bounds()
	dst := image.NewNRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))

	// Копируем содержимое src в dst
	draw.Draw(dst, dst.Bounds(), src, bounds.Min, draw.Src)

	return dst
}

func ToTiffBytes(im image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)

	opts := &tiff.Options{
		Compression: tiff.Uncompressed,
	}

	err := tiff.Encode(buf, im, opts)
	if err != nil {
		return nil, err
	}

	// Получаем байты TIFF изображения из буфера
	tiffBytes := buf.Bytes()
	return tiffBytes, nil
}

func ToImageTiffFromBytes(tiffBytes []byte) (image.Image, error) {
	reader := bytes.NewReader(tiffBytes)
	tiffImg, err := tiff.Decode(reader)
	if err != nil {
		return nil, err
	}

	return tiffImg, nil
}

func ToTiff(im image.Image) (image.Image, error) {
	bytes, err := ToTiffBytes(im)
	if err != nil {
		return nil, err
	}
	return ToImageTiffFromBytes(bytes)
}

// var buf bytes.Buffer
//	if err := jpeg.Encode(&buf, subImage, nil); err != nil {
//		log.Fatalf("Error encoding subimage: %v", err)
//	}
//	return buf.Bytes()
