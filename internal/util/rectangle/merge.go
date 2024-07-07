package rectangle

import (
	. "card_detector/internal/model"
	"sort"
)

// Какой процент общей площади должен быть у прямоугольников,
// чтобы их можно было объединить (в данном случае 80%).
const overlapThreshold = 0.73
const maxYDiffPercentage = 0.15 // Максимальная разница по Y в процентах от средней высоты
const maxYDiff = 10             // Максимальная разница по Y для объединения прямоугольников
const maxXGap = 10              // Максимальный промежуток по X для объединения прямоугольников по горизонтали

func MergeOverlappingRectangles(rectangles []TextArea) []TextArea {
	if len(rectangles) == 0 {
		return rectangles
	}

	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].Y < rectangles[j].Y
	})

	var mergedRectangles []TextArea
	currentRect := rectangles[0]

	for i := 1; i < len(rectangles); i++ {
		rect := rectangles[i]

		currentArea := currentRect.Width * currentRect.Height
		rectArea := rect.Width * rect.Height
		overlapArea := overlapArea(currentRect, rect)

		if float64(overlapArea) < overlapThreshold*float64(min(currentArea, rectArea)) {
			// Проверка объединения по горизонтали
			if isHorizontallyAdjacent(currentRect, rect) {
				// Объединить прямоугольники по горизонтали
				minX := min(currentRect.X, rect.X)
				maxX := max(currentRect.X+currentRect.Width, rect.X+rect.Width)
				maxY := max(currentRect.Y+currentRect.Height, rect.Y+rect.Height)
				currentRect = TextArea{
					X:      minX,
					Y:      min(currentRect.Y, rect.Y),
					Width:  maxX - minX,
					Height: maxY - min(currentRect.Y, rect.Y),
				}
			} else {
				mergedRectangles = append(mergedRectangles, currentRect)
				currentRect = rect
			}
		} else {
			minX := min(currentRect.X, rect.X)
			maxX := max(currentRect.X+currentRect.Width, rect.X+rect.Width)
			maxY := max(currentRect.Y+currentRect.Height, rect.Y+rect.Height)
			currentRect = TextArea{
				X:      minX,
				Y:      min(currentRect.Y, rect.Y),
				Width:  maxX - minX,
				Height: maxY - min(currentRect.Y, rect.Y),
			}
		}
	}

	mergedRectangles = append(mergedRectangles, currentRect)

	return mergedRectangles
}

// overlapArea возвращает площадь пересечения двух прямоугольников
func overlapArea(r1, r2 TextArea) int {
	x1 := max(r1.X, r2.X)
	y1 := max(r1.Y, r2.Y)
	x2 := min(r1.X+r1.Width, r2.X+r2.Width)
	y2 := min(r1.Y+r1.Height, r2.Y+r2.Height)

	if x2 < x1 || y2 < y1 {
		return 0
	}

	return (x2 - x1) * (y2 - y1)
}

//func isHorizontallyAdjacent(r1, r2 Box) bool {
//	return abs(r1.Y-r2.Y) <= maxYDiff && (r1.X+r1.Width+maxXGap >= r2.X || r2.X+r2.Width+maxXGap >= r1.X)
//}

func isHorizontallyAdjacent(r1, r2 TextArea) bool {
	avgHeight := (r1.Height + r2.Height) / 2
	maxYDiff := int(float64(avgHeight) * maxYDiffPercentage)
	return abs(r1.Y-r2.Y) <= maxYDiff && (r1.X+r1.Width+maxXGap >= r2.X || r2.X+r2.Width+maxXGap >= r1.X)
}

// @Deprecated

func MergeOverlappingRectangles2(rectangles []TextArea) []TextArea {
	// Если массив пустой, вернуть его
	if len(rectangles) == 0 {
		return rectangles
	}

	// Отсортировать массив по значению Y
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].Y < rectangles[j].Y
	})

	var mergedRectangles []TextArea
	currentRect := rectangles[0]

	for i := 1; i < len(rectangles); i++ {
		rect := rectangles[i]

		// Если текущий прямоугольник не пересекается с предыдущим,
		// добавить предыдущий прямоугольник в результат и обновить currentRect
		if rect.Y > currentRect.Y+currentRect.Height || rect.Y > currentRect.Y+maxYDiff {
			mergedRectangles = append(mergedRectangles, currentRect)
			currentRect = rect
		} else {
			// Иначе, объединить текущий прямоугольник с предыдущим
			minX := min(currentRect.X, rect.X)
			maxX := max(currentRect.X+currentRect.Width, rect.X+rect.Width)
			maxY := max(currentRect.Y+currentRect.Height, rect.Y+rect.Height)
			currentRect = TextArea{
				X:      minX,
				Y:      min(currentRect.Y, rect.Y),
				Width:  maxX - minX,
				Height: maxY - min(currentRect.Y, rect.Y),
			}
		}
	}

	// Добавить последний прямоугольник в результат
	mergedRectangles = append(mergedRectangles, currentRect)

	return mergedRectangles
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
