package model

type DetectWorld struct {
	Text string
	Rect Rectangle
	Prob float32
}

type Rectangle struct {
	PTop1 Point
	PTop2 Point
	PBot1 Point
	PBot2 Point
}

func (r *Rectangle) GetHeight() int {
	return r.PTop1.Y - r.PBot1.Y
}

type Point struct {
	X int
	Y int
}
