package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(verifyTreeOrder([]int{4, 6, 5, 9, 8}))
}

func verifyTreeOrder(postorder []int) bool {
	var dfs func(nums []int, mx, mi int) bool
	dfs = func(nums []int, mx, mi int) bool {
		if len(nums) == 0 {
			return true
		}
		n := len(nums)
		rootV := nums[n-1]
		if rootV > mx || rootV < mi {
			return false
		}
		idx := findInx(nums[:n-1], rootV)
		left := nums[:idx]
		right := nums[idx : n-1]
		a := dfs(left, rootV, mi)
		b := dfs(right, mx, rootV)
		return a && b
	}
	inf := math.MaxInt64 - 10
	return dfs(postorder, inf, -inf)
}

// postorder 中无重复数字

// 找一个 idx，idx-1<va && id>va
func findInx(nums []int, va int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	if nums[0] > va {
		return 0
	}
	if nums[n-1] < va {
		return n
	}

	for i := 0; i < n; i++ {
		if nums[i] > va {
			return i
		}
	}
	return n
}
