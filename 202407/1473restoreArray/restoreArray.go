package main

import (
	"fmt"
)

func main() {
	fmt.Println(restoreArray([][]int{{2, 1}, {3, 4}, {3, 2}}))
	fmt.Println(restoreArray([][]int{{100000, -100000}}))
}

func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	g := make(map[int][]int)
	in := make(map[int]int)
	for _, ch := range adjacentPairs {
		a, b := ch[0], ch[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
		in[a]++
		in[b]++
	}
	start := 0 // 题目保证有解
	for k, v := range in {
		if v == 1 {
			start = k
			break
		}
	}
	ans := []int{start}
	used := make(map[int]bool)
	used[start] = true
	for len(ans) < n {
		next := g[start]
		if len(next) == 0 {
			break
		}
		for _, k := range next {
			if used[k] {
				continue
			}
			used[k] = true
			start = k
			ans = append(ans, k)
		}
	}
	return ans
}
