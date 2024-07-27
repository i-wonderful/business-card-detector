package field_sort

import (
	"card_detector/internal/model"
	"card_detector/internal/service/field_sort/helper"
)

// categorizeByIcons - распределение полей на основании найденных значков.
// @return []model.DetectWorld - не распределенные данные
func (s *Service) categorizeByIcons(recognized map[string]interface{}, data []model.DetectWorld, boxes []model.TextArea) []model.DetectWorld {

	var skype string
	var telegram []string

	if _, ok := recognized[FIELD_SKYPE]; ok {
		skype = recognized[FIELD_SKYPE].(string)
	}

	if _, ok := recognized[FIELD_TELEGRAM]; ok {
		telegram = recognized[FIELD_TELEGRAM].([]string)
	}

	if skype == "" {
		if isOk, box := isContainsLabel(boxes, "skype"); isOk {
			worldSkype, index := helper.FindNearestWorld(data, box)
			if index != -1 {
				skype = worldSkype.Text
				data = remove(data, index)
			}
		}
	}

	if len(telegram) == 0 {
		if isOk, box := isContainsLabel(boxes, "telegram"); isOk {
			worldTelegram, index := helper.FindNearestWorld(data, box)
			if index != -1 {
				telegram = append(telegram, worldTelegram.Text)
				data = remove(data, index)
			}
		}
	}

	recognized[FIELD_SKYPE] = skype
	recognized[FIELD_TELEGRAM] = telegram

	return data
}

func isContainsLabel(boxes []model.TextArea, label string) (bool, *model.TextArea) {
	for _, box := range boxes {
		if box.Label == label {
			return true, &box
		}
	}
	return false, nil
}
