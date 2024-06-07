package main

import (
	"fmt"
)

func main() {
	fmt.Println(lastVisitedIntegers([]int{1, 2, -1, -1, -1}))
	fmt.Println(lastVisitedIntegers([]int{-1, -1, 94, 56, -1, 32, -1, -1, -1}))
}

/*
给你一个整数数组 nums ，其中 nums[i] 要么是一个正整数，要么是 -1 。我们需要为每个 -1 找到相应的正整数，我们称之为最后访问的整数。

为了达到这个目标，定义两个空数组：seen 和 ans。

从数组 nums 的头部开始遍历。

	如果遇到正整数，把它添加到 seen 的 头部。
	如果遇到 -1，则设 k 是到目前为止看到的 连续 -1 的数目(包括当前 -1)，
	    如果 k 小于等于 seen 的长度，把 seen 的第 k 个元素添加到 ans。
	    如果 k 严格大于 seen 的长度，把 -1 添加到 ans。

请你返回数组 ans。
*/

func lastVisitedIntegers(nums []int) []int {
	seen := make([]int, 0)
	ans := make([]int, 0)
	last := -1
	for i, ch := range nums {
		if ch > 0 {
			last = -1
			seen = append(seen, ch)
			continue
		}
		if ch == -1 {
			k := 1
			if last >= 0 {
				k = i - last + 1
			}

			if k <= len(seen) {
				ans = append(ans, seen[len(seen)-k])
			} else {
				ans = append(ans, -1)
			}
			if last == -1 {
				last = i
			}
		}

	}
	return ans
}

func reverse(nums []int) []int {
	ans := make([]int, len(nums))
	copy(ans, nums)
	le, ri := 0, len(ans)-1
	for le < ri {
		ans[le], ans[ri] = ans[ri], ans[le]
		le++
		ri--
	}
	return ans
}
