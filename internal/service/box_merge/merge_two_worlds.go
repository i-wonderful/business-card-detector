package box_merge

import (
	. "card_detector/internal/model"
	. "card_detector/internal/util/calc"
	. "card_detector/internal/util/str"
)

func mergeTwoWorlds(item1, item2 *DetectWorld) *DetectWorld {
	return &DetectWorld{
		Text: ClearTrashSymbols(item1.Text) + " " + ClearTrashSymbols(item2.Text),
		Box: NewBoxFromPoints(
			mergeTopLeftPoints(item1.Box.PTop1, item2.Box.PTop1),
			mergeTopRightPoints(item1.Box.PTop2, item2.Box.PTop2),
			mergeBottomRightPoints(item1.Box.PBot1, item2.Box.PBot1),
			mergeBottomLeftPoints(item1.Box.PBot2, item2.Box.PBot2),
		),
		Prob: (item1.Prob + item2.Prob) / 2,
	}
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
