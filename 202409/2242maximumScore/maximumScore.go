package main

import (
	"sort"
)

func main() {

}

func maximumScore(scores []int, edges [][]int) int {
	n := len(scores)
	g := make([][]pair, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], pair{y, scores[y]})
		g[y] = append(g[y], pair{x, scores[x]})
	}

	for i, vs := range g {
		sort.Slice(vs, func(i, j int) bool { return vs[i].Score > vs[j].Score })
		if len(vs) > 3 {
			g[i] = vs[:3]
		}
	}
	ans := -1
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		for _, v := range g[x] {
			for _, w := range g[y] {
				if v.To != y && w.To != x && v.To != w.To {
					ans = max(ans, v.Score+w.Score+scores[y]+scores[x])
				}
			}
		}
	}
	return ans
}

type pair struct {
	To    int
	Score int
}
