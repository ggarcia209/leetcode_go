package removeElement

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var tests = []struct {
		input  []int
		val    int
		want   []int
		length int
	}{
		{[]int{3, 2, 2, 3}, 3, []int{2, 2}, 2},
		{[]int{1, 2, 3, 4}, 3, []int{1, 2, 4}, 3},
		{[]int{3}, 3, []int{}, 0},
		{[]int{3}, 2, []int{3}, 1},
		{[]int{}, 2, []int{}, 0},
		{[]int{1, 2, 3}, 3, []int{1, 2}, 2},
		{[]int{3, 2, 1}, 3, []int{2, 1}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}, 4, []int{1, 2, 3, 5, 6, 5, 3, 2, 1}, 9},
	}
	for _, test := range tests {
		input := test.input
		length := removeElement(input, test.val)
		for i, n := range input {
			if i == len(test.want) {
				break
			}
			if n != test.want[i] {
				t.Errorf("removeElement failed (array) - result: %v; want: %v", input, test.want)
			}
		}
		if length != test.length {
			t.Errorf("removeElement failed (length) - result: %d; want: %d", length, test.length)
		}
	}
}
