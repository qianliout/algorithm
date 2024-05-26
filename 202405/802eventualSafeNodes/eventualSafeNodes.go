package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}))
	fmt.Println(eventualSafeNodes([][]int{{}, {0, 2, 3, 4}, {3}, {4}, {}}))
}

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	in := make([][]int, n)
	out := make([][]int, n)
	for i, ch := range graph {
		out[i] = ch
		for _, j := range ch {
			if j != i {
				in[j] = append(in[j], i)
			}
		}
	}
	end := make(map[int]bool)
	for i, ch := range out {
		if len(ch) == 0 {
			end[i] = true
		}
	}

	ans := make([]int, 0)
	for k := range end {
		ch := in[k]
		ans = append(ans, k)
		for _, x := range ch {
			flag := false
			for _, o := range out[x] {
				if !end[o] {
					flag = true
					break
				}
			}
			if !flag {
				ans = append(ans, x)
			}
		}
	}

	sort.Ints(ans)

	return ans
}
