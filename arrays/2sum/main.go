package main

import (
	"fmt"
	"time"
)

// ex: l = [2, 4, 7 11, 19]; t = 9;
// a = [0,1] (l[0] + l[1] = t)

func main() {
	nums := []int{11, 7, 4, 6, 2, 19}
	start := time.Now()
	sum := twoSum(nums, 6)
	fmt.Println(time.Since(start))
	fmt.Println(sum)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] > target {
				continue
			}
			t := nums[i] + nums[j]
			if t == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
