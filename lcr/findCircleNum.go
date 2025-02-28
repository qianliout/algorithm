package main

import (
	"fmt"
)

func main() {
	fmt.Println(findCircleNum([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := NewUnionFind(n)
	for i := range isConnected {
		for j, v := range isConnected[i] {
			if v == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.Count
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{Parent: parent, Count: n}
}

type UnionFind struct {
	Parent []int
	Rank   []int
	Count  int
}

func (u *UnionFind) Union(a int, b int) {
	ar := u.FindParent(a)
	br := u.FindParent(b)
	u.Parent[ar] = br
	// 可以在写入数据时不做路径圧缩，到每次查询时做路径圧缩

	if ar != br {
		u.Count--
	}
}

func (u *UnionFind) Connect(a, b int) bool {
	return u.FindParent(a) == u.FindParent(b)
}

func (u *UnionFind) FindParent(a int) int {
	if u.Parent[a] != a {
		// 路径压缩
		// 这是为了进行路径压缩（path compression），具体有以下好处：
		// 加速后续查找：通过将当前节点直接连接到查询得到的根节点上，减少了后续对该节点及其子节点进行查找时的层级深度，从而加快了下一次查找的速度。
		// 优化性能：在不增加额外操作复杂度的前提下，通过对路径上的每个节点进行修改，使得树形结构更加扁平化，进而提升了整个并查集结构的效率。
		u.Parent[a] = u.FindParent(u.Parent[a])

		// 当然也可以在查询过程中不修改数据
		// return u.Find(u.Parent[x])
		u.Parent[a] = u.FindParent(u.Parent[a])
	}
	return u.Parent[a]
}
