package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	left, right, up, down := 0, n-1, 0, m-1
	all := m * n
	ans := make([]int, 0)
	for len(ans) < all {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
	}

	return ans[:m*n]
}
