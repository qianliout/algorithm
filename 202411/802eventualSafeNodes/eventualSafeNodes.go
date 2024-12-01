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
			// 有自闭环，但是自闭环也算一条路径
			in[j] = append(in[j], i)
		}
	}
	q := make([]int, 0)
	for i, c := range out {
		if c == 0 {
			q = append(q, i)
		}
	}
	ans := make([]int, 0)
	for len(q) > 0 {
		fir := q[0]
		q = q[1:]
		ans = append(ans, fir)
		// 找到以这些点作为出度的点
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

func eventualSafeNodes2(graph [][]int) []int {
	n := len(graph)
	in := make([][]int, n)
	out := make([]int, n)
	for i, ch := range graph {
		out[i] = len(ch)
		for _, j := range ch {
			// 有自闭环，但是自闭环也算一条路径
			in[j] = append(in[j], i)
		}
	}
	// 找到出度为0的点
	q := make([]int, 0)
	for i, c := range out {
		if c == 0 {
			q = append(q, i)
		}
	}
	ans := make([]int, 0)
	for len(q) > 0 {
		fir := q[0]
		ans = append(ans, fir)
		q = q[1:]
		// 找到以这些点作为出度的点
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
