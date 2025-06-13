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
	}
	return &RankUnionFind{Parent: p, Rank: r, Count: totalNodes}
}

func (u *RankUnionFind) Find(x int) int {
	if u.Parent[x] != x {
		u.Parent[x] = u.Find(u.Parent[x]) // 路径压缩
	}
	return u.Parent[x]
}

func (u *RankUnionFind) Union(x, y int) {
	xRoot := u.Find(x)
	yRoot := u.Find(y)

	if xRoot != yRoot {
		u.Count--
		// 在并查集的按秩合并（rank union）中，只有当两个集合的秩（rank）相等时，合并后根节点的秩才需要加一（即u.Rank[xRoot]++）。
		// 如果u.Rank[xRoot] < u.Rank[yRoot]或u.Rank[xRoot] > u.Rank[yRoot]，较小秩的树直接挂到较大秩的树下，不会改变较大树的高度，所以不需要更新rank。
		// 只有秩相等时，合并会导致树高加一，这时才需要更新rank
		if u.Rank[xRoot] < u.Rank[yRoot] {
			u.Parent[xRoot] = yRoot
		} else if u.Rank[xRoot] > u.Rank[yRoot] {
			u.Parent[yRoot] = xRoot
		} else {
			u.Parent[yRoot] = xRoot
			u.Rank[xRoot]++
		}
	}
}

func (u *RankUnionFind) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
