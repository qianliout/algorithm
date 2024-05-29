package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
}

func networkDelayTime(times [][]int, n int, k int) int {
	// 下标统一都做减一操作
	// 注意点1：这里最好不定义成 math.MaxInt32,因为下面有加法，可能会有溢出
	inf := math.MaxInt32 / 10
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, ch := range times {
		x, y, z := ch[0]-1, ch[1]-1, ch[2]
		g[x][y] = z
	}
	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[k-1] = 0
	done := make([]bool, n)
	for {
		x := -1
		for i, ok := range done {
			// 这里用 dis[i]<=dis[x] 还是用 dis[i]<dis[x]都能得到正确结果，还是还是推荐使用<=
			if !ok && (x < 0 || dis[i] <= dis[x]) {
				x = i
			}
		}

		if x < 0 { // 说明所有 n 个元素都更新的完了
			break
		}
		// 这里最好是>= 因为下面更新 dis[y] 时没有做判断，是直接加的，可能会比 inf 大
		// 即使不做判断，也不会有值超过 inf 因为是用的 min()操作
		// 说是不可达了
		if dis[x] >= inf {
			return -1
		}
		// 对于第一次的循环，x 一定是起始点,所以在这里更新 done 数组
		done[x] = true
		for y, d := range g[x] {
			// 这里可以做一步判断，判断是否 >= inf,也可以不判断，因为上面 >= inf 就都认为不可达
			// 也可以不判断
			if d >= inf {
				continue
			}
			dis[y] = min(dis[y], dis[x]+d)
		}
	}
	return slices.Max(dis)
}
