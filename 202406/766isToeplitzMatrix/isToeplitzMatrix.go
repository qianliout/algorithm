package main

func main() {

}

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return true
	}
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		pre := matrix[0][i]
		for j, k := 1, i+1; j < n && k < m; j, k = j+1, k+1 {
			if matrix[j][k] != pre {
				return false
			}
		}
	}

	for j := 0; j < n; j++ {
		pre := matrix[j][0]
		for i, k := j+1, 1; k < m && i < n; i, k = i+1, k+1 {
			if matrix[i][k] != pre {
				return false
			}
		}
	}

	return true
}
