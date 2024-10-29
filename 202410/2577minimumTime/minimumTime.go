package main

func main() {

}

func minimumTime(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dis := make([]int, m*n)

}
