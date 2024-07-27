package model

import (
	"math"
)

type TextArea struct {
	Label  string
	X      int
	Y      int
	Width  int
	Height int
}

func (b *TextArea) IsVertical() bool {
	return b.Height > b.Width
}

func (b *TextArea) IsSquare() bool {
	d := float64(b.Width) / float64(b.Height)
	return math.Abs(d-1.0) < 0.2
}
