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
	le, ri := 0, n-1
	for le < ri {
		if ri < n && nums[ri] == nums[le] {
			ri--
			continue
		}
		mid := le + (ri-le)/2
		if mid >= 0 && mid < n && nums[mid] <= nums[0] {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	if le < 0 || le >= n {
		return nums[0]
	}
	return nums[le]
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
