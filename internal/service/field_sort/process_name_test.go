package field_sort

import (
	. "card_detector/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNearest(t *testing.T) {
	testCasesFindNearest := []struct {
		name            string
		item            DetectWorld
		worlds          []DetectWorld
		expected        string
		expectedIsAbove bool
	}{
		{
			name: "Marinov Angel",
			item: DetectWorld{
				Text: "Marinov",
				Box: Rectangle{
					PTop1: Point{73, 349},
					PTop2: Point{312, 349},
					PBot1: Point{312, 395},
					PBot2: Point{73, 395},
					H:     46,
				},
			},
			worlds: []DetectWorld{
				{
					Text: "Angel",
					Box: Rectangle{
						PTop1: Point{69, 284},
						PTop2: Point{221, 279},
						PBot1: Point{223, 342},
						PBot2: Point{73, 395},
						H:     46,
					},
				},
			},
			expected:        "Angel",
			expectedIsAbove: true,
		}, {
			name: "OKSANA KILOVA",
			item: DetectWorld{
				Text: "OKSANA",
				Box: Rectangle{
					PTop1: Point{189, 446},
					PTop2: Point{448, 444},
					PBot1: Point{449, 498},
					PBot2: Point{190, 501},
					H:     54,
				},
			},
			worlds: []DetectWorld{
				{
					Text: "KILOVA",
					Box: Rectangle{
						PTop1: Point{359, 500},
						PTop2: Point{585, 497},
						PBot1: Point{586, 553},
						PBot2: Point{360, 555},
						H:     55},
				},
			},
			expected:        "KILOVA",
			expectedIsAbove: false,
		},
	}

	for _, tc := range testCasesFindNearest {
		t.Run(tc.name, func(t *testing.T) {
			result, isAbove := findNearest(&tc.item, tc.worlds)
			if assert.NotNil(t, result) {
				assert.Equal(t, tc.expected, result.Text)
				assert.Equal(t, tc.expectedIsAbove, isAbove)
			}
		})
	}
}
