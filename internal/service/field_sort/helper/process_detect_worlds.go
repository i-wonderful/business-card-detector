package helper

import (
	"card_detector/internal/model"
	"math"
)

func GetOnlyWorlds(detectWorlds []model.DetectWorld) []string {
	worlds := make([]string, len(detectWorlds))
	for i, world := range detectWorlds {
		worlds[i] = world.Text
	}
	return worlds
}

func FindNearestWorld(data []model.DetectWorld, box *model.TextArea) (model.DetectWorld, int) {
	var nearest model.DetectWorld
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
