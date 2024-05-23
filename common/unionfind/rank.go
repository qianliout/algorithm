package unionfind

type RankUnionFind struct {
	Parent []int // 可以理解成下标i的最终父节点就是 Parent[i]
	Rank   []int
	Count  int
}

func NewRankUnionFind(totalNodes int) *RankUnionFind {
	p := make([]int, totalNodes)
	r := make([]int, totalNodes)
	for i := 0; i < totalNodes; i++ {
		p[i] = i
		r[i] = 1
	}
	return &RankUnionFind{Parent: p, Rank: r, Count: totalNodes}
}

func (u *RankUnionFind) Find(x int) int {
	if u.Parent[x] != x {
		// 路径压缩
		u.Parent[x] = u.Find(u.Parent[x])
	}
	return u.Parent[x]
}

func (u *RankUnionFind) Union(x, y int) {
	xRoot := u.Find(x)
	yRoot := u.Find(y)

	// 把低的rank赋值给高的node，这里有些不理解，
	if xRoot != yRoot {
		u.Count--
		if u.Rank[xRoot] > u.Rank[yRoot] {
			u.Parent[yRoot] = xRoot
		} else if u.Rank[xRoot] < u.Rank[yRoot] {
			u.Parent[xRoot] = yRoot
		} else {
			u.Parent[yRoot] = xRoot
			u.Rank[xRoot]++
		}
	}
}

func (u *RankUnionFind) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
