package box_merge

import (
	"card_detector/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		name     string
		boxes    []model.DetectWorld
		expected model.DetectWorld
	}{
		//{
		//	name: "Two overlapping boxes",
		//	boxes: []model.DetectWorld{
		//		{
		//			Text: "Klientu vaditajs",
		//			Box: model.NewBoxFromPoints(
		//				model.NewPoint(108, 255),
		//				model.NewPoint(276, 255),
		//				model.NewPoint(276, 279),
		//				model.NewPoint(108, 279),
		//			),
		//			Prob: 0.9732199,
		//		},
		//		{
		//			Text: "Ekomercijas departaments",
		//			Box: model.NewBoxFromPoints(
		//				model.NewPoint(111, 276),
		//				model.NewPoint(397, 276),
		//				model.NewPoint(397, 300),
		//				model.NewPoint(111, 300),
		//			),
		//		},
		//	},
		//	expected: model.DetectWorld{
		//		Text: "Klientu vaditajs Ekomercijas departaments",
		//		Box: model.NewBoxFromPoints(
		//			model.NewPoint(108, 255),
		//			model.NewPoint(397, 255),
		//			model.NewPoint(397, 300),
		//			model.NewPoint(108, 300),
		//		),
		//	},
		//},
		{
			"Merge Vladislav Belov",
			[]model.DetectWorld{
				{
					Text: "Vladislav",
					Box: model.NewBoxFromPoints(
						model.NewPoint(89, 307),
						model.NewPoint(423, 307),
						model.NewPoint(423, 367),
						model.NewPoint(89, 367),
					),

					Prob: 0.98891234,
				},
				{
					Text: "Belov",
					Box: model.NewBoxFromPoints(
						model.NewPoint(82, 381),
						model.NewPoint(294, 383),
						model.NewPoint(293, 447),
						model.NewPoint(81, 444),
					),
				},
			},
			model.DetectWorld{
				Text: "Vladislav Belov",
				Box: model.NewBoxFromPoints(
					model.NewPoint(82, 307),
					model.NewPoint(423, 307),
					model.NewPoint(423, 447),
					model.NewPoint(81, 444),
				),
			},
		},
		{
			"Merge Vladislav Belov2",
			[]model.DetectWorld{
				{
					Text: "Vladislav",
					Box: model.NewBoxFromPoints(
						model.NewPoint(97, 315),
						model.NewPoint(427, 315),
						model.NewPoint(427, 377),
						model.NewPoint(97, 377),
					),
					Prob: 0.98896325,
				},
				{
					Text: "Belov",
					Box: model.NewBoxFromPoints(
						model.NewPoint(89, 389),
						model.NewPoint(299, 391),
						model.NewPoint(298, 454),
						model.NewPoint(88, 452),
					),
					Prob: 0.9925569,
				},
			},
			model.DetectWorld{
				Text: "Vladislav Belov",
				Box: model.NewBoxFromPoints(
					model.NewPoint(89, 315),
					model.NewPoint(427, 315),
					model.NewPoint(427, 454),
					model.NewPoint(88, 452),
				),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			merged := MergeBoxesVertical(tc.boxes)
			if assert.Equal(t, 1, len(merged)) {
				assert.Equal(t, tc.expected.Text, merged[0].Text)
				assert.Equal(t, tc.expected.Box.PTop1, merged[0].Box.PTop1)
				assert.Equal(t, tc.expected.Box.PTop2, merged[0].Box.PTop2)
				assert.Equal(t, tc.expected.Box.PBot1, merged[0].Box.PBot1)
				assert.Equal(t, tc.expected.Box.PBot2, merged[0].Box.PBot2)
			}
		})
	}
}
