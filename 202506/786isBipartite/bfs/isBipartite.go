package main

import (
	. "outback/algorithm/common/unionfind"
)

func main() {

}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	uf := NewSizeUnionFind(n)

	for i, ch := range graph {
		// 容易忘记
		// if len(ch) == 0 {
		// 	continue
		// }
		// fir := ch[0]
		for _, w := range ch {
			if uf.IsConnected(i, w) {
				return false
			}
			uf.Union(ch[0], w)
		}
	}
	return true
}

/*
我们知道如果是二分图的话，那么图中每个顶点的所有邻接点都应该属于同一集合，且不与顶点处于同一集合。
因此我们可以使用并查集来解决这个问题，我们遍历图中每个顶点，将当前顶点的所有邻接点进行合并，
并判断这些邻接点中是否存在某一邻接点已经和当前顶点处于同一个集合中了，若是，则说明不是二分图。
*/
