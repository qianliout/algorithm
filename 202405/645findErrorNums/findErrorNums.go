package main

import (
	"fmt"
)

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}

func findErrorNums(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i, ch := range nums {
		if i+1 != ch {
			return []int{ch, i + 1}
		}
	}
	return []int{}
}
