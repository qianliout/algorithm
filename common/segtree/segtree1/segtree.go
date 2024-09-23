package segtree

/*

// 最简单的方式，只支持单点更新,节点从0开始

根节点node[0]
对于任务节点 i，他的父节点是(i-1)/2
左孩子节点：i*2+1
右孩子节点：i*2+2
*/

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

func (s *SegTree) query(o, l, r, start, end int) int {
	if end < l || r < start {
		return 0
	}
	if start <= l && end >= r {
		return s.Node[o]
	}

	mid := l + (r-l)/2
	left := s.query(o*2+1, l, mid, start, end)
	right := s.query(o*2+2, mid+1, r, start, end)
	return left + right
}
