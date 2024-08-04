package main

func main() {

}

func checkValid(matrix [][]int) bool {
	n := len(matrix)
	for _, ch := range matrix {
		cnt := make([]int, n)
		for _, i := range ch {
			if i > n || i < 1 {
				return false
			}
			cnt[i-1]++
		}
		for i := 0; i < n; i++ {
			if cnt[i] <= 0 {
				return false
			}
		}
	}

	for j := 0; j < n; j++ {
		cnt := make([]int, n)
		for i := 0; i < n; i++ {
			a := matrix[i][j]
			if a > n || a < 1 {
				return false
			}
			cnt[a-1]++
		}
		for i := 0; i < n; i++ {
			if cnt[i] <= 0 {
				return false
			}
		}
	}
	return true
}
