package main

import "fmt"

func main() {
	nums1 := []int{5,7,7,8,10}
	indices := searchRange(nums1, 8)
	fmt.Println(indices)
}

func searchRange(nums []int, target int) []int {
	indices := []int{-1, -1}
	stop1, stop2 := false, false
	for i, num := range nums {
		if num == target && !stop1 {
			indices[0] = i
			stop1 = true
		}
		j := len(nums) - (1+i)
		opp := nums[j]
		if opp == target && !stop2 {
			indices[1] = j
			stop2 = true
		}
		if stop1 && stop2 {
			break
		}
	}
	return indices
}
