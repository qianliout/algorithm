package main

func main() {

}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	for i := n - 2; i >= 0; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}
