package field_sort

import (
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
func TestExtractSkype(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid Skype ID",
			input:    "My Skype: skype.sample_id",
			expected: "skype.sample_id",
		},
		//{
		//	name:     "Valid Skype ID with spaces",
		//	input:    "My Skype ID is: skype.live: sample_id",
		//	expected: "sample_id",
		//},
		{
			name:     "Valid Skype ID with case-insensitive pattern",
			input:    "My Skype: SAMPLE_ID",
			expected: "SAMPLE_ID",
		},
		{
			name:     "No Skype ID",
			input:    "This text does not contain",
			expected: "",
		},
		//{
		//	name:     "Multiple Skype IDs",
		//	input:    "My Skype IDs are: skype:id1 and skype:id2",
		//	expected: "id1",
		//},
		{
			name:     "Extract with @",
			input:    "Mail: b2b@lLinebet.com Skype: partners@Linebet.com",
			expected: "partners@Linebet.com",
		},
		{
			"Without colon",
			"Skype flavio.tamega",
			"flavio.tamega",
		},
		{
			"With two points",
			"Skype:alex.softgamings.com",
			"alex.softgamings.com",
		},
		//{
		//	"With point and colon",
		//	"live:.cid.e53090522ec2bf11",
		//	"live:.cid.e53090522ec2bf11",
		//},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := extractSkype(test.input)
			if result != test.expected {
				t.Errorf("Expected %q, but got %q", test.expected, result)
			}
		})
	}
}
