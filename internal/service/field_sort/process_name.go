package field_sort

import (
	. "card_detector/internal/model"
	. "card_detector/internal/service/field_sort/helper"
	. "card_detector/internal/util/str"
	"math"
	"regexp"
	"strings"
	"unicode"
)

// processNameByExistingNames - определить имя по существующим именам
func (s *Service) processNameByExistingNames(line string, name *string) bool {
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
func (s *Service) processSurnameIfSingleName(name *string, worldName *DetectWorld, worlds []DetectWorld) {
	if len(worlds) == 0 || *name == "" || nameIsFull(*name) {
		return
	}
	nearest, isAbove := findNearest(worldName, worlds)
	if nearest == nil {
		return
	}
	if isAbove {
		*name = nearest.Text + " " + *name
	} else {
		*name += " " + nearest.Text
	}
}

const distanceThreshold = 0.2

func findNearest(item *DetectWorld, worlds []DetectWorld) (*DetectWorld, bool) {
	var nearest *DetectWorld
	minDistance := math.MaxFloat64
	itemBottom := item.Box.PBot1.Y
	itemTop := item.Box.PTop1.Y
	isAbove := false

	for _, world := range worlds {
		if !IsOnlyLetters(world.Text) {
			continue
		}

		worldTop := world.Box.PTop1.Y
		worldBottom := world.Box.PBot1.Y
		maxDistance := math.Max(float64(item.Box.H), float64(world.Box.H)) * distanceThreshold

		distanceToTop := math.Abs(float64(itemBottom - worldTop))
		distanceToBottom := math.Abs(float64(worldBottom - itemTop))

		if (distanceToTop <= maxDistance || distanceToBottom <= maxDistance) && (distanceToTop < minDistance || distanceToBottom < minDistance) {
			nearest = &world
			if distanceToTop < distanceToBottom {
				minDistance = distanceToTop
				isAbove = false
			} else {
				minDistance = distanceToBottom
				isAbove = true
			}
		}
	}

	return nearest, isAbove
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
