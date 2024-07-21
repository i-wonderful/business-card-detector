package helper

import (
	. "card_detector/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNearestWorld(t *testing.T) {
	tests := []struct {
		name     string
		data     []DetectWorld
		box      TextArea
		expected DetectWorld
		expIndex int
	}{
		{
			name: "No objects in Y range",
			data: []DetectWorld{
				{Box: Rectangle{PTop1: Point{X: 100, Y: 300}, PBot1: Point{X: 150, Y: 350}}},
				{Box: Rectangle{PTop1: Point{X: 200, Y: 400}, PBot1: Point{X: 250, Y: 450}}},
			},
			box: TextArea{X: 50, Y: 100, Height: 50},
			//expected: nil,
			expIndex: -1,
		},
		{
			name: "Single object in Y range",
			data: []DetectWorld{
				{Box: Rectangle{PTop1: Point{X: 100, Y: 100}, PBot1: Point{X: 150, Y: 150}}},
			},
			box:      TextArea{X: 50, Y: 100, Height: 50},
			expected: DetectWorld{Box: Rectangle{PTop1: Point{X: 100, Y: 100}, PBot1: Point{X: 150, Y: 150}}},
			expIndex: 0,
		},
		{
			name: "Multiple objects, one nearest",
			data: []DetectWorld{
				{Text: "world1", Box: Rectangle{PTop1: Point{X: 200, Y: 100}, PBot1: Point{X: 250, Y: 150}}},
				{Text: "world2", Box: Rectangle{PTop1: Point{X: 120, Y: 100}, PBot1: Point{X: 170, Y: 150}}},
				{Text: "world3", Box: Rectangle{PTop1: Point{X: 150, Y: 100}, PBot1: Point{X: 200, Y: 150}}},
			},
			box:      TextArea{X: 50, Y: 100, Height: 50},
			expected: DetectWorld{Text: "world2", Box: Rectangle{PTop1: Point{X: 120, Y: 100}, PBot1: Point{X: 170, Y: 150}}},
			expIndex: 1,
		},
		{
			name: "Sebastian skype",
			data: []DetectWorld{
				{Text: "sebastian.jepsson86", Box: NewBoxFromPoints(
					Point{161, 554},
					Point{390, 571},
					Point{388, 598},
					Point{159, 582})},
			},
			box: TextArea{X: 111, Y: 543, Width: 41, Height: 46},
			expected: DetectWorld{Text: "sebastian.jepsson86", Box: NewBoxFromPoints(
				Point{161, 554},
				Point{390, 571},
				Point{388, 598},
				Point{159, 582})},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, index := FindNearestWorld(tt.data, &tt.box)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.expIndex, index)
		})
	}
}
