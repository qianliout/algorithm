package main

func main() {

}
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat) == 0 || len(mat) == 0 {
		return mat
	}
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}
	m := len(mat[0])
	for i := 0; i < r*c; i++ {
		x, y := find(i, m)
		a, b := find(i, c)
		res[a][b] = mat[x][y]
	}

	return res
}

func find(n int, c int) (int, int) {
	return n / c, n % c
}
