package helper

import (
	"testing"
)

func TestExtractBrokenUrl(t *testing.T) {
	tests := []struct {
		text    string
		domains []string
		zone    string
		want    string
	}{
		{"https://www.google.com", []string{"google"}, "com", "https://www.google.com"},
		{"www.nsano.comn", []string{"nsano"}, "com", "www.nsano.comn"},
		{"reypanda.com", []string{"revpanda"}, "com", "reypanda.com"},
	}

	for _, tt := range tests {
		testname := tt.text
		t.Run(testname, func(t *testing.T) {
			ans := ExtractBrokenUrl(tt.text, tt.domains, tt.zone)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
