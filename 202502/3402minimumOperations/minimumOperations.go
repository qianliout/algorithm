package main

func main() {

}

func minimumOperations(grid [][]int) int {
	cnt := 0
	n, m := len(grid), len(grid[0])
	for j := 0; j < m; j++ {
		for i := 1; i < n; i++ {
			if grid[i][j] > grid[i-1][j] {
				continue
			}
			cnt += grid[i-1][j] - grid[i][j] + 1
			grid[i][j] = grid[i-1][j] + 1 // 这一步很容易出错
		}
	}
	return cnt
}
