package model

import (
	"fmt"
	"math"
)

type TextArea struct {
	Label  string
	X      int
	Y      int
	Width  int
	Height int
	Prob   float32
}

func (b *TextArea) IsVertical() bool {
	return b.Height > b.Width
}

func (b *TextArea) IsSquare() bool {
	d := float64(b.Width) / float64(b.Height)
	return math.Abs(d-1.0) < 0.2
}

func (b *TextArea) ToString() string {
	return fmt.Sprintf("{%s, prob: %f, x: %d, y: %d, width: %d, height: %d}", b.Label, b.Prob, b.X, b.Y, b.Width, b.Height)
}
