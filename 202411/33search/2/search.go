package main

import (
	"fmt"
)

func main() {
	// fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{3, 1}, 1))
}

func search(nums []int, target int) int {
	n := len(nums)
	// 找出旋转点
	le, ri := 0, n
	for le < ri {
		mid := le + (ri-le)/2
		// 左端点
		// 这里一定是小于 [3,1]
		if mid >= 0 && mid < n && nums[mid] < nums[0] {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	// 因为各值不相同
	if target == nums[0] {
		return 0
	}
	if target > nums[0] {
		ri = le
		le = 0
	} else {
		ri = n
	}

	for le < ri {
		mid := le + (ri-le)/2
		// 左端点写法
		if mid >= 0 && mid < n && nums[mid] >= target {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	if le < 0 || le >= n || nums[le] != target {
		return -1
	}
	return le
}

// 整数数组 nums 按升序排列，数组中的值 互不相同 。
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k],
// nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
