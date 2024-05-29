package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2))
}

// 朴素的解法，能过，但是内存会用超
func maxProbability1(n int, edges [][]int, sp []float64, start int, end int) float64 {
	// inf := float64(math.MaxInt32 / 10)
	g := make([][]float64, n)
	for i := range g {
		g[i] = make([]float64, n)
	}
	for i, ch := range edges {
		x, y, z := ch[0], ch[1], sp[i]
		g[x][y] = z
		g[y][x] = z
	}
	dis := make([]float64, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[start] = 1
	done := make([]bool, n)
	for {
		x := -1
		for i, ok := range done {
			if !ok && (x < 0 || dis[i] >= dis[x]) {
				x = i
			}
		}
		if x < 0 {
			break
		}
		// 不可达
		if dis[x] < 0 {
			break
		}
		done[x] = true

		for y, d := range g[x] {
			if y == x {
				continue
			}
			dis[y] = max(dis[y], dis[x]*d)
		}
	}
	return max(dis[end], 0)
}

func maxProbability(n int, edges [][]int, sp []float64, start int, end int) float64 {
	// inf := float64(math.MaxInt32 / 10)
	g := make(map[int]map[int]float64, n)
	// for i := range g {
	// 	g[i] = make(map[int]float64, n)
	// }
	for i, ch := range edges {
		x, y, z := ch[0], ch[1], sp[i]
		if g[x] == nil {
			g[x] = make(map[int]float64)
		}
		if g[y] == nil {
			g[y] = make(map[int]float64)
		}
		g[x][y] = z
		g[y][x] = z
	}
	dis := make([]float64, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[start] = 1
	done := make([]bool, n)
	for {
		x := -1
		for i, ok := range done {
			if !ok && (x < 0 || dis[i] >= dis[x]) {
				x = i
			}
		}
		if x < 0 {
			break
		}
		// 不可达
		if dis[x] < 0 { // 初值是一个可能到的值，也就是-1
			break
		}
		done[x] = true

		for y, d := range g[x] {
			if y == x {
				continue
			}
			dis[y] = max(dis[y], dis[x]*d)
		}
	}
	return max(0, dis[end])
}
