package boxes

import "card_detector/internal/model"

// Scaling - scaling boxes after resize img
func Scaling(boxes []model.TextArea, scaleX, scaleY float64) {
	for i := range boxes {
		boxes[i].X = (int)(float64(boxes[i].X) * scaleX)
		boxes[i].Y = (int)(float64(boxes[i].Y) * scaleY)
		boxes[i].Width = (int)(float64(boxes[i].Width) * scaleX)
		boxes[i].Height = (int)(float64(boxes[i].Height) * scaleY)
	}
}
