package main

import (
	"fmt"
	"math"
)

func main() {
	hats := [][]int{{1, 3, 5, 10, 12, 13, 14, 15, 16, 18, 19, 20, 21, 27, 34, 35, 38, 39, 40}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}, {3, 7, 10, 12, 13, 14, 15, 17, 21, 25, 29, 31, 35, 40}, {2, 3, 7, 8, 9, 11, 12, 14, 15, 16, 17, 18, 19, 20, 22, 24, 25, 28, 29, 32, 33, 34, 35, 36, 38}, {6, 12, 17, 20, 22, 26, 28, 30, 31, 32, 34, 35}, {1, 4, 6, 7, 12, 13, 14, 15, 21, 22, 27, 28, 30, 31, 32, 35, 37, 38, 40}, {6, 12, 21, 25, 38}, {1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 34, 35, 36, 37, 38, 39, 40}}
	fmt.Println(numberWays(hats)) // 842465346
}

func numberWays(hats [][]int) int {
	n := len(hats) // 有多少个人
	ht := 40       // 这里可以有程序算出最大的帽子是多少
	m := 1<<n - 1
	base := int(math.Pow10(9)) + 7
	// 每个帽子有多少人喜欢
	g := make([][]int, ht+1)
	for i, hat := range hats {
		for _, h := range hat {
			g[h] = append(g[h], i)
		}
	}
	mem := make([][]int, ht+1)
	for i := range mem {
		mem[i] = make([]int, m)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	// i是指选了第 i 个帽子，j 表示当前所选人员集合
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 容易出错点1：这里j==m 一定要先判断，然后再判断 j>ht
		if j == m { // 每个人都选了帽子，所以出现了一种可选的集合
			return 1
		}

		// ht是包含的，所以这里一定是>
		if i > ht {
			return 0
		}

		if mem[i][j] != -1 {
			return mem[i][j]
		}
		res := dfs(i+1, j) // 所有人都不选第 i 号帽子的选法

		for _, p := range g[i] {
			if (j>>p)&1 != 0 {
				// 这个人已经选了
				continue
			}
			res = (res + dfs(i+1, j|(1<<p)))
		}
		mem[i][j] = res
		return res
	}
	res := dfs(0, 0)

	return res % base
}
