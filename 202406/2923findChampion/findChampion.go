package main

func main() {

}

func findChampion(grid [][]int) int {
	n := len(grid)
	for row := 0; row < n; row++ {
		ok := true
		for col := 0; col < n; col++ {
			if col != row && grid[col][row] == 1 {
				ok = false
				break
			}
		}
		if ok {
			return row
		}
	}
	return 0
}
