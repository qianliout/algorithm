package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(nums []int) []int {
	end := len(nums) - 1
	i := 0
	for i <= end {
		if nums[i]%2 == 0 {
			i++
		} else {
			nums[i], nums[end] = nums[end], nums[i]
			end--
		}
	}
	return nums
}
