package main

import (
	"fmt"
)

func main() {
	fmt.Println(minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
}

func minJumps(arr []int) int {
	n := len(arr)
	cnt := make(map[int][]int)
	// 应付等值的情况
	for i, ch := range arr {
		cnt[ch] = append(cnt[ch], i)
	}
	queue := []int{0}

	visit := make([]bool, n)
	visit[0] = true
	step := 0

	for len(queue) > 0 {
		lev := make([]int, 0)
		for _, no := range queue {
			if no == n-1 {
				return step
			}
			visit[no] = true
			if no-1 >= 0 && !visit[no-1] {
				visit[no-1] = true
				lev = append(lev, no-1)
			}
			if no+1 < n && !visit[no+1] {
				visit[no+1] = true
				lev = append(lev, no+1)
			}
			// 等值
			for _, nex := range cnt[arr[no]] {
				if !visit[nex] {
					visit[nex] = true
					lev = append(lev, nex)
				}
			}
			// 不加这一行，在极端情况下会超时，因为上面的：for _, nex := range cnt[arr[no]]
			delete(cnt, arr[no])
		}
		if len(lev) > 0 {
			step++
		}
		queue = lev
	}
	return -1
}
