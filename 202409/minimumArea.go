package main

func main() {

}

func minimumArea(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	left, right, up, down := m-1, 0, n-1, 0
	for i := range grid {
		for j, x := range grid[i] {
			if x != 1 {
				continue
			}
			left = min(left, j)
			right = max(right, j)
			up = min(up, i)
			down = max(down, i)
		}
	}
	return (right - left + 1) * (down - up + 1)
}
