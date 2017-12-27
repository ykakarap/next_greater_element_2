package main

import (
	"fmt"
)

func main() {
	test := []int{1, 2, 1}
	fmt.Printf("Answer is %v", nextGreaterElements(test))

}

func nextGreaterElements(nums []int) []int {
	ans := make([]int, len(nums))
	// Initialize everything with -1
	for kk := range ans {
		ans[kk] = -1
	}

	// Cyclic loop through the array
	for k, v := range nums {
		searched := 0
		for i := k; searched < len(nums); i = next(i, len(nums)) {
			searched = searched + 1
			if nums[i] > v {
				ans[k] = nums[i]
				// Found the next greatest value. Terminate now
				break
			}
		}
	}

	return ans
}

func next(idx, len int) int {
	if idx < len-1 {
		return idx + 1
	} else {
		return 0
	}
}
