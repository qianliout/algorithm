package dijkstra

import (
	"math"
)

// https://leetcode.cn/problems/reachable-nodes-in-subdivided-graph/description/

type neighbor struct {
	x, d int
}

// Dijkstra算法实现（不使用堆，使用双重循环）
// 时间复杂度: O(V²) - V为顶点数，每次需要线性搜索最小距离顶点
// 空间复杂度: O(V) - 需要距离数组和访问标记数组
// 工作原理: 每次选择未访问顶点中距离最小的顶点，更新其邻居的距离
func dijkstra(g [][]neighbor, start int, end int) []int {
	inf := math.MaxInt / 10
	// 距离数组，存储从起点到每个顶点的最短距离
	dis := make([]int, end+1)
	// 访问标记数组，标记顶点是否已确定最短距离
	visited := make([]bool, end+1)

	// 初始化距离数组，所有距离设为无穷大
	for i := range dis {
		dis[i] = inf
	}
	// 起点到自身的距离为0
	dis[start] = 0

	// 主循环：每次选择一个未访问的顶点
	for {
		// 找到未访问顶点中距离最小的顶点
		minDis := inf
		minVertex := -1

		// 线性搜索最小距离顶点
		for i := 0; i <= end; i++ {
			if !visited[i] && dis[i] < minDis {
				minDis = dis[i]
				minVertex = i
			}
		}

		// 如果没有找到可访问的顶点，说明所有可达顶点都已处理完
		if minVertex == -1 {
			break
		}

		// 标记当前顶点为已访问
		visited[minVertex] = true

		// 更新当前顶点的所有邻居的距离
		for _, next := range g[minVertex] {
			// 计算通过当前顶点到达邻居的新距离
			newDis := dis[minVertex] + next.d
			// 如果新距离更短，则更新距离
			if dis[next.x] > newDis {
				dis[next.x] = newDis
			}
		}
	}

	return dis
}

// 注意：原来的堆相关代码已移除，因为不再需要
// 这个实现更简单直观，但时间复杂度较高
// 适用于顶点数较少或对性能要求不高的场景
