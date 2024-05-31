package model

type DetectWorld struct {
	Text string
	Box  Rectangle
	Prob float32
}

type Rectangle struct {
	PTop1 Point
	PTop2 Point
	PBot1 Point
	PBot2 Point
	H     int
	W     int
}

func NewBoxFromPoints(p1, p2, p3, p4 Point) Rectangle {
	h1 := p4.Y - p1.Y
	h2 := p3.Y - p2.Y
	h := (h1 + h2) / 2
	return Rectangle{
		PTop1: p1,
		PTop2: p2,
		PBot1: p3,
		PBot2: p4,
		H:     h,
	}
}

func (r *Rectangle) GetHeight() int {
	return r.PTop1.Y - r.PBot1.Y
}

type Point struct {
	X int
	Y int
}
