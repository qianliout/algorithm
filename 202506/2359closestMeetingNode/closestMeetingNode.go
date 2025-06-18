package main

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
	dis1 := cacal(edges, node1)
	dis2 := cacal(edges, node2)
	ans := -1
	minDis := n // 技巧
	for i := 0; i < n; i++ {
		d := max(dis1[i], dis2[i]) // 如果d==n 说明两个节点不可达
		if d < minDis {
			minDis = d
			ans = i // 更新答案为当前节点
		}
	}
	return ans
}

type node struct {
	x, fa int
}

// 请你返回一个从 node1 和 node2 都能到达节点的编号，使节点 node1 和节点 node2 到这个节点的距离 较大值最小化。
// 如果有多个答案，请返回 最小 的节点编号。如果答案不存在，返回 -1 。

func cacal(edges []int, start int) []int {
	n := len(edges)
	dis := make([]int, n)
	for i := range dis {
		dis[i] = n // 初始化为最大值
	}
	d := 0
	for start >= 0 && dis[start] == n {
		dis[start] = d
		d++
		start = edges[start]
	}
	return dis
}
