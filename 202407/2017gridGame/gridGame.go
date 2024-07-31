package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(gridGame([][]int{{2, 5, 4}, {1, 5, 1}}))
}

// dfs的做法没有搞定
// 不能这样写，因为第一个机器人的最优解不是获取最大的分数
func gridGame2(grid [][]int) int64 {
	var dfs func(i, j int) int
	m, n := len(grid), len(grid[0])

	used := make([][]int, m+1)
	for i := range used {
		used[i] = make([]int, n+1)
	}

	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		if used[i][j] > 0 {
			return used[i][j]
		}
		ans := grid[i][j]
		a := dfs(i+1, j)
		b := dfs(i, j+1)
		c := max(a, b) + ans
		used[i][j] = c
		return c
	}
	var dfs2 func(i, j int)
	dfs2 = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if dfs(i, j) == used[i][j] {
			grid[i][j] = 0
		}
		dfs2(i+1, j)
		dfs2(i, j+1)
	}

	dfs(0, 0)
	dfs2(0, 0)
	// used := make([][]int, m+1)
	for i := range used {
		used[i] = make([]int, n+1)
	}
	ans := dfs(0, 0)
	return int64(ans)
}

// 由于矩阵只有两行，我们可以枚举第一个机器人的拐弯位置，于是剩余的点数被划分成了两部分，即第一行的后缀和第二行的前缀，另一个机器人只能在这两部分中取最大值。为了最小化第二个机器人收集的点数，取所有最大值中的最小值，即为答案。
func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	ans := math.MaxInt64 / 10
	left := 0
	for _, ch := range grid[0] {
		left += ch
	}
	right := 0
	for i := 0; i < n; i++ {
		left -= grid[0][i]
		// 第二个机器人只能走一条路
		ans = min(ans, max(left, right))
		right += grid[1][i]
	}

	return int64(ans)
}
