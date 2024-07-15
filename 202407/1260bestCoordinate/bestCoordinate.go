package main

import (
	"math"
)

func main() {

}

func bestCoordinate(towers [][]int, radius int) []int {
	n := 110
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}
	x1, y1, ans := 0, 0, 0
	for _, t := range towers {
		x, y, q := t[0], t[1], t[2]
		// 辐射
		for i := max(0, x-radius); i <= x+radius; i++ {
			for j := max(0, y-radius); j <= y+radius; j++ {
				// cal 一定要返回 float，不能取整数
				d := cal(x, y, i, j)
				if d > float64(radius) {
					continue
				}
				// 只能是在计算之后取整数
				g[i][j] += int(math.Floor(float64(q) / (1 + d)))

				if g[i][j] > ans {
					ans = g[i][j]
					x1, y1 = i, j
				} else if g[i][j] == ans {
					// 请你返回数组 [cx, cy] ，表示 信号强度 最大的 整数 坐标点 (cx, cy) 。如果有多个坐标网络信号一样大，请你返回字典序最小的 非负 坐标。
					if i < x1 || (i == x1 && j < y1) {
						x1, y1 = i, j
					}
				}
			}
		}
	}

	return []int{x1, y1}
}

func cal(x, y, i, j int) float64 {
	a := i - x
	b := j - y
	return math.Sqrt(float64(a*a + b*b))
}
