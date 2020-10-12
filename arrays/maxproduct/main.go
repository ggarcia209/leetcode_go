package main

import (
	"fmt"
	"time"
)

// 1,2,3,-3,-5,6,7,0,9,2,-3,2,1,-2

func main() {
	nums := []int{2, -5, -2, -4, 3}
	start := time.Now()
	max := maxProduct(nums)
	fmt.Println(time.Since(start))
	fmt.Println(max)

}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0] // first number encountered is always the the max at start
	sub := max     // max product of substring
	if len(nums) == 1 {
		return max
	}

	var j int
	var mark bool
	for i, n := range nums {
		if n <= 0 && !mark {
			mark = true
			j = i + 1 // index of first value in substring passed to recusive call
		}
		if i == 0 {
			continue // don't multiply max by self
		}
		sub *= n
		if sub > max {
			max = sub
		}
		if sub == 0 {
			break
		}

	}

	if sub <= 0 {
		// find product of substring beginning after last negative number
		// Only one recursive call at most, substring product always positive
		sub2 := maxProduct(nums[j:])
		if sub2 > max {
			max = sub2
		}
	}

	return max
}
