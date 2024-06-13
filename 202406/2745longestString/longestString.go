package main

import (
	"fmt"
	"slices"
)

func main() {
	// fmt.Println(longestString(1, 39, 14))
	fmt.Println(longestString(3, 1, 21))
	fmt.Println(longestString2(3, 1, 21))
	fmt.Println(longestString2(2, 5, 1))
}

// 这才是这个题目的正解
func longestString(x int, y int, z int) int {
	ans := min(x, y) * 2
	if x != y {
		ans++
	}

	return (ans + z) * 2
}

// 递归的方法没有想好。因为是需要尝试3种选择方法
func longestString2(x int, y int, z int) int {

	need := map[int][]int{
		0: []int{1},
		1: []int{0, 2},
		2: []int{0, 2},
	}

	var dfs func(cost []int, pre int) int

	dfs = func(cost []int, pre int) int {
		if slices.Max(cost) == 0 {
			return 0
		}
		nex := need[pre]
		ans := 0
		for j := range nex {
			i := nex[j]
			if cost[i] > 0 {
				cost[i] -= 1
				ans = max(ans, dfs(cost, i)+2)
				cost[i] += 1
			}
		}
		return ans
	}
	ans := 0
	for i := 0; i < 1; i++ {
		cost := []int{x, y, z}
		if cost[i] > 0 {
			cost[i]--
			ans = max(ans, dfs(cost, i)+2)
		}
	}
	return ans
}
