package boxes

import (
	"card_detector/internal/model"
	"reflect"
	"testing"
)

func TestMergeCardBoxes(t *testing.T) {
	tests := []struct {
		name     string
		input    []model.TextArea
		expected []model.TextArea
	}{
		{
			name: "No card boxes",
			input: []model.TextArea{
				{Label: "not_card", X: 0, Y: 0, Width: 10, Height: 10},
				{Label: "also_not_card", X: 20, Y: 20, Width: 5, Height: 5},
			},
			expected: []model.TextArea{
				{Label: "not_card", X: 0, Y: 0, Width: 10, Height: 10},
				{Label: "also_not_card", X: 20, Y: 20, Width: 5, Height: 5},
			},
		},
		{
			name: "Single card box",
			input: []model.TextArea{
				{Label: "card", X: 0, Y: 0, Width: 10, Height: 10},
				{Label: "not_card", X: 20, Y: 20, Width: 5, Height: 5},
			},
			expected: []model.TextArea{
				{Label: "not_card", X: 20, Y: 20, Width: 5, Height: 5},
				{Label: "card", X: 0, Y: 0, Width: 10, Height: 10},
			},
		},
		{
			name: "Multiple card boxes",
			input: []model.TextArea{
				{Label: "card", X: 0, Y: 0, Width: 10, Height: 10},
				{Label: "not_card", X: 20, Y: 20, Width: 5, Height: 5},
				{Label: "card", X: 5, Y: 5, Width: 15, Height: 15},
			},
			expected: []model.TextArea{
				{Label: "not_card", X: 20, Y: 20, Width: 5, Height: 5},
				{Label: "card", X: 0, Y: 0, Width: 20, Height: 20},
			},
		},
		{
			name: "Overlapping card boxes",
			input: []model.TextArea{
				{Label: "card", X: 0, Y: 0, Width: 10, Height: 10},
				{Label: "card", X: 5, Y: 5, Width: 10, Height: 10},
			},
			expected: []model.TextArea{
				{Label: "card", X: 0, Y: 0, Width: 15, Height: 15},
			},
		},
		{
			name: "Only card boxes",
			input: []model.TextArea{
				{Label: "card", X: 0, Y: 0, Width: 10, Height: 10},
				{Label: "card", X: 20, Y: 20, Width: 5, Height: 5},
			},
			expected: []model.TextArea{
				{Label: "card", X: 0, Y: 0, Width: 25, Height: 25},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeCardBoxes(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("mergeCardBoxes() = %v, want %v", result, tt.expected)
			}
		})
	}
}
