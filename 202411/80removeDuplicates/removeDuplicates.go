package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 1, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{1, 2}))
	fmt.Println(removeDuplicates([]int{1, 1}))
}

func removeDuplicates2(nums []int) int {
	start := 0
	n := len(nums)
	cnt := 1
	for i := 1; i < n; i++ {
		if nums[i] == nums[start] {
			cnt++
			if cnt > 2 {
				continue
			}
			start++
			nums[start] = nums[i]
		} else {
			cnt = 1
			start++
			nums[start] = nums[i]
		}
	}
	fmt.Println(nums)
	return start + 1
}

func removeDuplicates(nums []int) int {
	start := 0
	for _, ch := range nums {
		if start >= 2 && ch == nums[start-1] && ch == nums[start-2] {
			continue
		}
		nums[start] = ch
		start++
	}
	return start + 1
}
