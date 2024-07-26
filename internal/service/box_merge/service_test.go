package box_merge

import (
	. "card_detector/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService(t *testing.T) {
	testCases := []struct {
		name     string
		boxes    []DetectWorld
		expected []DetectWorld
	}{
		{
			name: "Skype merge",
			boxes: []DetectWorld{
				{
					Text: "Email",
					Box: Rectangle{
						PTop1: Point{X: 118, Y: 676},
						PTop2: Point{X: 210, Y: 676},
						PBot1: Point{X: 210, Y: 706},
						PBot2: Point{X: 118, Y: 706},
						H:     30,
					},
					Prob: 0.9866112,
				},
				{
					Text: "Skype",
					Box: Rectangle{
						PTop1: Point{X: 119, Y: 713},
						PTop2: Point{X: 218, Y: 716},
						PBot1: Point{X: 217, Y: 750},
						PBot2: Point{X: 117, Y: 746},
						H:     33,
					},
					Prob: 0.9973361,
				},
				{
					Text: "| davidzammitfcca",
					Box: Rectangle{
						PTop1: Point{X: 271, Y: 715},
						PTop2: Point{X: 570, Y: 713},
						PBot1: Point{X: 570, Y: 741},
						PBot2: Point{X: 271, Y: 744},
						H:     28,
					},
					Prob: 0.95386744,
				},
			},
			expected: []DetectWorld{
				{
					Text: "Skype davidzammitfcca",
					Box: Rectangle{
						PTop1: Point{X: 108, Y: 255},
						PTop2: Point{X: 397, Y: 255},
						PBot1: Point{X: 397, Y: 300},
						PBot2: Point{X: 108, Y: 300},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			merged := MergeBoxesVertical(tc.boxes)

			assert.Equal(t, len(tc.expected), len(merged))
			for i, m := range merged {
				assert.Equal(t, tc.expected[i].Text, m.Text)
			}
		})

	}
}
