package main

func main() {

}

func differenceOfDistinctValues(grid [][]int) [][]int {
	n, m := len(grid), len(grid[0])
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	var cal func(c, r int) int
	cal = func(c, r int) int {
		up, down := make(map[int]int), make(map[int]int)
		// 左上
		for i, j := c-1, r-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			up[grid[i][j]]++
		}
		// 右下
		for i, j := c+1, r+1; i < n && j < m; i, j = i+1, j+1 {
			down[grid[i][j]]++
		}
		return abs(len(up) - len(down))
	}
	for i := range ans {
		for j := range ans[i] {
			ans[i][j] = cal(i, j)
		}
	}
	return ans
}

func abs(a int) int {
	if a <= 0 {
		return -a
	}
	return a
}
