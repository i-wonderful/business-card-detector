package field_sort

import (
	. "card_detector/internal/model"
	. "card_detector/internal/service/field_sort/helper"
	. "card_detector/internal/util/str"
	"regexp"
	"strings"
	"unicode"
)

// processNameByKnownNames checks if the given line contains any known names and modifies the name string accordingly.
//
// line: the input string to check for known names.
// name: a pointer to a string to be modified if a known name is found.
// bool: true if a known name is found and processed, false otherwise.
func (s *Service) processNameByKnownNames(line string, name *string) bool {
	isFind, nameFind := IsContainsWith(line, s.names)
	if isFind {
		*name = InsertSpaceIfNeeded(line, nameFind)
		return true
	}
	return false
}

// processNameByMailName - определить имя на основании почты. В почте могут встречаться часть имени.
func (s *Service) processNameByMailName(line string, mailName string, name *string) bool {
	if nameIsFull(*name) || mailName == "" {
		return false
	}

	if len(mailName) < 3 {
		return s.processNameByInitials(line, mailName, name)
	}

	re := regexp.MustCompile(`[._]`)
	mailNames := re.Split(mailName, -1)

	for _, m := range mailNames {
		if len(m) < 3 {
			continue
		}
		if ContainsIgnoreCase(line, m) || ContainsIgnoreCase(m, line) {
			*name += " " + line
			*name = strings.Trim(*name, " ")
			return true
		}
	}
	return false
}

// processNameByInitials - определить имя на основании инициалов встреченных в email.
func (s *Service) processNameByInitials(line string, mailName string, name *string) bool {
	if len(mailName) > 3 {
		return false
	}
	initials := getInitials(line)
	if initials == "" {
		return false
	}

	if ContainsIgnoreCase(mailName, initials) || ContainsIgnoreCase(initials, mailName) {
		*name = line
		return true
	}
	return false
}

// processSurnameIfSingleName updates the name if it is a single name by adding the nearest world text.
//
// Parameters: name *string, item *DetectWorld, worlds []DetectWorld
// Returns nothing.
func (s *Service) processSurnameIfSingleName(name *string, nameBox *DetectWorld, worldBoxes []DetectWorld) {
	if len(worldBoxes) == 0 || *name == "" || nameIsFull(*name) {
		return
	}
	nearest := FindNearestByY(nameBox, worldBoxes)
	if nearest == nil {
		return
	}
	isAbove := IsAbove(*nameBox, *nearest)
	if isAbove {
		*name += " " + nearest.Text
		nameBox.Box.PBot1 = nearest.Box.PBot1
		nameBox.Box.PBot2 = nearest.Box.PBot2
	} else {
		*name = nearest.Text + " " + *name
		nameBox.Box.PTop1 = nearest.Box.PTop1
		nameBox.Box.PTop2 = nearest.Box.PTop2
	}
}

func nameIsFull(val string) bool {
	if val == "" {
		return false
	}
	names := strings.Split(val, " ")
	return len(names) >= 2
}

// getInitials - получить инициалы
func getInitials(name string) string {
	words := strings.Fields(name)
	if len(words) == 1 {
		return ""
	}

	initials := []rune{}
	for _, word := range words {
		if len(word) > 0 && unicode.IsLetter(rune(word[0])) {
			initials = append(initials, rune(word[0]))
		}
	}
	if len(initials) == 0 {
		return ""
	}
	return string(initials)
}
