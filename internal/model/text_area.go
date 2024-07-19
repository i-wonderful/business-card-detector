package model

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
