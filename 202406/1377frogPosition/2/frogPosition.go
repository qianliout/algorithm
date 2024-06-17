package main

import (
	"fmt"
)

func main() {
	fmt.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 2, 4))
	fmt.Println(frogPosition(8, [][]int{{2, 1}, {3, 2}, {4, 1}, {5, 1}, {6, 4}, {7, 1}, {8, 7}}, 7, 7))
}

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	g := make([][]int, n+1)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	head := &Node{
		Id:    1,
		Ti:    1,
		Child: make([]*Node, 0),
	}
	visit2 := make([]bool, n+1)
	gen(head, g, visit2)
	var ans float64
	var find bool
	visit := make([]bool, n+1)
	dfs(head, &ans, &find, target, t, visit)
	if !find {
		return 0
	}
	return ans
}

func gen(head *Node, g [][]int, visit []bool) {
	if head == nil {
		return
	}
	headID := head.Id
	visit[headID] = true
	for _, x := range g[headID] {
		if visit[x] {
			continue
		}
		visit[x] = true
		head.Child = append(head.Child, &Node{
			Id:    x,
			Ti:    1,
			Child: make([]*Node, 0),
		})
	}
	for i := range head.Child {
		gen(head.Child[i], g, visit)
	}
}

func dfs(node *Node, ans *float64, find *bool, target int, t int, visit []bool) {
	if node == nil {
		return
	}

	if visit[node.Id] {
		return
	}

	if t < 0 {
		return
	}

	if *find {
		return
	}
	if node.Id == target {
		*ans = max(*ans, node.Ti)
		*find = true
		return
	}
	visit[node.Id] = true
	if len(node.Child) > 0 {
		p := float64(1) / float64(len(node.Child))
		for _, ch := range node.Child {
			ch.Ti = p * node.Ti
			dfs(ch, ans, find, target, t-1, visit)
		}
	}
}

type Node struct {
	Id    int
	Ti    float64
	Child []*Node
}
