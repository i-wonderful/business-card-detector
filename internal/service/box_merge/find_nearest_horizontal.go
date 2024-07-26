package box_merge

import (
	. "card_detector/internal/model"
	. "card_detector/internal/util/calc"
	"math"
)

// Поиск ближайшего бокса с другим текстом
func findNearestHorizontal(item DetectWorld, worlds []DetectWorld) (*DetectWorld, int) {
	var nearest *DetectWorld
	index := -1
	minDistance := math.MaxInt32

	for i := range worlds {
		world := worlds[i] // Avoid taking the address of an iterator variable
		if item.Text == world.Text {
			continue
		}
		if isOnSameLevel(item.Box, world.Box) {
			dist := horizontalDistance(item.Box, world.Box)
			if dist < minDistance {
				minDistance = dist
				nearest = &world
				index = i
			}
		}
	}

	return nearest, index
}

// Проверяет, находятся ли два бокса на одном уровне (перекрытие по вертикали больше 70%)
func isOnSameLevel(box1, box2 Rectangle) bool {
	top1, bot1 := Min(box1.PTop1.Y, box1.PTop2.Y), Max(box1.PBot1.Y, box1.PBot2.Y)
	top2, bot2 := Min(box2.PTop1.Y, box2.PTop2.Y), Max(box2.PBot1.Y, box2.PBot2.Y)

	overlap := Max(0, Min(bot1, bot2)-Max(top1, top2))
	minHeight := Min(bot1-top1, bot2-top2)

	return float64(overlap)/float64(minHeight) > 0.8
}

// Вычисляет горизонтальное расстояние между двумя боксами
func horizontalDistance(box1, box2 Rectangle) int {
	center1 := (box1.PTop1.X + box1.PTop2.X + box1.PBot1.X + box1.PBot2.X) / 4
	center2 := (box2.PTop1.X + box2.PTop2.X + box2.PBot1.X + box2.PBot2.X) / 4

	return int(math.Abs(float64(center1 - center2)))
}
