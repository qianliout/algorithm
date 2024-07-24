package main

func main() {

}

// 要求必须写出时间复杂度为 O(m log(n)) 或 O(n log(m)) 的算法
// 这种解法不能满足题目的意思
func findPeakGrid(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if check(mat, i, j) {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func check(mat [][]int, i, j int) bool {
	if i-1 >= 0 && mat[i-1][j] >= mat[i][j] {
		return false
	}
	if i+1 < len(mat) && mat[i+1][j] >= mat[i][j] {
		return false
	}
	if j-1 >= 0 && mat[i][j-1] >= mat[i][j] {
		return false
	}
	if j+1 < len(mat[i]) && mat[i][j+1] >= mat[i][j] {
		return false
	}
	return true
}
