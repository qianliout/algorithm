package main

import (
	"fmt"
)

func main() {
	fmt.Println(canMakeSquare([][]byte{[]byte("BWB"), []byte("BWB"), []byte("BWB")}))
}

func canMakeSquare(grid [][]byte) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if check(grid, i, j) {
				return true
			}
		}
	}
	return false
}
func check(grid [][]byte, i, j int) bool {
	cnt := make(map[byte]int)
	cnt[grid[i][j]]++
	cnt[grid[i+1][j]]++
	cnt[grid[i][j+1]]++
	cnt[grid[i+1][j+1]]++
	return cnt['W'] >= 3 || cnt['B'] >= 3
}
