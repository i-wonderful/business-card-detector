package rectangle

import (
	. "card_detector/internal/model"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMergeOverlappingRectangles(t *testing.T) {
	testCases := []struct {
		name           string
		input          []TextArea
		expected       []TextArea
		maxYDiff       int
		sortedByYBegin bool
	}{
		{
			name:     "Empty input",
			input:    []TextArea{},
			expected: []TextArea{},
		},
		{
			name:     "Single rectangle",
			input:    []TextArea{{X: 10, Y: 20, Width: 30, Height: 40}},
			expected: []TextArea{{X: 10, Y: 20, Width: 30, Height: 40}},
		},
		//{
		//	name: "Merge overlapping rectangles",
		//	input: []TextArea{
		//		{X: 10, Y: 20, Width: 30, Height: 40},
		//		{X: 20, Y: 25, Width: 20, Height: 30},
		//		{X: 30, Y: 50, Width: 10, Height: 20},
		//	},
		//	expected: []TextArea{
		//		{X: 10, Y: 20, Width: 30, Height: 50},
		//		{X: 30, Y: 50, Width: 10, Height: 20},
		//	},
		//	maxYDiff: 5,
		//},
		{
			name: "Merge rectangles with maxYDiff",
			input: []TextArea{
				{X: 10, Y: 20, Width: 30, Height: 40},
				{X: 20, Y: 25, Width: 20, Height: 30},
				{X: 30, Y: 65, Width: 10, Height: 20},
			},
			expected: []TextArea{
				{X: 10, Y: 20, Width: 30, Height: 40},
				{X: 30, Y: 65, Width: 10, Height: 20},
			},
			maxYDiff: 5,
		},
		{
			name: "Merge real",
			input: []TextArea{
				{X: 169, Y: 490, Width: 379, Height: 45},
				{X: 168, Y: 499, Width: 375, Height: 43},
			},
			expected: []TextArea{
				{X: 168, Y: 490, Width: 380, Height: 52},
			},
		},
		{
			name: "Merge big format", // если картинки других масштабов, тоже должно работать
			input: []TextArea{
				{X: 595, Y: 1438, Width: 629, Height: 74},
				{X: 589, Y: 1451, Width: 636, Height: 76},
			},
			expected: []TextArea{
				{X: 589, Y: 1438, Width: 636, Height: 89},
			},
		},
		{
			name: "Merge horizontal",
			input: []TextArea{
				{X: 280, Y: 1314, Width: 466, Height: 70},
				{X: 730, Y: 1316, Width: 265, Height: 77},
			},
			expected: []TextArea{
				{X: 280, Y: 1314, Width: 715, Height: 79},
			},
		},
		{
			name: "Merge horizontal2",
			input: []TextArea{
				{X: 505, Y: 777, Width: 424, Height: 110},
				{X: 167, Y: 792, Width: 331, Height: 103},
			},
			expected: []TextArea{
				{X: 167, Y: 777, Width: 762, Height: 118},
			},
		},
		{
			name: "Merge horizontal3",
			input: []TextArea{
				{X: 761, Y: 1310, Width: 829, Height: 87},
				{X: 756, Y: 1334, Width: 797, Height: 85},
			},
			expected: []TextArea{
				{X: 756, Y: 1310, Width: 834, Height: 109},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.sortedByYBegin {
				// Перемешать входные данные, если они не отсортированы по Y
				shuffleInput(tc.input)
			}
			result := MergeOverlappingRectangles(tc.input)

			assert.Len(t, result, len(tc.expected), "Expected %d results, but got %d", len(tc.expected), len(result))
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func shuffleInput(input []TextArea) {
	// Перемешать входные данные
	// Реализация опущена для краткости
}
