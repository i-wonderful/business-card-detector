package boxes

import (
	"card_detector/internal/model"
	"card_detector/internal/service/detect/onnx"
)

// MergeCardBoxes - combines all card objects into one card
func MergeCardBoxes(boxes []model.TextArea) []model.TextArea {
	var cardBox *model.TextArea
	result := make([]model.TextArea, 0, len(boxes))

	for _, box := range boxes {
		if box.Label == onnx.CARD_CLASS {
			if cardBox == nil {
				cardBox = &model.TextArea{
					Label:  onnx.CARD_CLASS,
					X:      box.X,
					Y:      box.Y,
					Width:  box.Width,
					Height: box.Height,
				}
			} else {
				// Расширяем cardBox, чтобы включить текущий box
				minX := min(cardBox.X, box.X)
				minY := min(cardBox.Y, box.Y)
				maxX := max(cardBox.X+cardBox.Width, box.X+box.Width)
				maxY := max(cardBox.Y+cardBox.Height, box.Y+box.Height)

				cardBox.X = minX
				cardBox.Y = minY
				cardBox.Width = maxX - minX
				cardBox.Height = maxY - minY
			}
		} else {
			result = append(result, box)
		}
	}

	if cardBox != nil {
		result = append(result, *cardBox)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
