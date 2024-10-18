package main

import (
	"math"
)

func minAreaRect(points [][]int) int {
	cnt := make(map[Point]int)
	for _, ch := range points {
		x, y := ch[0], ch[1]
		cnt[Point{x, y}]++
	}
	inf := math.MaxInt / 100
	ans := inf
	for a := range cnt {
		for b := range cnt {
			if a.X == b.X || a.Y == b.Y {
				continue
			}
			c := Point{a.X, b.Y}
			d := Point{b.X, a.Y}

			if _, ok := cnt[c]; !ok {
				continue
			}
			if _, ok := cnt[d]; !ok {
				continue
			}

			ans = min(ans, abs(a.X-b.X)*abs(a.Y-b.Y))
		}
	}
	if ans == inf {
		return 0
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Point struct {
	X, Y int
}

// 一个矩形可以通过一条对角线也就是两个点唯一确认。那么可以先将所有的点加入哈希表，然后枚举两个点，
// 如果这两个点x坐标和y坐标都不一致就可以构成一条对角线。最后判断由这条对角线确定的矩形的另外两个点是否在哈希表中，如果存在就是一个合法的矩形并更新最小面积。
