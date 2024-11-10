package main

func main() {
	setZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})
}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	// -231 <= matrix[i][j] <= 231 - 1
	col := make([]int, m)
	row := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b := matrix[i][j]
			if b == 0 {
				col[i] = 1
				row[j] = 1
			}
		}
	}
	for i := 0; i < m; i++ {
		if col[i] == 1 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < n; j++ {
		if row[j] == 1 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}
