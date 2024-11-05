package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func removeDuplicates1(nums []int) int {
	k := 0
	i := 0
	n := len(nums)
	for i < n {
		nums[k] = nums[i]
		k++
		j := i + 1
		for j < n && nums[j] == nums[i] {
			j++
		}
		i = j
	}
	fmt.Println(nums)
	return k
}

func removeDuplicates2(nums []int) int {
	start := 0
	for _, ch := range nums {
		if ch == nums[start] {
			continue
		}

		start++
		nums[start] = ch
	}
	return start + 1
}

func removeDuplicates(nums []int) int {
	start := 0
	for _, ch := range nums {
		if start >= 1 && ch == nums[start-1] {
			continue
		}
		nums[start] = ch
		start++
	}
	return start
}
