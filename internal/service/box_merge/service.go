package box_merge

import (
	. "card_detector/internal/model"
	. "card_detector/internal/util/calc"
	"math"
	"sort"
	"unicode"
)

// MergeBoxes - merge boxes with close values
func MergeBoxes(boxes []DetectWorld) []DetectWorld {
	sortByHeight(boxes)

	rez := []DetectWorld{}
	i := 0

	for i < len(boxes) {
		prev := &boxes[i]
		if i+1 < len(boxes) {
			next := &boxes[i+1]
			isCloser25Percent := isCloser25Values(prev.Box.H, next.Box.H)
			isCloser2Percent := isCloser2Values(prev.Box.PTop1.X, next.Box.PTop1.X)
			if isCloser25Percent && isCloser2Percent &&
				isOnlyLetters(prev.Text) && isOnlyLetters(next.Text) {

				if i+2 < len(boxes) {
					next2 := &boxes[i+2]
					isCloser25Percent2 := isCloser25Values(next.Box.H, next2.Box.H)
					isCloser2Percent2 := isCloser2Values(next.Box.PTop1.X, next2.Box.PTop1.X)
					if isCloser25Percent2 && isCloser2Percent2 &&
						isOnlyLetters(next.Text) && isOnlyLetters(next2.Text) {

						rez = append(rez, DetectWorld{
							Text: prev.Text + " " + next.Text + " " + next2.Text,
							Box: NewBoxFromPoints(
								mergeTopLeftPoints(prev.Box.PTop1, next2.Box.PTop1),
								mergeTopRightPoints(prev.Box.PTop2, next2.Box.PTop2),
								mergeBottomRightPoints(prev.Box.PBot1, next2.Box.PBot1),
								mergeBottomLeftPoints(prev.Box.PBot2, next2.Box.PBot2),
							),
							Prob: (prev.Prob + next.Prob + next2.Prob) / 3,
						})
						i += 3
						continue
					}
				}

				rez = append(rez, DetectWorld{
					Text: prev.Text + " " + next.Text,
					Box: NewBoxFromPoints(
						mergeTopLeftPoints(prev.Box.PTop1, next.Box.PTop1),
						mergeTopRightPoints(prev.Box.PTop2, next.Box.PTop2),
						mergeBottomRightPoints(prev.Box.PBot1, next.Box.PBot1),
						mergeBottomLeftPoints(prev.Box.PBot2, next.Box.PBot2),
					),
					Prob: (prev.Prob + next.Prob) / 2,
				})
				i += 2
			} else {
				rez = append(rez, *prev)
				i++
			}
		} else {
			rez = append(rez, *prev)
			i++
		}
	}
	return rez
}

func sortByHeight(worlds []DetectWorld) []DetectWorld {
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
		if !unicode.IsLetter(r) /*&& r != ' ' && r != '-' && r != ',' && r != '.'*/ { // todo ???
			return false
		}
	}
	return true
}

func sortByProbAndY(worlds []DetectWorld) []DetectWorld {
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

// Числа отличаются не более чем на 25%
func isCloser25Values(a, b int) bool {
	diff := math.Abs(float64(a - b))
	limitPercent := 0.25 * float64(Max(a, b))
	return diff <= limitPercent
}

func isCloser2Values(a, b int) bool {
	diff := math.Abs(float64(a - b))
	limitPercent := 0.02 * float64(Max(a, b)) // todo 0.03 (?)
	return diff <= limitPercent
}

func mergeTopLeftPoints(a, b Point) Point {
	return Point{
		X: Min(a.X, b.X),
		Y: Min(a.Y, b.Y),
	}
}

func mergeTopRightPoints(a, b Point) Point {
	return Point{
		X: Max(a.X, b.X),
		Y: Min(a.Y, b.Y),
	}
}

func mergeBottomRightPoints(a, b Point) Point {
	return Point{
		X: Max(a.X, b.X),
		Y: Max(a.Y, b.Y),
	}
}

func mergeBottomLeftPoints(a, b Point) Point {
	return Point{
		X: Min(a.X, b.X),
		Y: Max(a.Y, b.Y),
	}
}
