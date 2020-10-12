package main

import "fmt"

var nums = []int{}

func main() {
	condup := containsDuplicate(nums)
	fmt.Printf("contains duplicate: %v", condup)
}

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, n := range nums {
		if !seen[n] {
			seen[n] = true
			continue
		} else {
			return true
		}
	}
	return false
}
