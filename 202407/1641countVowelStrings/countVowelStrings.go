package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(getMaximumGenerated(0))
}

func countVowelStrings(n int) int {
	// 当前已经选好 i 个字符了，且最后一个字符是 j，因为相临的字符可以一样，所以 j 的初值可以是0
	var dfs func(i, j int) int
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, 5)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, j int) int {
		if i >= n {
			return 1
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		ans := 0
		for k := j; k < 5; k++ {
			ans += dfs(i+1, k)
		}
		mem[i][j] = ans
		return ans
	}

	return dfs(0, 0)
}

func getMaximumGenerated(n int) int {
	nums := make([]int, n+1)
	nums[0] = 0
	if n > 0 {
		nums[1] = 1
	}

	for i := 0; i <= n; i++ {
		if 2*i >= 2 && 2*i <= n {
			nums[2*i] = nums[i]
		}
		if 2*i+1 >= 2 && 2*i+1 <= n {
			nums[2*i+1] = nums[i] + nums[i+1]
		}
	}

	return slices.Max(nums)
}
