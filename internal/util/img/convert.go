package img

import (
	"bytes"
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

// var buf bytes.Buffer
//	if err := jpeg.Encode(&buf, subImage, nil); err != nil {
//		log.Fatalf("Error encoding subimage: %v", err)
//	}
//	return buf.Bytes()
