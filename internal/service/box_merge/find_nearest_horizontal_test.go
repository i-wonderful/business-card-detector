package box_merge

import (
	. "card_detector/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNearestHorizontal(t *testing.T) {
	tests := []struct {
		name     string
		item     DetectWorld
		worlds   []DetectWorld
		expected *DetectWorld
	}{
		{
			name: "Email nearest",
			item: DetectWorld{
				Text: "Email",
				Box: Rectangle{
					PTop1: Point{X: 118, Y: 676},
					PTop2: Point{X: 210, Y: 676},
					PBot1: Point{X: 210, Y: 706},
					PBot2: Point{X: 118, Y: 706},
					H:     30},
			},
			worlds: []DetectWorld{
				{
					Text: "Website",
					Box: Rectangle{
						PTop1: Point{X: 124, Y: 757},
						PTop2: Point{X: 255, Y: 757},
						PBot1: Point{X: 255, Y: 786},
						PBot2: Point{X: 124, Y: 786},
						H:     29}},
				{
					Text: "dz@vallettapay.com",
					Box: Rectangle{
						PTop1: Point{X: 301, Y: 674},
						PTop2: Point{X: 617, Y: 674},
						PBot1: Point{X: 617, Y: 702},
						PBot2: Point{X: 301, Y: 702},
						H:     28}},
			},
			expected: &DetectWorld{
				Text: "dz@vallettapay.com",
				Box: Rectangle{
					PTop1: Point{X: 301, Y: 674},
					PTop2: Point{X: 617, Y: 674},
					PBot1: Point{X: 617, Y: 702},
					PBot2: Point{X: 301, Y: 702},
					H:     28}},
		},
		{
			"Phone find",
			DetectWorld{
				Text: "Phone",
				Box: Rectangle{
					PTop1: Point{X: 115, Y: 633},
					PTop2: Point{X: 219, Y: 636},
					PBot1: Point{X: 218, Y: 667},
					PBot2: Point{X: 114, Y: 663},
					H:     30},
			},
			[]DetectWorld{
				{
					Text: "+35679707069",
					Box: Rectangle{
						PTop1: Point{X: 294, Y: 631},
						PTop2: Point{X: 570, Y: 626},
						PBot1: Point{X: 571, Y: 657},
						PBot2: Point{X: 294, Y: 663},
						H:     31},
				},
			},
			&DetectWorld{
				Text: "+35679707069",
				Box: Rectangle{
					PTop1: Point{X: 294, Y: 631},
					PTop2: Point{X: 570, Y: 626},
					PBot1: Point{X: 571, Y: 657},
					PBot2: Point{X: 294, Y: 663},
					H:     31},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := findNearestHorizontal(tt.item, tt.worlds)
			assert.NotNil(t, result, "expected %v, got nil", tt.expected)
			assert.Equal(t, tt.expected, result, "expected %v, got %v", tt.expected, result)
		})
	}
}
