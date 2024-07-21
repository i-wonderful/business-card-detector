package helper

import "testing"

func TestExtractSkypeSkype(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Skype:sidagarwal17", "sidagarwal17"},
		{"s:sidagarwal17", "sidagarwal17"},
		{"Skype:sidagarwal17@example.com", "sidagarwal17@example.com"},
		{"Office Address:Unit-No.C-617,1-ThumSector-62,Noida", ""},
		{"Skype:alex.softgamings.com", "alex.softgamings.com"},
		{"S:russ.yershon", "russ.yershon"},
		{"skype: sidagarwal17", "sidagarwal17"},
		{"Skype flavio.tamega", "flavio.tamega"},
		{"Mail: b2b@lLinebet.com Skype: partners@Linebet.com", "partners@Linebet.com"},
		{"My: skype:id1 and skype:id2", "id1"},
		{"My Skype: SAMPLE_ID", "SAMPLE_ID"},
		{"My Skype: skype.sample_id", "skype.sample_id"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := ExtractSkypeSkype(test.input)
			if result != test.expected {
				t.Errorf("Expected %q, but got %q", test.expected, result)
			}
		})
	}
}
