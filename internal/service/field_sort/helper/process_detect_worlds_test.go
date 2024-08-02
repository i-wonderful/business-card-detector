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
			result, index := FindNearestWorldToBox(tt.data, &tt.box)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.expIndex, index)
		})
	}
}

func TestFindNearestByY(t *testing.T) {
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
						W:     152,
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
					W:     259,
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
						H:     55,
						W:     226,
					},
				},
			},
			expected:        "KILOVA",
			expectedIsAbove: false,
		},
		{
			name: "Marketing (problem5 1_12)",
			item: DetectWorld{
				Text: "Daniele Costa ",
				Box: Rectangle{
					PTop1: Point{376, 180},
					PTop2: Point{596, 183},
					PBot1: Point{595, 215},
					PBot2: Point{376, 212},
					H:     32,
					W:     220,
				},
			},
			worlds: []DetectWorld{
				{
					Text: "Marketing",
					Box: Rectangle{
						PTop1: Point{375, 217},
						PTop2: Point{546, 223},
						PBot1: Point{545, 259},
						PBot2: Point{373, 253},
						H:     36,
						W:     171,
					},
				},
				{
					Text: "PAY FOR FUH",
					Box: Rectangle{
						PTop1: Point{613, 120},
						PTop2: Point{705, 120},
						PBot1: Point{705, 134},
						PBot2: Point{613, 134},
						H:     14,
						W:     92,
					},
				},
			},
			expected:        "Marketing",
			expectedIsAbove: false,
		},
		{
			name: "Arkadijs long job (problem5 0_10)",
			item: DetectWorld{
				Text: "Arkadijs Narcuks",
				Box: Rectangle{
					PTop1: Point{179, 273},
					PTop2: Point{471, 276},
					PBot1: Point{471, 404},
					PBot2: Point{179, 402},
					H:     128,
					W:     292,
				},
			},
			worlds: []DetectWorld{
				{
					Text: "LPB",
					Box: Rectangle{
						PTop1: Point{963, 226},
						PTop2: Point{1168, 226},
						PBot1: Point{1168, 297},
						PBot2: Point{963, 297},
						H:     71,
						W:     205,
					},
				},
				{
					Text: "Klientu vaditajs",
					Box: Rectangle{
						PTop1: Point{180, 434},
						PTop2: Point{460, 437},
						PBot1: Point{460, 472},
						PBot2: Point{180, 470},
						H:     35,
						W:     280,
					},
				},
				{
					Text: "E-komercijas departaments",
					Box: Rectangle{
						PTop1: Point{180, 472},
						PTop2: Point{669, 470},
						PBot1: Point{669, 505},
						PBot2: Point{180, 508},
						H:     35,
						W:     489,
					},
				},
			},
			expected:        "Klientu vaditajs",
			expectedIsAbove: false,
		},
	}

	for _, tc := range testCasesFindNearest {
		t.Run(tc.name, func(t *testing.T) {
			result := FindNearestByY(&tc.item, tc.worlds)
			if assert.NotNil(t, result) {
				assert.Equal(t, tc.expected, result.Text)
				//assert.Equal(t, tc.expectedIsAbove, isAbove)
			}
		})
	}
}

func TestUpOrDown(t *testing.T) {

	tests := []struct {
		name  string
		item1 DetectWorld
		item2 DetectWorld
		want  bool
	}{
		{
			name: "OKSANA KILOVA (problem5 0_1)",
			item1: DetectWorld{
				Text: "OKSANA",
				Box: Rectangle{
					PTop1: Point{189, 446},
					PTop2: Point{448, 444},
					PBot1: Point{449, 498},
					PBot2: Point{190, 501},
					H:     54,
					W:     259,
				},
			},
			item2: DetectWorld{
				Text: "KILOVA",
				Box: Rectangle{
					PTop1: Point{359, 500},
					PTop2: Point{585, 497},
					PBot1: Point{586, 553},
					PBot2: Point{360, 555},
					H:     55,
					W:     226,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsAbove(tt.item1, tt.item2), "IsAbove(%v, %v)", tt.item1, tt.item2)
		})
	}
}
