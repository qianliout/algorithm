package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumMoves([][]int{{1, 1, 0}, {1, 1, 1}, {1, 2, 1}}))
}

type pair struct {
	x, y, idx int
}

func minimumMoves(grid [][]int) int {
	from, to := make([]pair, 0), make([]pair, 0)
	for i := range grid {
		for j, ch := range grid[i] {
			for ch > 1 {
				from = append(from, pair{i, j, len(from)})
				ch--
			}
			if ch == 0 {
				to = append(to, pair{i, j, len(to)})
			}
		}
	}
	pp := make([][]pair, 0)
	used := make(map[int]bool)

	dfs(from, used, []pair{}, &pp)
	ans := math.MaxInt
	for i := 0; i < len(pp); i++ {
		cnt := 0
		for j, ch := range pp[i] {
			cnt += abs(ch.x-to[j].x) + abs(ch.y-to[j].y)
		}
		ans = min(ans, cnt)
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dfs(from []pair, used map[int]bool, path []pair, ans *[][]pair) {
	if len(path) == len(from) {
		*ans = append(*ans, append([]pair{}, path...))
		return
	}
	for i := 0; i < len(from); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, from[i])
		dfs(from, used, path, ans)
		used[i] = false
		path = path[:len(path)-1]
	}
}
