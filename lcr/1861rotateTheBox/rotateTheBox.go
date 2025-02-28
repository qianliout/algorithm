package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func rotateTheBox(boxGrid [][]byte) [][]byte {
	m, n := len(boxGrid), len(boxGrid[0])
	rot := make([][]byte, n)
	for i := range rot {
		rot[i] = make([]byte, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rot[j][m-i-1] = boxGrid[i][j]
		}
	}
	// 下落
	for i := 0; i < m; i++ {
		last := n - 1
		for j := n - 1; j >= 0; j-- {
			if rot[j][i] == '.' {
				continue
			}
			if rot[j][i] == '*' {
				last = j - 1
				continue
			}
			if rot[j][i] == '#' {
				rot[last][i] = '#'
				if last != j { // 这一步判断最容易出错
					rot[j][i] = '.'
				}
				last--
			}
		}
	}

	return rot
}

// func rotateTheBox(boxGrid [][]byte) [][]byte {
// 	res := rotate(boxGrid)
// 	ans := help(res)
// 	return ans
// }

func rotate(boxGrid [][]byte) [][]byte {
	m, n := len(boxGrid), len(boxGrid[0])
	res := make([][]byte, n)
	for i := range res {
		res[i] = make([]byte, m)
	}
	for i := range boxGrid {
		for j := range boxGrid[i] {
			res[j][m-1-i] = boxGrid[i][j]
		}
	}
	for i := range res {
		fmt.Println(string(res[i]))
	}
	return res
}

func help(boxGrid [][]byte) [][]byte {
	m, n := len(boxGrid), len(boxGrid[0])

	for j := 0; j < n; j++ {
		for i := m - 1; i >= 0; i-- {
			ch := boxGrid[i][j]
			if ch == '*' || ch == '.' {
				continue
			}
			end := i
			// 找到最下边可以存的地地方
			for k := i + 1; k < m; k++ {
				ch2 := boxGrid[k][j]
				if ch2 == '.' {
					end = k
					continue
				}
				break
			}
			if end != i && end < m {
				boxGrid[i][j], boxGrid[end][j] = boxGrid[end][j], boxGrid[i][j]
			}
		}
	}

	return boxGrid
}
