package main

import "fmt"

// 1-2 negative ints +
// 1-2 positive ints
//    OR
// 1 negative +
// 0 +
// 1 positive

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sums := threeSum(nums)
	fmt.Println(sums)
}

func threeSum(nums []int) [][]int {
	var results [][]int
	var pos []int
	var neg []int
	ct := make(map[int]int)

	// create slices with each unique negative int and positive int (includes 0)
	// map each int to it's count in list
	for _, n := range nums {
		if n < 0 && ct[n] == 0 {
			neg = append(neg, n)
		} else if n > 0 && ct[n] == 0 {
			pos = append(pos, n)
		}
		ct[n]++
	}

	if ct[0] > 2 {
		results = append(results, []int{0, 0, 0})
	}
	// Add 1 negative (n) and 1 positive (p), check if -(n+p) in map
	// (n + p) + -(n + p) = 0
	for _, n := range neg {
		roundCt := ct
		roundCt[n]--
		for _, p := range pos {
			roundCt[p]--
			m := (n + p) * -1
			if roundCt[m] > 0 { // valid - in list
				if (m == n || m == p) && roundCt[m] > 1 { // multiple ints of same value -  ok
					results = append(results, []int{n, m, p})
					continue
				} else if m != n && m != p && roundCt[m] > 0 { // unique ints
					results = append(results, []int{n, m, p})
					continue
				} else {
					continue
				}

			}
		}
		ct[n] = 0
	}
	return results
}
