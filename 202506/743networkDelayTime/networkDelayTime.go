package main

import "slices"

func main() {

}

func networkDelayTime(times [][]int, n int, k int) int {
	g := make([][]int, n)
	inf := 1 << 26
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, ch := range times {
		x, y, z := ch[0]-1, ch[1]-1, ch[2]
		g[x][y] = z
	}
	dis := make([]int, n)
	done := make([]bool, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[k-1] = 0
	for {
		x := -1
		for i, ok := range done {
			if !ok && (x == -1 || dis[i] <= dis[x]) {
				x = i
			}
		}
		if x == -1 {
			break
		}
		// if dis[x] == inf {
		// 	return -1 // 说明有节点不可达
		// }

		done[x] = true
		for y, d := range g[x] {
			dis[y] = min(dis[y], dis[x]+d)
		}
	}

	ans := slices.Max(dis)
	if ans == inf {
		return -1
	}
	return ans
}
