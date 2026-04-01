package main

func main() {

}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	for i := n - 2; i >= 0; i++ {
		tri := triangle[i]
		for j, c := range tri {
			tri[j] = min(triangle[i+1][j], triangle[i+1][j+1]) + c
		}
	}
	return triangle[0][0]
}
