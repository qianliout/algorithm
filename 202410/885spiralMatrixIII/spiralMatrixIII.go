package main

func main() {

}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	result := make([][]int, 0)
	visit := 0
	r, c := rStart, cStart
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIndex := 0
	// 确定初始边界
	// 确定边界的方法得掌握
	up, down, left, right := rStart-1, rStart+1, cStart-1, cStart+1
	for visit < rows*cols {
		if in(rows, cols, r, c) {
			result = append(result, []int{r, c})
			visit++
		}
		// 向左达到边界
		if dirIndex%4 == 0 && c == right {
			// 方向改成向下
			dirIndex++
			// 左边的边界增加一
			right++
		} else if dirIndex%4 == 1 && r == down {
			dirIndex++
			down++
		} else if dirIndex%4 == 2 && c == left {
			dirIndex++
			left--
		} else if dirIndex%4 == 3 && r == up {
			dirIndex = 0
			up--
		}
		r, c = r+dirs[dirIndex][0], c+dirs[dirIndex][1]
	}
	return result
}

func in(m, n, r, c int) bool {
	if r < 0 || r >= m || c < 0 || c >= n {
		return false
	}
	return true
}
