package main

func main() {

}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	row, col := make([]int, m), make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			col[i] = max(col[i], grid[i][j])
			row[j] = max(row[j], grid[i][j])
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans += min(col[i], row[j]) - grid[i][j]
		}
	}

	return ans
}
