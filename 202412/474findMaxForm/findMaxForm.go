package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
}

func findMaxForm1(strs []string, m int, n int) int {

	nums := make([]Node, len(strs))
	for i, ch := range strs {
		nums[i] = Node{strings.Count(ch, "1"), strings.Count(ch, "0")}
	}
	le := len(strs)

	var dfs func(i, j, k int) int

	dfs = func(i, j, k int) int {
		if k < 0 {
			return 0
		}
		o, z := nums[k].one, nums[k].zero
		ans := dfs(i, j, k-1)
		if i-z >= 0 && j-o >= 0 {
			ans = max(ans, dfs(i-z, j-o, k-1)+1)
		}
		return ans
	}
	ans := dfs(m, n, le-1)

	return ans
}

type Node struct {
	one, zero int
}

// 该子集中 最多 有 m 个 0 和 n 个 1 。
func findMaxForm(strs []string, m int, n int) int {
	k := len(strs)
	nums := make([]Node, len(strs))
	for i, ch := range strs {
		nums[i] = Node{strings.Count(ch, "1"), strings.Count(ch, "0")}
	}
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, k+1)
		}
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			for c := 1; c <= k; c++ {
				o, z := nums[c-1].one, nums[c-1].zero
				f[i][j][c] = f[i][j][c-1]
				if i-z >= 0 && j-o >= 0 {
					f[i][j][c] = max(f[i][j][c], f[i-z][j-o][c-1]+1)
				}
			}
		}
	}
	return f[m][n][k]
}
