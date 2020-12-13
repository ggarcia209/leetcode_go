package removeElement

/*
Given an array nums and a value val, remove all instances of that value
in-place and return the new length.

Do not allocate extra space for another array,
you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond
the new length.
*/

func removeElement(nums []int, val int) int {
	j := 0 // index offset
	for i := 0; i >= 0; i++ {
		if i-j > len(nums)-1 {
			break
		}
		if nums[i-j] == val {
			if i-j != len(nums)-1 {
				nums = append(nums[:i-j], nums[i+1-j:]...) // remove element
				j++                                        // offset index by 1 everytime an element is deleted
			} else {
				nums = nums[:i-j] // remove last element
			}
		}
	}
	return len(nums)
}
