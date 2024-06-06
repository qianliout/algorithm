package main

import (
	"fmt"
)

func main() {
	fmt.Println(goodDaysToRobBank([]int{5, 3, 3, 3, 5, 6, 2}, 2))
	fmt.Println(goodDaysToRobBank([]int{1, 1, 1, 1, 1}, 0))
	fmt.Println(goodDaysToRobBank([]int{1, 2, 3, 4}, 1))
	fmt.Println(goodDaysToRobBank([]int{1, 2, 5, 4, 1, 0, 2, 4, 5, 3, 1, 2, 4, 3, 2, 4, 8}, 2))
}

func goodDaysToRobBank(nums []int, ti int) []int {
	n := len(nums)
	ans := make([]int, 0)
	pre := make([]int, n) // 非递增的个数
	suf := make([]int, n) // 非递减的个数
	for i := 0; i < n; i++ {
		pre[i] = 1
		suf[i] = 1
	}
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			pre[i] = pre[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			suf[i] = suf[i+1] + 1
		}
	}
	for i := ti; i+ti < n; i++ {
		if pre[i] > ti && suf[i] > ti {
			ans = append(ans, i)
		}
	}
	return ans
}
