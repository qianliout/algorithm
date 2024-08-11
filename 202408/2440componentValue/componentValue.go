package main

import (
	"fmt"
)

func main() {
	// nums = [6,2,2,2,6], edges = [[0,1],[1,2],[1,3],[3,4]]
	fmt.Println(componentValue([]int{6, 2, 2, 2, 6}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}))
}

func componentValue(nums []int, edges [][]int) int {
	sum := 0
	for _, ch := range nums {
		sum += ch
	}

	n := len(nums)
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(i, fa, target int) int
	dfs = func(x, fa, target int) int {
		sz := nums[x]
		for _, ch := range g[x] {
			if ch != fa {
				res := dfs(ch, x, target)
				if res == -1 {
					return res
				}
				sz += res
			}
		}
		if sz > target {
			return -1
		}
		if sz == target {
			return 0
		}

		return sz
	}
	for i := n; i >= 1; i-- {
		if sum%i == 0 {
			target := sum / i
			if dfs(0, -1, target) == 0 {
				return i - 1
			}
		}
	}
	return 0
}
