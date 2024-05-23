package main

import (
	. "outback/algorithm/common/unionfind"
)

func main() {

}

/*
思路：

	利用并查集的思想，相同字母可以形成连通区域。
	从左上角顶点开始，同时向右和向下搜索，若字母相同则合并。
	合并时，若发现 x 和 y 的 parent 相同，即形成环。

备注：

	从左上角顶点开始，向右向下搜索即可，不需要考虑向左和向上。扩展：windows经典游戏扫雷，由于我们可以在屏幕随便点击某个方块，这时需要考虑多个方向。
	为加快求解时间，直接短路返回了（找到一条环就返回true）。若要输出所有的环，则需完全遍历一遍。

两个联通怎么就能说明有环呢？

	每次都是从左到右，从上到小，如果没有环，那么是不可能有共同的祖先的
	兼容从 A 点出发走，一条路从左到右，一条路从上到下，走到grid[i][j]时， 如果 grid[i][j] 和grid[i-1][j]是在一起了，那只能说明另一条路来过这里
*/
func containsCycle(grid [][]byte) bool {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false
	}
	m, n := len(grid), len(grid[0])
	uf := NewRankUnionFind(m*n + 1)
	for i := range grid {
		for j, ch := range grid[i] {
			if i > 0 && grid[i-1][j] == ch {
				// 两个相同，且有共同的祖先，就是有环了
				if uf.IsConnected(i*n+j, (i-1)*n+j) {
					return true
				}
				uf.Union(i*n+j, (i-1)*n+j)
			}
			if j > 0 && grid[i][j-1] == ch {
				// 两个相同，且有共同的祖先，就是有环了
				if uf.IsConnected(i*n+j, i*n+j-1) {
					return true
				}
				uf.Union(i*n+j, i*n+j-1)
			}
		}
	}
	return false
}
