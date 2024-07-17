package box_merge

import (
	"card_detector/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {

	boxes := []model.DetectWorld{
		{
			Text: "Klientu vaditajs",
			Box: model.NewBoxFromPoints(
				model.NewPoint(108, 255),
				model.NewPoint(276, 255),
				model.NewPoint(276, 279),
				model.NewPoint(108, 279),
			),
			Prob: 0.9732199,
		},
		{
			Text: "Ekomercijas departaments",
			Box: model.NewBoxFromPoints(
				model.NewPoint(111, 276),
				model.NewPoint(397, 276),
				model.NewPoint(397, 300),
				model.NewPoint(111, 300),
			),
		},
	}

	merged := MergeBoxes(boxes)

	assert.Equal(t, 1, len(merged))
	assert.Equal(t, "Klientu vaditajs Ekomercijas departaments", merged[0].Text)
	assert.Equal(t, model.NewPoint(108, 255), merged[0].Box.PTop1)
	assert.Equal(t, model.NewPoint(397, 255), merged[0].Box.PTop2)
	assert.Equal(t, model.NewPoint(108, 300), merged[0].Box.PBot2)
	assert.Equal(t, model.NewPoint(397, 300), merged[0].Box.PBot1)
}
