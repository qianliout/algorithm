package main

import (
	"fmt"
)

func main() {
	fmt.Println(matrixScore([][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}))
}

// 题不难，就是麻烦
func matrixScore(grid [][]int) int {
	row(grid)
	col(grid)
	ans := 0
	for i := 0; i < len(grid); i++ {
		ans += cal(grid[i])
	}
	return ans
}

func cal(row []int) int {
	ans := 0
	n := len(row)
	for i := 0; i < len(row); i++ {
		if row[i] == 1 {
			ans += 1 << (n - i - 1)
		}
	}
	return ans
}

func row(grid [][]int) {
	for c := 0; c < len(grid); c++ {
		o, z := len(grid[c]), len(grid[c])
		for i := 0; i < len(grid[c]); i++ {
			if grid[c][i] == 1 {
				o = min(i, o)
			} else {
				z = min(i, z)
			}
		}

		if o > z {
			for i := 0; i < len(grid[0]); i++ {
				if grid[c][i] == 0 {
					grid[c][i] = 1
				} else {
					grid[c][i] = 0
				}
			}
		}
	}
}

func col(grid [][]int) {
	for i := 0; i < len(grid[0]); i++ {
		o, z := 0, 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == 1 {
				o++
			} else {
				z++
			}
		}
		// 就要翻转
		if o < z {
			for j := 0; j < len(grid); j++ {
				if grid[j][i] == 1 {
					grid[j][i] = 0
				} else {
					grid[j][i] = 1
				}
			}
		}
	}
}
