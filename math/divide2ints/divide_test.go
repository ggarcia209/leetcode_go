package divide2ints

import (
	"testing"
)

func TestDivide(t *testing.T) {
	var tests = []struct {
		dividend int
		divisor  int
		quotient int
	}{
		{100, 10, 10},
		{100, 25, 4},
		{75, 10, 7},
		{100, -10, -10},
		{0, 1, 0},
		{2, 1, 2},
		{12, 14, 0},
		{-1, 1, -1},
		{-2, -1, 2},
		{0, -1, 0},
		{1, 1, 1},
		{-2147483648, -1, 2147483647},
	}
	for _, test := range tests {
		quotient := divide(test.dividend, test.divisor)
		if quotient != test.quotient {
			t.Errorf("divide failed - result: %d; want: %d", quotient, test.quotient)
		}
	}
}
