package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	cnt := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 找第三点
			cnt += help(nums, i, j)
		}
	}
	return cnt
}

func help(nums []int, i, j int) int {
	n := len(nums)
	le, ri := 0, n
	for le < ri {
		mid := le + (ri-le)/2
		// 找左端点
		if mid >= 0 && mid < n && nums[i]+nums[j] > nums[mid] {
			ri = mid
		} else {
			le = mid + 1
		}
	}

	lo := le
	le, ri = 0, n
	for le < ri {
		mid := le + (ri-le+1)/2
		// 找右端点
		if mid > 0 && mid < n && nums[i]+nums[j] > nums[mid] {
			le = mid
		} else {
			ri = mid - 1
		}
	}

	return max(0, le-lo+1)
}

func help2(nums []int, i, j int) int {
	n := len(nums)
	le, ri := j+1, n
	for le < ri {
		mid := le + (ri-le)/2
		// 这样写是不可以的，得好好理解本质
		if mid >= j+1 && mid < n && nums[i]+nums[j] > nums[mid] {
			ri = mid
		} else {
			le = mid + 1
		}
	}

	lo := le
	le, ri = j+1, n
	for le < ri {
		mid := le + (ri-le+1)/2
		if mid >= j+1 && mid < n && nums[i]+nums[j] > nums[mid] {
			le = mid
		} else {
			ri = mid - 1
		}
	}

	return max(0, le-lo+1)
}
