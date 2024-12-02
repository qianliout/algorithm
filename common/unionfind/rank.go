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
		// 这是为了进行路径压缩（path compression），具体有以下好处：
		// 加速后续查找：通过将当前节点直接连接到查询得到的根节点上，减少了后续对该节点及其子节点进行查找时的层级深度，从而加快了下一次查找的速度。
		// 优化性能：在不增加额外操作复杂度的前提下，通过对路径上的每个节点进行修改，使得树形结构更加扁平化，进而提升了整个并查集结构的效率。
		u.Parent[x] = u.Find(u.Parent[x])

		// 当然也可以在查询过程中不修改数据
		// return u.Find(u.Parent[x])
	}
	return u.Parent[x]
}

func (u *RankUnionFind) Union(x, y int) {
	xRoot := u.Find(x)
	yRoot := u.Find(y)

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
