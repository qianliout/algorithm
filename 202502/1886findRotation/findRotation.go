package main

func main() {

}

func findRotation(mat [][]int, target [][]int) bool {

	for i := 0; i < 4; i++ {
		if check(mat, target) {
			return true
		}
		mat = rotate(mat)
	}

	return false
}

func rotate(boxGrid [][]int) [][]int {
	m, n := len(boxGrid), len(boxGrid[0])
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
	}
	for i := range boxGrid {
		for j := range boxGrid[i] {
			res[j][m-1-i] = boxGrid[i][j]
		}
	}
	return res
}
func check(a [][]int, b [][]int) bool {
	m, n := len(a), len(a[0])

	if m != len(b) || n != len(b[0]) {
		return false
	}
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
