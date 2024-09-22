package segtree

type SegTree struct {
	Data []int
	Node []int
	N    int
}

func NewSegTree(data []int) *SegTree {
	tr := &SegTree{
		Data: data,
		Node: make([]int, 4*len(data)),
		N:    len(data),
	}
	tr.Build(0, 0, tr.N-1)
	return tr
}

func (s *SegTree) Build(o int, l, r int) {
	if l == r {
		s.Node[o] = s.Data[l]
		return
	}
	mid := l + (r-l)/2
	s.Build(o*2+1, l, mid)
	s.Build(o*2+2, mid+1, r)
	s.Node[o] = s.Node[o*2+1] + s.Node[o*2+2]
}

func (s *SegTree) Update(idx int, value int) {
	s.update(0, 0, s.N-1, idx, value)
}

func (s *SegTree) update(o int, l, r int, idx, value int) {
	if l == r {
		s.Data[idx] = value
		s.Node[o] = value
		return
	}
	mid := l + (r-l)/2
	if idx <= mid {
		s.update(2*o+1, l, mid, idx, value)
	} else {
		s.update(2*o+2, mid+1, r, idx, value)
	}
	s.Node[o] = s.Node[o*2+1] + s.Node[o*2+2]
}

func (s *SegTree) Query(l, r int) int {
	return s.query(0, 0, s.N-1, l, r)
}

// 要着重理解这几个参数的意义
// o:本次查询的区间起点，也是递归的入口：树状数组的根节点
// l,r,是次所查询的树状数组的左右边界，在递归过程中，会不断缩小查询范围
// start,end 是本次要查询的原数据的左右范围，闭区间
func (s *SegTree) query(o, l, r, start, end int) int {
	// 本次所查询的树状数组的范围没有与查询范围有交集
	if end < l || r < start {
		return 0
	}

	// 树装数组的范围已经完全在查询范围内了，就没有必要再递归下次了
	if start <= l && end >= r {
		return s.Node[o]
	}

	mid := l + (r-l)/2
	left := s.query(o*2+1, l, mid, start, end)
	right := s.query(o*2+2, mid+1, r, start, end)
	return left + right
}
