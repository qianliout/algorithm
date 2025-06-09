package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	le, up := n-1, 0
	for le >= 0 && up < m {
		c := matrix[up][le]
		if c == target {
			return true
		} else if c > target {
			le--
		} else if c < target {
			up++
		}
	}
	return false
}
