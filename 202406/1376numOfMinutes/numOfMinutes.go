package main

import (
	"fmt"
)

func main() {
	fmt.Println(numOfMinutes(11, 4, []int{5, 9, 6, 10, -1, 8, 9, 1, 9, 3, 4}, []int{0, 213, 0, 253, 686, 170, 975, 0, 261, 309, 337}))
}

// 不对
func numOfMinutes1(n int, headID int, manager []int, informTime []int) int {
	g := make([][]int, n)
	for i, x := range manager {
		if x != -1 {
			g[x] = append(g[x], i)
		}
	}
	ans := 0
	queue := make([]int, 0)
	queue = append(queue, headID)

	for len(queue) > 0 {
		ti := 0
		lev := make([]int, 0)
		for _, no := range queue {
			ti = max(ti, informTime[no])
			for _, nx := range g[no] {
				lev = append(lev, nx)
			}
		}
		// 这样写不对，不对的原因是，每一层传给下一层后，马上就会再向下传递，不会第这一层传完
		ans += ti
		queue = lev
	}
	return ans
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	g := make([][]int, n)
	for i, x := range manager {
		if x != -1 {
			g[x] = append(g[x], i)
		}
	}
	head := &Node{
		Id:    headID,
		Ti:    informTime[headID],
		Child: make([]*Node, 0),
	}
	gen(head, informTime, g)
	var ans int
	dfs(head, &ans)
	return ans
}

func gen(head *Node, informTime []int, g [][]int) {
	if head == nil {
		return
	}
	headID := head.Id
	for _, x := range g[headID] {
		head.Child = append(head.Child, &Node{
			Id:    x,
			Ti:    informTime[x],
			Child: make([]*Node, 0),
		})
	}
	for i := range head.Child {
		gen(head.Child[i], informTime, g)
	}
}

func dfs(node *Node, ans *int) {
	if node == nil {
		return
	}
	*ans = max(*ans, node.Ti)
	for i := range node.Child {
		node.Child[i].Ti += node.Ti
		dfs(node.Child[i], ans)
	}
}

type Node struct {
	Id    int
	Ti    int
	Child []*Node
}
