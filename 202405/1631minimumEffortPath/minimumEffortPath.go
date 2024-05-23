package main

import (
	"fmt"
	"sort"

	. "outback/algorithm/common/unionfind"
	. "outback/algorithm/common/utils"
)

func main() {
	heights := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	fmt.Println(minimumEffortPath(heights))
}

func minimumEffortPath(heights [][]int) int {
	edges := make([][]int, 0)
	m, n := len(heights), len(heights[0])

	for i := range heights {
		for j, ch := range heights[i] {
			pos := i*n + j
			if i < m-1 {
				edges = append(edges, []int{AbsSub(heights[i+1][j], ch), pos, pos + n})
			}
			if j < n-1 {
				edges = append(edges, []int{AbsSub(heights[i][j+1], ch), pos, pos + 1})
			}
		}
	}
	sort.Sort(Edge(edges))

	uf := NewRankUnionFind(m*n + 1)

	for _, ch := range edges {
		uf.Union(ch[1], ch[2])
		if uf.IsConnected(0, m*n-1) {
			return ch[0]
		}
	}
	return 0

}

type Edge [][]int

func (vi Edge) Len() int {
	return len(vi)
}

func (vi Edge) Less(i, j int) bool {
	return vi[i][0] <= vi[j][0]
}

func (vi Edge) Swap(i, j int) {
	vi[i], vi[j] = vi[j], vi[i]
}
