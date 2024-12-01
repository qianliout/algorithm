package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin([]int{1, 3, 5}))
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}

func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right = right - 1
		}
	}
	return nums[left]
}

func findMin2(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right = right - 1
		}
	}
	return nums[left]
}
