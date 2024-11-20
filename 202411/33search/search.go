package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 24, 5, 6, 7, 0, 1, 2}, 0))
}

func search(nums []int, target int) int {
	// 先找旋转点
	// 找大于等nums[0]的右边界
	n := len(nums)
	le, ri := 0, n
	for le < ri {
		// 右边界的写法
		mid := le + (ri-le+1)/2
		if mid < n && mid >= 0 && nums[mid] >= nums[0] {
			le = mid
		} else {
			// 查的是右边界，如是不在，那就缩小右边界
			ri = mid - 1
		}
	}
	// 根据target和num[0]的大于，确定target 是在左边还是右边,从而确认 le,ri的值
	if target >= nums[0] {
		// 在左边
		// ri = le
		le = 0
		// ri此时表示 大于等 nums[0]的右边界，就是大于等于 target 的右边界
	} else {
		// 在右边
		le = le + 1
		ri = n
	}
	for le < ri {
		// 找小于target 的右边界解法
		mid := le + (ri-le+1)/2
		if mid >= 0 && mid < n && nums[mid] <= target {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	return le
}

// 整数数组 nums 按升序排列，数组中的值 互不相同 。
