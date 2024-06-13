package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(specialPerm([]int{2, 3, 6}))
	fmt.Println(specialPerm([]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}))
}

/*
给你一个下标从 0 开始的整数数组 nums ，它包含 n 个 互不相同 的正整数。如果 nums 的一个排列满足以下条件，我们称它是一个特别的排列：
    对于 0 <= i < n - 1 的下标 i ，要么 nums[i] % nums[i+1] == 0 ，要么 nums[i+1] % nums[i] == 0 。
请你返回特别排列的总数目，由于答案可能很大，请将它对 109 + 7 取余 后返回。
*/

func specialPerm(nums []int) int {
	mod := int(math.Pow10(9)) + 7
	var dfs func(j, set int) int
	n := len(nums)
	mask := 1 << n
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, mask+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(j, set int) int {
		if set == 0 {
			return 1 // 说明已经把所有的元素都用完了
		}
		if mem[j][set] != -1 {
			return mem[j][set]
		}

		res := 0

		for k, num := range nums {
			// 说明已经选了
			if set>>k&1 == 0 {
				continue
			}
			// 说明还没有选，选了就把这个位设置成0
			// 都是正整数，所以不有判断负数
			if num%nums[j] == 0 || nums[j]%num == 0 {
				res = (res + dfs(k, set^(1<<k))) % mod
			}
		}
		mem[j][set] = res % mod
		return mem[j][set]
	}

	res := 0
	for j := range nums {
		res += dfs(j, (mask-1)^(1<<j))
	}
	return res % mod
}
