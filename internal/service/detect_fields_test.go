package service

import "testing"

func TestDetectEmail(t *testing.T) {
	//testCases := []struct {
	//	input    []string
	//	expected map[string]string
	//}{
	//	{
	//		input:    []string{"johndoe@email.com"},
	//		expected: map[string]string{"email": "johndoe@email.com"},
	//	},
	//	{
	//		input:    []string{": russell@connectingbrands.co.uk"},
	//		expected: map[string]string{"email": "russell@connectingbrands.co.uk"},
	//	}, {
	//		input:    []string{"kate@oddsdigger.com", "elterova.ekaterina", "@elterova"},
	//		expected: map[string]string{"email": "kate@oddsdigger.com"},
	//	},
	//}
	// todo

	//detector := NewDetector("../config/professions.txt",
	//	"../config/company.txt")
	//for _, tc := range testCases {
	//	fields := detector.DetectFields(tc.input)
	//	for k, v := range tc.expected {
	//		if fields[k] != v {
	//			t.Errorf("Expected %s:%s, got %s:%s", k, v, k, fields[k])
	//		}
	//	}
	//}
}

func TestDetectSkype(t *testing.T) {
	//testCases := []struct {
	//	input    []string
	//	expected map[string]string
	//}{
	//	{
	//		input:    []string{"elterova.ekaterina"},
	//		expected: map[string]string{"skype": "elterova.ekaterina"},
	//	},
	//}

	// todo
	//detector := NewDetector("../config/professions.txt", "../config/company.txt")
	//for _, tc := range testCases {
	//	fields := detector.DetectFields(tc.input)
	//	for k, v := range tc.expected {
	//		if fields[k] != v {
	//			t.Errorf("Expected %s:%s, got %s:%s", k, v, k, fields[k])
	//		}
	//	}
	//}
}

//func TestDetectJobTitle(t *testing.T) {
//	testCases := []struct {
//		input    []string
//		expected map[string]string
//	}{
//		{
//			input:    []string{"AFFILIATE MANAGER"},
//			expected: map[string]string{"jobTitle": "AFFILIATE MANAGER"},
//		},
//	}
//
//	detector := &Detector{
//		professions: readFile("professions.txt"),
//	}
//	for _, tc := range testCases {
//		fields := detector.DetectFields(tc.input)
//		for k, v := range tc.expected {
//			if fields[k] != v {
//				t.Errorf("Expected %s:%s, got %s:%s", k, v, k, fields[k])
//			}
//		}
//	}
//}

//func TestDetectFields(t *testing.T) {
//

//
//	// Test name extraction
//	data = []string{"John Doe"}
//	fields = DetectFields(data)
//	if fields["name"] != "John Doe" {
//		t.Errorf("Expected name John Doe, got %s", fields["name"])
//	}
//
//	// Test phone number
//	data = []string{"+1 234 567 8910"}
//	fields = DetectFields(data)
//	if fields["phone"] != "+1 234 567 8910" {
//		t.Errorf("Expected phone +1 234 567 8910, got %s", fields["phone"])
//	}
//
//	// Test website
//	data = []string{"www.website.com"}
//	fields = DetectFields(data)
//	if fields["site"] != "www.website.com" {
//		t.Errorf("Expected site www.website.com, got %s", fields["site"])
//	}
//
//	// Test company name
//	data = []string{"ACME Ltd"}
//	fields = DetectFields(data)
//	if fields["company"] != "ACME Ltd" {
//		t.Errorf("Expected company ACME Ltd, got %s", fields["company"])
//	}
//
//	// Test Skype
//	data = []string{"johndoe"}
//	fields = DetectFields(data)
//	if fields["skype"] != "johndoe" {
//		t.Errorf("Expected skype johndoe, got %s", fields["skype"])
//	}
//
//	// Test other field
//	data = []string{"Software Developer"}
//	fields = DetectFields(data)
//	if fields["other"] != "Software Developer" {
//		t.Errorf("Expected other Software Developer, got %s", fields["other"])
//	}
//
//}
