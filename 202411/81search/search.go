package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{1, 0, 1, 1, 1}, 0))
}

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target || nums[left] == target || nums[right] == target {
			return true
		} else if nums[mid] == nums[right] {
			right--
			continue
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && nums[right] > target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if nums[mid] > nums[right] {
			if nums[left] < target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}
