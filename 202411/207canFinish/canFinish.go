package main

import (
	"fmt"
)

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	in := make([]int, numCourses)
	for _, ch := range prerequisites {
		x, y := ch[0], ch[1]
		g[y] = append(g[y], x)
		in[x]++
	}
	q := make([]int, 0)
	for i, c := range in {
		if c == 0 {
			q = append(q, i)
		}
	}
	cnt := 0
	for len(q) > 0 {
		fir := q[0]
		cnt++
		q = q[1:]
		for _, x := range g[fir] {
			in[x]--
			if in[x] == 0 {
				q = append(q, x)
			}
		}
	}
	return cnt == numCourses
}
