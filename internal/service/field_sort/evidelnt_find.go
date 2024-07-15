package field_sort

const FIELD_PHONE = "phone"

// evident fields sort algorithm - нахождение очевидных полей
// @todo: доработать алгоритм
// @return map[string]interface{}, []string - recognized fields and not recognized fields
func (s *Service) evidentSort(data []string) (map[string]interface{}, []string) {
	notDetectItems := []string{}
	phones := []string{}

	for _, line := range data {
		line = clearTrashSymbols(line)

		if phoneRegex.MatchString(line) && notContainsLetters(line) {
			p := phoneRegex.FindString(line)
			phones = append(phones, p)
		} else {
			notDetectItems = append(notDetectItems, line)
		}
	}

	recognized := map[string]interface{}{
		FIELD_PHONE: clearPhones(phones),
	}
	return recognized, notDetectItems
}
