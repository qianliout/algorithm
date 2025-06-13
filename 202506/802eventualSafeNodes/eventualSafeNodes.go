package main

import (
	"sort"
)

func main() {

}

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	in := make([][]int, n)
	out := make([]int, n)
	for i, ch := range graph {
		out[i] = len(ch)
		for _, j := range ch {
			in[j] = append(in[j], i)
		}
	}
	q := make([]int, 0)
	for i, o := range out {
		if o == 0 {
			q = append(q, i)
		}
	}
	ans := make([]int, 0)
	for len(q) > 0 {
		fir := q[0]
		q = q[1:]
		ans = append(ans, fir)
		for _, j := range in[fir] {
			out[j]--
			if out[j] == 0 {
				q = append(q, j)
			}
		}

	}
	sort.Ints(ans)
	return ans
}
