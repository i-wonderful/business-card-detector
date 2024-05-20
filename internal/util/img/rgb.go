package img

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func rgbaToYCbCr(src *image.NRGBA) *image.YCbCr {
	bounds := src.Bounds()
	dst := image.NewYCbCr(bounds, image.YCbCrSubsampleRatio420)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := src.At(x, y).RGBA()
			Y, Cb, Cr := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			dst.Y[dst.YOffset(x, y)] = Y
			dst.Cb[dst.COffset(x, y)] = Cb
			dst.Cr[dst.COffset(x, y)] = Cr
		}
	}

	return dst
}

func OpenJPEGAsNRGBA(filename string) (*image.RGBA, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, color.RGBAModel.Convert(img.At(x, y)))
		}
	}

	return rgba, nil
}

func YCbCrToRGBA(src *image.YCbCr) *image.NRGBA {
	bounds := src.Bounds()
	dst := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			iy := src.YOffset(x, y)
			icb := src.COffset(x, y)
			icr := src.COffset(x, y)
			yccColor := color.YCbCr{
				Y:  src.Y[iy],
				Cb: src.Cb[icb],
				Cr: src.Cr[icr],
			}
			r, g, b, _ := yccColor.RGBA()
			rgbColor := color.NRGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: 255,
			}
			dst.Set(x, y, rgbColor)
		}
	}

	return dst
}
