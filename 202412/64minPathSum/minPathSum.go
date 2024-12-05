package main

func main() {

}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m+10)
	for i := range f {
		f[i] = make([]int, n+10)
	}
	f[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		f[i][0] = f[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		f[0][j] = f[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]
		}
	}

	return f[m-1][n-1]
}

func uniquePathsWithObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m+10)
	for i := range f {
		f[i] = make([]int, n+10)
	}
	f[0][0] = 1
	if grid[0][0] == 1 {
		return 0
	}
	for i := 1; i < m; i++ {
		if grid[i][0] == 1 {
			break
		}
		f[i][0] = 1
	}
	for j := 1; j < n; j++ {
		if grid[0][j] == 1 {
			break
		}
		f[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}

	return f[m-1][n-1]
}
