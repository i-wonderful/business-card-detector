package boxes

import "card_detector/internal/model"

func Transpose(boxes []model.TextArea, offsetX, offsetY int) {
	for i, box := range boxes {
		box.X = box.X - offsetX
		box.Y = box.Y - offsetY
		boxes[i] = box
	}
}
