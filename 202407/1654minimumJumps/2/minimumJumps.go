package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumJumps([]int{18, 13, 3, 9, 8, 14}, 3, 8, 6))
	fmt.Println(minimumJumps([]int{14, 4, 18, 1, 15}, 3, 15, 9))
	fmt.Println(minimumJumps([]int{128, 178, 147, 165, 63, 11, 150, 20, 158, 144, 136}, 61, 170, 135))
	fmt.Println(minimumJumps([]int{162, 118, 178, 152, 167, 100, 40, 74, 199, 186, 26, 73, 200, 127, 30, 124, 193, 84, 184, 36,
		103, 149, 153, 9, 54, 154, 133, 95, 45, 198, 79, 157, 64, 122, 59, 71, 48, 177, 82, 35, 14, 176, 16, 108, 111, 6, 168, 31, 134, 164, 136, 72, 98}, 29, 98, 80))
}

func minimumJumps1(forbidden []int, a int, b int, x int) int {
	forbid := make(map[int]bool)
	for _, ch := range forbidden {
		forbid[ch] = true
	}

	// 这样去重会有问题，主要原因是：对于点 i来说，如果他是向前跳跳到 i，那么向后跳时，也是可以跳到 i 的
	visit := make(map[int]bool) // 防止有环，来回跳
	queue := []int{0}
	visit[0] = true
	res := 0
	maxIdx := 60000
	for len(queue) > 0 {
		lev := make([]int, 0)
		for _, first := range queue {
			if first == x {
				return res
			}
			if first+a <= maxIdx && !visit[first+a] && !forbid[first+a] {
				visit[first+a] = true
				lev = append(lev, first+a)
			}
			if first-b > 0 && !visit[first-b] && !forbid[first-b] {
				visit[first-b] = true
				lev = append(lev, first-b)
			}
		}
		res++
		queue = lev
	}
	return -1
}

func minimumJumps(forbidden []int, a int, b int, x int) int {
	forbid := make(map[int]bool)
	for _, ch := range forbidden {
		forbid[ch] = true
	}

	visit := make(map[pair]bool) // 防止有环，来回跳
	queue := []pair{pair{0, 0}}
	visit[pair{0, 0}] = true
	// visit[pair{0, 1}] = true
	maxIdx := 60000
	res := 0
	for len(queue) > 0 {
		lev := make([]pair, 0)
		for i := range queue {
			node := queue[i]
			if node.Idx == x {
				return res
			}
			pos := pair{Idx: node.Idx + a, Pos: 0}
			// 不能直接判断成 <=x+b,题目中说只能向后跳一次，但是可以向生跳一次又向前跳一下，再向后跳
			// if pos.Idx <= x+b && !visit[pos] && !forbid[pos.Idx] {
			if pos.Idx <= maxIdx && !visit[pos] && !forbid[pos.Idx] {
				visit[pos] = true
				lev = append(lev, pos)
			}
			if node.Pos < 1 {
				back := pair{Idx: node.Idx - b, Pos: 1}
				if back.Idx >= 0 && !visit[back] && !forbid[back.Idx] {
					visit[back] = true
					lev = append(lev, back)
				}
			}
		}
		res++
		queue = lev
	}
	return -1
}

type pair struct {
	Idx int
	Pos int //  表示后退了多少次
}
