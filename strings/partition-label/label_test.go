package label

import (
	"testing"
)

var tests = []struct {
	input string
	want  []int
}{
	{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
	{"ababceecff", []int{4, 4, 2}},
	{"a", []int{1}},
	{"", []int{0}},
	{"aaa", []int{3}},
	{"abc", []int{1, 1, 1}},
	{"abca", []int{4}},
	{"abcad", []int{4, 1}},
}

func TestPartitionLabels(t *testing.T) {
	// success (4ms/3.1MB)
	for _, test := range tests {
		res := partitionLabels(test.input)
		for i, v := range res {
			if v != test.want[i] {
				t.Errorf("fail: %v; want: %v", res, test.want)
			}
		}
	}
}

func TestPartitionLabelsV2(t *testing.T) {
	for _, test := range tests {
		res := partitionLabelsV2(test.input)
		for i, v := range res {
			if v != test.want[i] {
				t.Errorf("fail: %v; want: %v", res, test.want)
			}
		}
	}
}
