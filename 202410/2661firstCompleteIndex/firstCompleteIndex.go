package main

import "fmt"

func main() {
	fmt.Println(firstCompleteIndex([]int{1, 4, 5, 2, 6, 3}, [][]int{{4, 3, 5}, {1, 2, 6}}))

}

func firstCompleteIndex(arr []int, mat [][]int) int {
	grip := make(map[int]pair)
	for i, row := range mat {
		for j, v := range row {
			grip[v] = pair{col: i, row: j}
		}
	}
	n, m := len(mat), len(mat[0])
	col := make([]int, n)
	row := make([]int, m)
	for i, ch := range arr {
		c, r := grip[ch].col, grip[ch].row
		col[c]++
		row[r]++
		// 这里的判断条件容易出错
		if col[c] == m || row[r] == n {
			return i
		}
	}
	return 0
}

type pair struct {
	col int
	row int
}
