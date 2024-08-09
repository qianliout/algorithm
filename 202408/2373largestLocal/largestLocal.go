package main

func main() {

}

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := range ans {
		ans[i] = make([]int, n-2)
	}
	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			// 1 <= grid[i][j] <= 100
			mx := 0
			for k := i; k <= i+2; k++ {
				for m := j; m <= j+2; m++ {
					mx = max(mx, grid[k][m])
				}
			}
			ans[i][j] = mx
		}
	}
	return ans
}
