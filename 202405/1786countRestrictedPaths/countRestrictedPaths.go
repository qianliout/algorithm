package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	fmt.Println(countRestrictedPaths(5, [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}))
	fmt.Println(countRestrictedPaths(9, [][]int{{6, 2, 35129}, {3, 4, 99499}, {2, 7, 43547}, {8, 1, 78671}, {2, 1, 66308}, {9, 6, 33462}, {5, 1, 48249}, {2, 3, 44414}, {6, 7, 44602}, {1, 7, 14931}, {8, 9, 38171}, {4, 5, 30827}, {3, 9, 79166}, {4, 8, 93731}, {5, 9, 64068}, {7, 5, 17741}, {6, 3, 76017}, {9, 4, 72244}}))
}

func countRestrictedPaths(n int, edges [][]int) int {
	base := int(math.Pow(10, 9)) + 7

	dis := make([]int, n+1)
	g := GenG2(edges, n+1)
	for i := 1; i <= n; i++ {
		dis[i] = GetMin2(g, n+1, i, n)
	}
	dis2 := make([]pair, n)
	for i := range dis[1:] {
		dis2[i] = pair{i + 1, dis[i+1]}
	}
	// 根据点距离从小到大排序
	sort.Slice(dis2, func(i, j int) bool { return dis2[i].d < dis2[j].d })
	// 开始 dp
	// dp[i] 为从点 i 到第点 n 的受限路径数量
	dp := make([]int, n+1)
	dp[n] = 1
	// 不失一般性，当我们要求 dp[i] 的时候，其实找的所有满足「与点 i 相连，且最短路比点 i 要小的点 j」，
	// 符合条件的点 j 有很多个，将所有的 dp[j] 累加即是 dp[i]。
	// dp[i] = dp[i - 1] if dis[i-1] < dis[i]
	// 这里有个条件 dp[i] 由dp[j]转化，所以 i 和 j一定是要相连的，如果用邻接矩阵的话，要判断是否相连
	// 这样写会超时,要优化
	for i := 0; i < n; i++ {
		node := dis2[i]
		for _, d := range g[node.x] {
			if node.d > dis[d.x] {
				dp[node.x] += dp[d.x]
			}
		}
	}

	return dp[1] % base
}

func GetMin2(g [][]pair, n, start, end int) int {
	// 下标统一都做减一操作
	// 注意点1：这里最好不定义成 math.MaxInt32,因为下面有加法，可能会有溢出
	inf := math.MaxInt
	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	hp := make(PriorityQueue, 0)
	heap.Push(&hp, &IntItem{Value: start, Priority: 0})

	for hp.Len() > 0 {
		pop := heap.Pop(&hp).(*IntItem)
		x := pop.Value
		dx := -pop.Priority
		// 这里最好是>= 因为下面更新 dis[y] 时没有做判断，是直接加的，可能会比 inf 大
		// 即使不做判断，也不会有值超过 inf 因为是用的 min()操作
		// 说是不可达了
		if dis[x] < dx {
			continue
		}
		// 对于第一次的循环，x 一定是起始点,所以在这里更新 done 数组
		for _, nex := range g[x] {
			// 这里可以做一步判断，判断是否 >= inf,也可以不判断，因为上面 >= inf 就都认为不可达
			// 也可以不判断
			nd := dx + nex.d
			if dis[nex.x] > nd {
				dis[nex.x] = nd
				heap.Push(&hp, &IntItem{Value: nex.x, Priority: -nd})
			}
		}
	}
	if dis[end] >= inf {
		return -1
	}
	return dis[end]
}

type pair struct {
	x, d int
}

func GenG2(edges [][]int, n int) [][]pair {
	g := make([][]pair, n)
	// 用邻接矩阵会超时,尝试邻接表
	for _, ch := range edges {
		x, y, z := ch[0], ch[1], ch[2]
		g[x] = append(g[x], pair{x: y, d: z})
		g[y] = append(g[y], pair{x: x, d: z})
	}
	return g
}
