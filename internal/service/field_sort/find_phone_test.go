package field_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_FindPhone(t *testing.T) {

	testCases := []struct {
		name     string
		input    []string
		expected map[string]interface{}
	}{
		{
			"Phones",
			[]string{
				"+7 (473) 200-0-300 © ",
				"+7 (473) 20-20-457 (©",
				"773-00606",
			},
			map[string]interface{}{
				"phone": []string{
					"+7 (473) 200-0-300",
					"+7 (473) 20-20-457",
				},
			},
		},
		{
			"Phone and job",
			[]string{
				"Head of Sales DepartmentDirect:+371 25371708",
			},
			map[string]interface{}{
				"jobTitle": "Head of Sales Department",
				"phone":    []string{"+371 25371708"},
			},
		},
		{
			"With labels",
			[]string{
				"PT Office +351220 991 583",
				"MT Office +356 770 40806 ",
				"UK Mobile +44 7921239 788",
			},
			map[string]interface{}{
				"phone": []string{"+351220 991 583", "+356 770 40806", "+44 7921239 788"},
			},
		},
		{
			"With labels, without spaces",
			[]string{
				"M+254720961738",
				"T+254113804990",
			},
			map[string]interface{}{
				"phone": []string{"+254720961738", "+254113804990"},
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
			result := service.Sort(tc.input)
			for k, v := range tc.expected {
				assert.Equal(t, v, result[k])
			}
		})
	}
}
