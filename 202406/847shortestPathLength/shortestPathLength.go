package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestPathLength([][]int{{1, 2, 3}, {0}, {0}, {0}}))               // 4
	fmt.Println(shortestPathLength([][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}})) // 4
}

// 这是一个等权无向图，题目要我们求从「一个点都没访问过」到「所有点都被访问」的最短路径。
func shortestPathLength(graph [][]int) int {
	n := len(graph)          // 表示有 n 个点
	vis := make([][]bool, n) // 每个点，每个状态的记录
	for i := range vis {
		vis[i] = make([]bool, 1<<n)
	}
	queue := make([]pair, 0)
	// 把所有点都加入，意思是尝试所有的点做为开始点
	for i := 0; i < n; i++ {
		ma := 1 << i // 表示第i个数已使用
		queue = append(queue, pair{idx: i, mask: ma, cost: 0})
		vis[i][ma] = true
	}
	for len(queue) > 0 {
		fir := queue[0]
		// bfs 最先达到的一定是最短的距离
		if fir.mask == 1<<n-1 {
			return fir.cost
		}
		queue = queue[1:]

		nex := graph[fir.idx]

		for _, d := range nex {
			nextMask := fir.mask | (1 << d)
			if !vis[d][nextMask] { // 说明这个点的这个状态已经计算过了，再去计算就重复了
				vis[d][nextMask] = true
				queue = append(queue, pair{idx: d, mask: nextMask, cost: fir.cost + 1})
			}
		}
	}
	return 0
}

type pair struct {
	idx  int // 表示点的位置
	mask int // 表示当前的状态
	cost int // 表示走到这个点的步数
}
