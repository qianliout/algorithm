package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(specialPerm([]int{2, 3, 6}))
}

func specialPerm(nums []int) int {
	n := len(nums)
	mask := 1 << n
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, mask)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	// DFS函数：j表示上一个选择的数字的索引，mask表示已选择的数字集合
	var dfs func(i, set int) int

	mod := int(math.Pow10(9)) + 7
	dfs = func(j, set int) int {
		if set == (1<<n)-1 {
			return 1
		}
		if mem[j][set] != -1 {
			return mem[j][set]
		}
		ans := 0
		for i := 0; i < n; i++ {
			// if set&(1<<i) == 1 { // 这样写是错的
			if set>>i&1 == 1 {
				continue
			}
			if nums[i]%nums[j] == 0 || nums[j]%nums[i] == 0 {
				ans += dfs(i, set|(1<<i)) // 把1加入到 set 中去
				ans = ans % mod
			}
		}
		mem[j][set] = ans
		return ans
	}
	ans := 0
	for j := 0; j < n; j++ {
		ans += dfs(j, 1<<j)
		ans = ans % mod
	}

	return ans
}

/*
给你一个下标从 0 开始的整数数组 nums ，它包含 n 个 互不相同 的正整数。如果 nums 的一个排列满足以下条件，我们称它是一个特别的排列：

对于 0 <= i < n - 1 的下标 i ，要么 nums[i] % nums[i+1] == 0 ，要么 nums[i+1] % nums[i] == 0 。
请你返回特别排列的总数目，由于答案可能很大，请将它对 109 + 7 取余 后返回
2 <= nums.length <= 14
1 <= nums[i] <= 109
*/
