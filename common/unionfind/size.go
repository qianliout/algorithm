package unionfind

type SizeUnionFind struct {
	Parent []int // 可以理解成下标i的最终父节点就是 Parent[i]
	Size   []int
	Count  int
}

func NewSizeUnionFind(totalNodes int) *SizeUnionFind {
	p := make([]int, totalNodes)
	s := make([]int, totalNodes)
	for i := 0; i < totalNodes; i++ {
		p[i] = i
		s[i] = 1
	}
	return &SizeUnionFind{Parent: p, Size: s, Count: totalNodes}
}

func (u *SizeUnionFind) Find(x int) int {
	if u.Parent[x] != x {
		u.Parent[x] = u.Find(u.Parent[x]) // 路径压缩
	}
	return u.Parent[x]
}

func (u *SizeUnionFind) Union(x, y int) {
	xRoot := u.Find(x)
	yRoot := u.Find(y)
	if xRoot != yRoot {
		u.Count--
		if u.Size[xRoot] > u.Size[yRoot] {
			u.Parent[yRoot] = xRoot
			u.Size[xRoot] += u.Size[yRoot]
		} else {
			u.Parent[xRoot] = yRoot
			u.Size[yRoot] += u.Size[xRoot]
		}
	}
}

func (u *SizeUnionFind) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
