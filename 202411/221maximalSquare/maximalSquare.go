package main

func main() {

}

func maximalSquare(matrix [][]byte) int {
	ans := 0
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	// 初值容易出错
	for i := 0; i < m; i++ {
		f[i][0] = int(matrix[i][0]) - '0'
		ans = max(ans, f[i][0])
	}
	for j := 0; j < n; j++ {
		f[0][j] = int(matrix[0][j]) - '0'
		ans = max(ans, f[0][j])
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			f[i][j] = min(f[i-1][j], f[i][j-1], f[i-1][j-1]) + 1
			ans = max(ans, f[i][j])
		}
	}
	return ans * ans
}

// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
