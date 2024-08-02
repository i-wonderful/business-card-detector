package helper

import (
	. "card_detector/internal/model"
	manage_str "card_detector/internal/util/str"
	"math"
)

func GetOnlyWorlds(detectWorlds []DetectWorld) []string {
	worlds := make([]string, len(detectWorlds))
	for i, world := range detectWorlds {
		worlds[i] = world.Text
	}
	return worlds
}

func FindNearestWorldToBox(data []DetectWorld, box *TextArea) (DetectWorld, int) {
	var nearest DetectWorld
	var index = -1
	minDistance := math.MaxFloat64 // Инициализируем минимальное расстояние максимально возможным значением

	// Определяем диапазон по оси Y с учетом небольшого увеличения высоты
	yTolerance := 1.2                                              // коэффициент для увеличения высоты на 20%
	yRangeStart := box.Y - int(float64(box.Height)*(yTolerance-1)) // Учитываем увеличение высоты снизу
	yRangeEnd := box.Y + int(float64(box.Height)*yTolerance)       // Учитываем увеличение высоты сверху

	for i, world := range data {
		// Проверяем, находится ли объект примерно на одной линии по оси Y
		if world.Box.PTop1.Y >= yRangeStart && world.Box.PBot1.Y <= yRangeEnd {
			// Вычисляем расстояние до объекта
			distance := math.Abs(float64(box.X - world.Box.PTop1.X))
			if distance < minDistance {
				minDistance = distance
				nearest = world
				index = i
			}
		}
	}

	return nearest, index
}

const maxVerticalDistanceThresholdPercent = 0.3
const maxHorizontalDistanceThresholdPercent = 0.6

// FindNearestByY finds the nearest by Y DetectWorld to the given item within a slice of DetectWorlds.
//
// item: a pointer to the DetectWorld item to find the nearest world to.
// worlds: a slice of DetectWorlds to search for the nearest one.
// Returns a pointer to the nearest DetectWorld and a boolean indicating if the nearest world is above or not.
func FindNearestByY(item *DetectWorld, worlds []DetectWorld) *DetectWorld {
	var nearest DetectWorld
	minDistance := math.MaxFloat64
	itemBottom := item.Box.PBot1.Y
	itemTop := item.Box.PTop1.Y
	itemLeft := item.Box.PTop1.X
	itemRight := item.Box.PTop2.X
	isFind := false

	for _, world := range worlds {
		if !manage_str.IsOnlyLettersExtend(world.Text) {
			continue
		}

		worldTop := world.Box.PTop1.Y
		worldBottom := world.Box.PBot1.Y
		worldLeft := world.Box.PTop1.X
		worldRight := world.Box.PTop2.X

		maxVerticalDistancePx := math.Max(float64(item.Box.H), float64(world.Box.H)) * maxVerticalDistanceThresholdPercent
		maxHorizontalDistancePx := math.Max(float64(item.Box.W), float64(world.Box.W)) * maxHorizontalDistanceThresholdPercent

		// Проверка горизонтального расстояния
		horizontalDistance := math.Min(math.Abs(float64(itemLeft-worldLeft)), math.Abs(float64(worldRight-itemRight)))
		if horizontalDistance > maxHorizontalDistancePx {
			continue
		}

		distanceToTop := math.Abs(float64(itemBottom - worldTop))
		distanceToBottom := math.Abs(float64(worldBottom - itemTop))

		if (distanceToTop <= maxVerticalDistancePx || distanceToBottom <= maxVerticalDistancePx) &&
			(distanceToTop < minDistance || distanceToBottom < minDistance) {
			nearest = world
			if distanceToTop < distanceToBottom {
				minDistance = distanceToTop
				isFind = true
			} else {
				minDistance = distanceToBottom
				isFind = true
			}
		}
	}

	if !isFind {
		return nil
	}

	return &nearest
}

func IsAbove(item1, item2 DetectWorld) bool {
	return item1.Box.PTop1.Y < item2.Box.PBot2.Y
}
