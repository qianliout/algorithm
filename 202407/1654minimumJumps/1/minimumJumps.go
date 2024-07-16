package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumJumps([]int{14, 4, 18, 1, 15}, 3, 15, 9))
	fmt.Println(minimumJumps([]int{128, 178, 147, 165, 63, 11, 150, 20, 158, 144, 136}, 61, 170, 135))
	fmt.Println(minimumJumps([]int{162, 118, 178, 152, 167, 100, 40, 74, 199, 186, 26, 73, 200, 127, 30, 124, 193, 84, 184, 36,
		103, 149, 153, 9, 54, 154, 133, 95, 45, 198, 79, 157, 64, 122, 59, 71, 48, 177, 82, 35, 14, 176, 16, 108, 111, 6, 168, 31, 134, 164, 136, 72, 98}, 29, 98, 80))
}

func minimumJumps(forbidden []int, a int, b int, x int) int {
	var dfs func(i, pre int) int
	forbid := make(map[int]bool)
	for _, ch := range forbidden {
		forbid[ch] = true
	}

	visit := make(map[pair]bool) // 防止有环，来回跳

	inf := math.MaxInt / 10

	visit[pair{0, 0}] = true // 主要考查这里的去重，容易出错点1
	maxIdx := 60000          // 这里也是容易出错的
	mem := make([][]int, maxIdx+10)
	for i := range mem {
		mem[i] = make([]int, 2)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	dfs = func(i, pre int) int {
		if i < 0 || i > maxIdx {
			return inf
		}
		if i == x {
			return 0
		}
		if mem[i][pre] != -1 {
			return mem[i][pre]
		}
		res := inf
		// 那就只能是向前跳了
		if !forbid[i+a] && !visit[pair{i + a, 0}] {
			visit[pair{i + a, 0}] = true
			nex := dfs(i+a, 0)
			// visit[i+a] = false
			res = min(res, nex+1)
		}
		//  尝试向后跳
		if pre < 1 && i-b >= 0 && !forbid[i-b] && !visit[pair{i - b, 1}] {
			visit[pair{i - b, 1}] = true
			nex := dfs(i-b, pre+1)
			visit[pair{i - b, 1}] = false
			res = min(res, nex+1)
		}
		mem[i][pre] = res
		return res
	}
	res := dfs(0, 0)
	if res >= inf {
		return -1
	}
	return res
}

type pair struct {
	idx int
	pos int
}
