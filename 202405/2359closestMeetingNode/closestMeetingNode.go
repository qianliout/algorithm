package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(closestMeetingNode([]int{1, 2, -1}, 0, 2))
	fmt.Println(closestMeetingNode([]int{2, 2, 3, -1}, 0, 1))
}

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	misDis := math.MaxInt32
	ans := -1
	// 先算出 node1,node2到各个点的距离
	n1 := cac(edges, node1)
	n2 := cac(edges, node2)
	for i := 0; i < len(n1); i++ {
		if n1[i] == -1 || n2[i] == -1 {
			continue
		}
		// 取两个距离中的较大值
		d := max(n1[i], n2[i])
		// 较大值中取较小值
		if misDis > d {
			misDis = d
			ans = i
		}
	}
	return ans
}

// 算出 start 这个点到各个点的距离
func cac(edges []int, start int) []int {
	dis := make([]int, len(edges))
	for i := range dis {
		dis[i] = -1
	}
	x := start
	d := 0
	for x != -1 && dis[x] == -1 {
		dis[x] = d
		d++
		x = edges[x]
	}
	return dis
}
