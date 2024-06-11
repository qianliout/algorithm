package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))
}

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	n := len(nums)
	ans := make([]int, len(queries))

	for i, ch := range queries {
		le, ri := 0, n
		for le < ri {
			mi := le + (ri-le+1)/2
			if mi >= 0 && mi < n && nums[mi] <= ch {
				le = mi
			} else {
				ri = mi - 1
			}
		}
		// 检测一下是正解
		if le < 0 || le >= n || ch >= nums[le] {
			ans[i] = le + 1 // 因为下标是从0开始的，题目中问的是个数
		} else {
			ans[i] = 0
		}
	}
	return ans
}
