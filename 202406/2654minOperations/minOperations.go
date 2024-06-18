package main

import (
	"fmt"
)

func main() {
	fmt.Println(minOperations([]int{2, 6, 3, 4}))
	fmt.Println(minOperations([]int{6, 10, 15}))
}

func minOperations(nums []int) int {
	n, getall, cnt := len(nums), 0, 0
	for i := range nums {
		getall = gcd(getall, nums[i])
		if nums[i] == 1 {
			cnt++
		}
	}
	if getall > 1 {
		return -1
	}
	if cnt > 0 {
		return n - cnt
	}
	minCnt := n // 表示可能约数成1的最短子数组

	for i := 0; i < n; i++ {
		g := 0
		for j := i; j < n; j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				minCnt = min(minCnt, j-i+1)
			}
		}
	}

	// 第一步由长度是 minCnt 的子数组先做约数，生成一个1，需要 minCnt-1 次操作
	// 然后由这个1,对改变剩下的其他数，需要 n-1次
	return minCnt - 1 + n - 1
}

// 求最大公约数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
