package main

import (
	"fmt"

	. "outback/algorithm/common/unionfind"
)

func main() {
	fmt.Println(findCircleNum([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}

func findCircleNum(conn [][]int) int {
	n := len(conn)
	// uf := NewSizeUnionFind(n)
	uf := NewRankUnionFind(n)
	for x, ch := range conn {
		for y, v := range ch {
			if v == 1 {
				uf.Union(x, y)
			}
		}
	}
	return int(uf.Count)
}

/*
有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。
省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。
返回矩阵中 省份 的数
*/
