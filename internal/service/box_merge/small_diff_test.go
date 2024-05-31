package box_merge

import "testing"

func TestWithinTenPercent(t *testing.T) {
	tests := []struct {
		a, b   int
		result bool
	}{
		{46, 59, true},
	}

	for _, test := range tests {
		if got := isCloser25Values(test.a, test.b); got != test.result {
			t.Errorf("withinTenPercent(%d, %d) = %v; want %v", test.a, test.b, got, test.result)
		}
	}
}
