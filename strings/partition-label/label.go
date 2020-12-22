package label

import (
	"strings"
)

/* string S of lowercase English letters is given. We want to partition this string
into as many parts as possible so that each letter appears in at most one part,
and return a list of integers representing the size of these parts. */

// success - 4ms/3.1MB
func partitionLabels(S string) []int {
	m := make(map[string][]int)
	prt := [][]int{}
	lens := []int{}
	j := 0
	ss := strings.Split(S, "")
	for i, s := range ss {
		// find first, last index of each letter
		if len(m[s]) == 0 {
			m[s] = []int{i, i}
			continue
		}
		m[s][1] = i
	}
	for _, k := range ss {
		if len(prt) == 0 { // create first partition by first/last indices
			prt = append(prt, []int{m[k][0], m[k][1]})
			continue
		}
		if m[k][0] > prt[j][1] { // new partition
			prt = append(prt, []int{m[k][0], m[k][1]})
			j++
		}
		if m[k][1] > prt[j][1] && m[k][0] < prt[j][1] { // update parition end index
			prt[j][1] = m[k][1] // new partition end index
		}
	}
	for _, p := range prt { // derive lengths of partitions from first/last indices
		lens = append(lens, (p[1] - p[0] + 1))
	}
	return lens
}

// success - 0ms/2.2MB
func partitionLabelsV2(S string) []int {
	seen := make(map[rune]bool)
	counts := []int{}
	n := 0 // number of distinct characters in substring
	c := 1 // number of characters in substring

	for i, r := range S {
		if !seen[r] {
			seen[r] = true
			n++
		}
		if seen[r] {
			// decrement count of distinct chars when 0 remaining instances in string
			if i < len(S)-1 && strings.Index(S[i+1:], string(r)) == -1 {
				n--
			}
			if i == len(S)-1 {
				n--
			}
			// add count when 0 remaining instances of distinct chars in
			// remaining substring; reset count
			if n == 0 {
				counts = append(counts, c)
				c = 0
			}
		}
		c++
	}

	return counts
}
