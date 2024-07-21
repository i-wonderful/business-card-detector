package field_sort

import (
	"card_detector/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_FindSkype(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected map[string]interface{}
	}{
		{
			"Skype",
			[]string{
				"Skype live.cid.9e53d8c1151b4b",
			},
			map[string]interface{}{
				"skype": "live:cid.9e53d8c1151b4b",
			},
		},
		{
			"Skype2",
			[]string{
				"s:live:cid.639e35052e7e9fe1 bla",
			},
			map[string]interface{}{
				"skype": "live:cid.639e35052e7e9fe1",
			},
		},
		{
			"Skype with two points",
			[]string{
				"Skype:alex.softgamings.com",
			},
			map[string]interface{}{
				"skype": "alex.softgamings.com",
			},
		},
		{
			"Skype with point and semicolon",
			[]string{"live:.cid.e53090522ec2bf11"},
			map[string]interface{}{
				"skype": "live:cid.e53090522ec2bf11",
			},
		},
		{
			"Skype without semicolon",
			[]string{"live.cid.639e35052e7e9fe1"},
			map[string]interface{}{
				"skype": "live:cid.639e35052e7e9fe1",
			},
		}, {
			"Skype with concat livecid",
			[]string{"livecid.e53090522ec2bf11"},
			map[string]interface{}{
				"skype": "live:cid.e53090522ec2bf11",
			},
		},
		{
			"Skype with S:",
			[]string{"S:russ.yershon"},
			map[string]interface{}{
				"skype": "russ.yershon",
			},
		},
		{
			"Without cid",
			[]string{"live:rajyousp"},
			map[string]interface{}{
				"skype": "live:rajyousp",
			},
		},
		{
			"Mix with other",
			[]string{
				"Office Address:Unit-No.C-617,I-Thum,Sector-62,Noida",
				"Skype: sidagarwal17",
			},
			map[string]interface{}{
				"skype": "sidagarwal17",
			},
		},
	}

	service := NewService(
		"../../../config/professions.txt",
		"../../../config/companies.txt",
		"../../../config/names.txt",
		true)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := service.Sort(wrapForSort(tc.input), nil)
			for k, v := range tc.expected {
				assert.Equal(t, v, result[k])
			}
		})
	}
}

func wrapForSort(inputs []string) []model.DetectWorld {
	result := make([]model.DetectWorld, len(inputs))
	for i, v := range inputs {
		result[i] = model.DetectWorld{Text: v}
	}
	return result
}
