package main

func main() {

}

func numberOfSubmatrices(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	sum1 := make([][]int, m+1)
	sum2 := make([][]int, m+1)
	for i := range sum1 {
		sum1[i] = make([]int, n+1)
		sum2[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			sum1[i+1][j+1] = sum1[i+1][j] + sum1[i][j+1] - sum1[i][j]
			sum2[i+1][j+1] = sum2[i+1][j] + sum2[i][j+1] - sum2[i][j]
			if grid[i][j] == 'X' {
				sum1[i+1][j+1] = sum1[i+1][j] + sum1[i][j+1] - sum1[i][j] + 1
			}
			if grid[i][j] == 'Y' {
				sum2[i+1][j+1] = sum2[i+1][j] + sum2[i][j+1] - sum2[i][j] + 1
			}
		}
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if sum1[i+1][j+1] >= 1 && sum1[i+1][j+1] == sum2[i+1][j+1] {
				ans++
			}
		}
	}
	return ans
}
