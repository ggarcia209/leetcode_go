package main

import (
	"fmt"
	"sort"
)

var nums1 = []int{1, 2}
var nums2 = []int{3, 4}

func main() {
	med := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("median: %v\n", med)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := append(nums1, nums2...)
	sort.IntSlice(total).Sort()

	median := 0.0

	if len(total)%2 == 0 {
		div := len(total) / 2
		lh, rh := total[:div], total[div:]
		median = (float64(lh[len(lh)-1]) + float64(rh[0])) / 2
	} else {
		div := len(total) / 2
		median = float64(total[div])
	}

	return median
}
