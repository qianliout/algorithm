package main

func main() {

}

func minimumOperations(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cnt := make([][]int, n)
	for i := range cnt {
		cnt[i] = make([]int, 10)
	}
	for _, row := range grid {
		for j, v := range row {
			cnt[j][v]++
		}
	}
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, 11)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	// i 表示在第i列，从上到下为一列
	// j 表示i+1列已经全部变成 j 了
	// 结果是在 上述情况下，会保留的数的总和
	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i < 0 {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		res := 0
		for k, c := range cnt[i] {
			if k != j {
				res = max(res, dfs(i-1, k)+c)
			}
		}
		mem[i][j] = res
		return res
	}
	// dfs(n-1,10) 表示从 n-1表开始看，因为 n-1列是最后一列，所以可以认为前面是一个不存在的数，为了好缓存，设置成10
	return m*n - dfs(n-1, 10)
}

// 如果下面相邻格子存在的话，它们的值相等，也就是 grid[i][j] == grid[i + 1][j]（如果存在）。
// 如果右边相邻格子存在的话，它们的值不相等，也就是 grid[i][j] != grid[i][j + 1]（如果存在）。
