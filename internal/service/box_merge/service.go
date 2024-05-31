package box_merge

import (
	"card_detector/internal/model"
	"math"
	"sort"
	"unicode"
)

// MergeBoxes - merge boxes with close values
func MergeBoxes(boxes []model.DetectWorld) []model.DetectWorld {
	sortByHeight(boxes)

	rez := []model.DetectWorld{}
	prev := &boxes[0]
	for i := 1; i < len(boxes); i++ {
		if isCloser25Values(prev.Box.H, boxes[i].Box.H) &&
			isCloser2Values(prev.Box.PTop1.X, boxes[i].Box.PTop1.X) &&
			isOnlyLetters(prev.Text) && isOnlyLetters(boxes[i].Text) {
			// merge items
			rez = append(rez, model.DetectWorld{
				Text: prev.Text + " " + boxes[i].Text,
				Box: model.NewBoxFromPoints(
					prev.Box.PTop1,
					prev.Box.PBot1,
					boxes[i].Box.PTop2,
					boxes[i].Box.PBot2,
				),
				Prob: (prev.Prob + boxes[i].Prob) / 2,
			})
			rez = append(rez, boxes[i+1:]...)
			prev = nil
			//if i+1 == len(boxes) {
			break
			//rez = append(rez, boxes[i])
			//}
			//prev = boxes[i+1]
			//i++
		} else {
			rez = append(rez, *prev)
			prev = &boxes[i]
		}
	}
	if prev != nil {
		rez = append(rez, *prev)
	}

	return rez
}

func sortByHeight(worlds []model.DetectWorld) []model.DetectWorld {
	sort.Slice(worlds, func(i, j int) bool {
		w1 := worlds[i]
		w2 := worlds[j]
		if isCloser25Values(w1.Box.H, w2.Box.H) {
			return w1.Box.PTop1.Y < w2.Box.PTop1.Y
		}
		return w1.Box.H > w2.Box.H
	})
	return worlds
}

func isContainsDigits(val string) bool {
	for _, r := range val {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func isOnlyLetters(val string) bool {

	for _, r := range val {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func sortByProbAndY(worlds []model.DetectWorld) []model.DetectWorld {
	sort.Slice(worlds, func(i, j int) bool {
		w1 := worlds[i]
		w2 := worlds[j]
		if math.Abs(float64(w1.Prob-w2.Prob)) > 0.04 {
			return w1.Prob > w2.Prob
		}
		return w1.Box.PTop1.Y < w2.Box.PTop1.Y
	})
	return worlds
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x

}

// Числа отличаются не более чем на 25%
func isCloser25Values(a, b int) bool {
	diff := math.Abs(float64(a - b))
	limitPercent := 0.25 * float64(max(a, b))
	return diff <= limitPercent
}

func isCloser2Values(a, b int) bool {
	diff := math.Abs(float64(a - b))
	limitPercent := 0.02 * float64(max(a, b))
	return diff <= limitPercent
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
