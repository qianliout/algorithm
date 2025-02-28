package main

func main() {
}

func numIslands(grid [][]byte) int {
	return 0
}

type UnionFind struct {
	Parent []int
	Rank   []int
	Count  int
}

func NewUF(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i]++
	}
	uf := &UnionFind{Parent: parent, Rank: rank, Count: n}
	return uf
}

func (u *UnionFind) Union(a, b int) {
	x, y := u.Find(a), u.Find(b)
	if x != y {
		u.Count--
		if u.Rank[x] < u.Rank[y] {
			u.Parent[x] = y
		} else if u.Rank[x] > u.Rank[y] {
			u.Parent[y] = x
		} else {
			u.Rank[x]++
			u.Parent[y] = x
		}
	}
}

func (u *UnionFind) Find(x int) int {
	if u.Parent[x] != x {
		u.Parent[x] = u.Find(u.Parent[x])
	}
	return u.Rank[x]
}
func (u *UnionFind) Connect(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
