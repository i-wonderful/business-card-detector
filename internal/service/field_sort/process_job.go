package field_sort

import (
	"card_detector/internal/model"
	"card_detector/internal/service/field_sort/helper"
)

func (s *Service) processJobByKnownJobs(line string, jobTitle *string) bool {
	isFind, namesFind := isContainsManyWith(line, s.professions)
	if isFind {
		*jobTitle += " " + line
		for _, n := range namesFind { // вставляем пробелы, если в должности они пропущены
			*jobTitle = helper.InsertSpaceIfNeeded(*jobTitle, n)
		}
		return true
	}
	return false
}

// processJobByNearestName processes a job by finding the nearest box to the name, in the given world boxes.
//
// Parameters:
//
//	job: a pointer to a string representing the job to be processed.
//	nameBox: a pointer to a DetectWorld struct representing the name box.
//	worldBoxes: a slice of DetectWorld structs representing the world boxes to search.
func (s *Service) processJobByNearestName(job *string, nameBox *model.DetectWorld, worldBoxes []model.DetectWorld) {
	if job != nil && len(*job) != 0 {
		return
	}

	if nameBox == nil || len(worldBoxes) == 0 {
		return
	}

	nearest := helper.FindNearestByY(nameBox, worldBoxes)
	if nearest == nil {
		return
	}

	nearest2 := helper.FindNearestByY(nearest, worldBoxes)
	*job = nearest.Text
	if nearest2 != nil {
		*job += " " + nearest2.Text
	}
}
