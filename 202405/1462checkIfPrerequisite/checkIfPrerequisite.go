package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkIfPrerequisite(2, [][]int{{1, 0}}, [][]int{{1, 0}, {0, 1}}))
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	in, out := make([][]int, numCourses), make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		in[i] = make([]int, 0)
		out[i] = make([]int, 0)
	}
	for _, no := range prerequisites {
		x, y := no[0], no[1]
		out[x] = append(out[x], y)
		in[y] = append(in[y], x)
	}

	ans := make([]bool, len(queries))
	for i := range queries {
		ans[i] = bfs(numCourses, in, out, queries[i])
	}
	return ans
}

func bfs(n int, in, out [][]int, pos []int) bool {
	x, y := pos[0], pos[1]
	vis := make([]bool, n)
	queue := make([]int, 0)
	queue = append(queue, y)
	for len(queue) > 0 {
		no := queue[0]
		queue = queue[1:]
		if vis[no] {
			continue
		}
		vis[no] = true
		for _, j := range in[no] {
			if j == x {
				return true
			}
			queue = append(queue, j)
		}
	}
	return false
}
