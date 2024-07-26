package box_merge

import (
	. "card_detector/internal/model"
	. "card_detector/internal/util/calc"
	manage_str "card_detector/internal/util/str"
	"math"
	"sort"
	"unicode"
)

const closerHPercent = 0.25 // если высота боксов не различается больше чем на 25%, то считаем боксы близкими
const percentCloserX = 0.15 // если расстояние между боксами не различается больше чем на 15%, то считаем боксы близкими
var keyWorlds = []string{"Phone", "tel", "Email", "mail", "Skype", "Website", "site", "web", "telegram"}

// MergeBoxesVertical - merge boxes with close values
func MergeBoxesVertical(boxes []DetectWorld) []DetectWorld {
	sortByHeight(boxes)

	rez := []DetectWorld{}
	i := 0
	isNeedMergeKeyWorlds := false

	for i < len(boxes) {
		prev := &boxes[i]

		if manage_str.IsContains(prev.Text, keyWorlds) {
			rez = append(rez, *prev)
			i++
			isNeedMergeKeyWorlds = true
			continue
		}

		if i+1 < len(boxes) {
			next := &boxes[i+1]
			isCloser25Percent := isCloserH(prev.Box.H, next.Box.H)
			hMax := Max(prev.Box.H, next.Box.H)
			isCloser2Percent := isCloserX(prev.Box.PTop1.X, next.Box.PTop1.X, hMax)
			if isCloser25Percent && isCloser2Percent &&
				isOnlyLetters(prev.Text) && isOnlyLetters(next.Text) {
				if i+2 < len(boxes) {
					next2 := &boxes[i+2]
					isCloser25Percent2 := isCloserH(next.Box.H, next2.Box.H)
					isCloser2Percent2 := isCloserX(next.Box.PTop1.X, next2.Box.PTop1.X, hMax)
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
				rez = append(rez, *mergeTwoWorlds(prev, next))
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

	if isNeedMergeKeyWorlds {
		rez = MergeKeyWorldsHorizontal(rez)
	}

	return rez
}

func MergeKeyWorldsHorizontal(boxes []DetectWorld) []DetectWorld {
	keyWorldsBoxes := []DetectWorld{}
	rez := []DetectWorld{}
	for _, w := range boxes {
		if len(w.Text) < 6 && manage_str.IsContains(w.Text, keyWorlds) {
			keyWorldsBoxes = append(keyWorldsBoxes, w)
		} else {
			rez = append(rez, w)
		}
	}

	if len(keyWorldsBoxes) == 0 {
		return rez
	}

	sortByY(rez)

	for _, kwBox := range keyWorldsBoxes {
		_, index := findNearestHorizontal(kwBox, rez)
		if index >= 0 {
			rez[index].Text = manage_str.ClearTrashSymbols(kwBox.Text) + " " + manage_str.ClearTrashSymbols(rez[index].Text)
			rez[index].Box.PTop1 = kwBox.Box.PTop1
			rez[index].Box.PBot2 = kwBox.Box.PBot2
		}
	}

	return rez
}

func deleteByIndex(worlds []DetectWorld, index int) []DetectWorld {
	return append(worlds[:index], worlds[index+1:]...)
}

func sortByHeight(worlds []DetectWorld) []DetectWorld {
	sort.Slice(worlds, func(i, j int) bool {
		w1 := worlds[i]
		w2 := worlds[j]
		if isCloserH(w1.Box.H, w2.Box.H) {
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

func sortByY(worlds []DetectWorld) []DetectWorld {
	sort.Slice(worlds, func(i, j int) bool {
		w1 := worlds[i]
		w2 := worlds[j]
		//if math.Abs(float64(w1.Prob-w2.Prob)) > 0.04 {
		//	return w1.Prob > w2.Prob
		//}
		return w1.Box.PTop1.Y < w2.Box.PTop1.Y
	})
	return worlds
}

// Числа отличаются не более чем на 25%
func isCloserH(a, b int) bool {
	diff := math.Abs(float64(a - b))
	limitPercent := closerHPercent * float64(Max(a, b))
	return diff <= limitPercent
}

func isCloserX(a, b, h int) bool {
	diff := math.Abs(float64(a - b))
	limitPercent := percentCloserX * float64(h)
	return diff <= limitPercent
}

func intContains(slice []int, value int) bool {
	if len(slice) == 0 {
		return false
	}
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
