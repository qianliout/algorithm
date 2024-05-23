package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isEscapePossible([][]int{{0, 1}, {1, 0}}, []int{0, 0}, []int{0, 2}))
}

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var base = int(math.Pow(10, 6))
var Source []int
var Target []int
var maxPoint int
var BlockM map[int]bool

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	Source = source
	Target = target

	m := len(blocked)
	maxPoint = (m + 1) * (m + 1) / 2
	BlockM = make(map[int]bool)
	for _, ch := range blocked {
		BlockM[ch[0]*base+ch[1]] = true
	}
	return bfs(source[0], source[1], true) && bfs(target[0], target[1], false)
}

func bfs(col, row int, start bool) bool {
	visit := make(map[int]bool)
	queue := make([][]int, 0)
	queue = append(queue, []int{col, row})
	visit[col*base+row] = true

	for len(queue) > 0 && len(visit) <= maxPoint {
		first := queue[0]

		queue = queue[1:]
		x, y := first[0], first[1]
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= base || ny < 0 || ny >= base {
				continue
			}
			nu := nx*base + ny
			if BlockM[nu] {
				continue
			}
			if visit[nu] {
				continue
			}
			if isArrive(nx, ny, start) {
				return true
			}
			visit[nu] = true
			queue = append(queue, []int{nx, ny})
		}
	}

	return len(visit) > maxPoint
}

// 从启点出发到了终点或者从终点出发到了启hko
func isArrive(col, row int, start bool) bool {
	if start {
		if col == Target[0] && row == Target[1] {
			return true
		}
	}
	if !start {
		if col == Source[0] && row == Source[1] {
			return true
		}
	}

	return false
}
