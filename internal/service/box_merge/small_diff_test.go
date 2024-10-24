package box_merge

import "testing"

func TestWithinTenPercent(t *testing.T) {
	tests := []struct {
		a, b   int
		result bool
	}{
		{46, 59, true},
		//{89, 82, true},
		//{28, 24, false},
	}

	for _, test := range tests {
		if got := isCloserHeight(test.a, test.b); got != test.result {
			t.Errorf("withinTenPercent(%d, %d) = %v; want %v", test.a, test.b, got, test.result)
		}
	}
}

func TestCloserHeight(t *testing.T) {
	tests := []struct {
		a, b, h int
		result  bool
	}{
		{89, 87, 20, true},
	}

	for _, test := range tests {
		if got := isCloserX(test.a, test.b, test.h); got != test.result {
			t.Errorf("withinTenPercent(%d, %d) = %v; want %v", test.a, test.b, got, test.result)
		}
	}
}
