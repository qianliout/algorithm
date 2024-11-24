package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxRemovals("bcda", "d", []int{0, 3}))
	fmt.Println(maxRemovals("abbaa", "aba", []int{0, 1, 2}))
}

func maxRemovals(source string, pattern string, targetIndices []int) int {
	m, n := len(source), len(pattern)
	tar := make(map[int]int)
	for _, ch := range targetIndices {
		tar[ch] = 1
	}
	inf := math.MaxInt64 / 10
	var dfs func(i, j int) int
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, j int) int {
		if i < j {
			return -inf
		}
		if i < 0 {
			return 0
		}
		if j < 0 {
			ans := 0
			for k := 0; k <= i; k++ {
				ans += tar[k]
			}
			return ans
		}

		if mem[i][j] != -1 {
			return mem[i][j]
		}
		// 选i 进行删除,如果 i 在 targetIndices 中，那么删除次数加一
		ans := dfs(i-1, j) + tar[i]
		// 不选 i 进行删除
		if j >= 0 && source[i] == pattern[j] {
			ans = max(ans, dfs(i-1, j-1))
		} else {
			// 如果j<0 或source[i] == pattern[j],那么此时就不可能是子序列了，这种情况下，是不合法的
			ans = max(ans, -inf)
		}
		mem[i][j] = ans
		return ans
	}
	ans := dfs(m-1, n-1)
	return ans
}

// 不选 source[i]，问题变成要使 pattern[0] 到 pattern[j] 是 source[0] 到 source[i−1] 的子序列，最多可以进行多少次删除操作，即 dfs(i−1,j)。如果 i 在 targetIndices 中，那么删除次数加一。
// 如果 source[i]=pattern[j]，那么匹配（都选），问题变成要使 pattern[0] 到 pattern[j−1] 是 source[0] 到 source[i−1] 的子序列，最多可以进行多少次删除操作，即 dfs(i−1,j−1)。
