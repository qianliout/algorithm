package main

func main() {

}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m+5)
	for i := range f {
		f[i] = make([]int, n+5)
		for j := range f[i] {
			f[i][j] = 0
		}
	}
	// 初值
	ans := 0
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			f[i][0] = 1
			ans = 1 // 这里是容易出错的
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			f[0][j] = 1
			ans = 1 // 这里是容易出错的
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				f[i][j] = min(f[i-1][j], f[i][j-1], f[i-1][j-1]) + 1
			}
			ans = max(f[i][j], ans)
		}
	}
	return ans * ans
}
