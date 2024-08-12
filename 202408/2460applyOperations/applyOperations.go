package main

import (
	"fmt"
)

func main() {
	fmt.Println(applyOperations([]int{1, 2, 2, 1, 1, 0}))
}

func applyOperations(nums []int) []int {
	n := len(nums)
	j := 0
	for i := 0; i < n-1; i++ {
		if nums[i] > 0 {
			if nums[i] == nums[i+1] {
				nums[i] = 2 * nums[i]
				nums[i+1] = 0
			}
			nums[j] = nums[i]
			j++
		}
	}

	if nums[n-1] > 0 {
		nums[j] = nums[n-1]
		j++
	}

	for i := j; i < n; i++ {
		nums[i] = 0
	}

	return nums
}
