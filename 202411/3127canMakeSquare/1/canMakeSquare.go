package main

import (
	"fmt"
)

func main() {
	fmt.Println(canMakeSquare([][]byte{[]byte("BWB"), []byte("BWB"), []byte("BWB")}))
}

func canMakeSquare(grid [][]byte) bool {
	for i := 0; i < len(grid)-1; i++ {
		for j := 0; j < len(grid[i])-1; j++ {
			if check(grid) {
				return true
			}

			if grid[i][j] == 'W' {
				grid[i][j] = 'B'
				if check(grid) {
					return true
				}
				grid[i][j] = 'W'
			}
			if grid[i][j] == 'B' {
				grid[i][j] = 'W'
				if check(grid) {
					return true
				}
				grid[i][j] = 'B'
			}
		}
	}
	return false
}

func check(grid [][]byte) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if check2(grid, i, j) {
				return true
			}
		}
	}
	return false
}
func check2(grid [][]byte, i, j int) bool {
	if string(grid[i][j:j+2]) == "WW" && string(grid[i+1][j:j+2]) == "WW" {
		return true
	}
	if string(grid[i][j:j+2]) == "BB" && string(grid[i+1][j:j+2]) == "BB" {
		return true
	}

	return false
}
