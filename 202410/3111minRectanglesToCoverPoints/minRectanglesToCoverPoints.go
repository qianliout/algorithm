package main

import (
	"sort"
)

func main() {

}

func minRectanglesToCoverPoints(points [][]int, w int) int {
	n := len(points)
	po := make([]pair, n)
	for i, ch := range points {
		po[i] = pair{x: ch[0], y: ch[1]}
	}
	sort.Slice(po, func(i, j int) bool { return po[i].x < po[j].x })
	ans := 0
	pre := -1
	for _, ch := range po {
		if ch.x > pre {
			ans++
			pre = ch.x + w
		}
	}
	return ans
}

type pair struct {
	x, y int
}
