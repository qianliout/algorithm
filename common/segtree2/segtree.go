package segtree

type SegTree struct {
	Data []int
	Node []int
	Lazy []Update
	N    int
}
type Update struct {
	Add int
}

func NewSegTree(data []int) *SegTree {
	n := len(data)
	tr := &SegTree{
		Data: data,
		Node: make([]int, 4*n),
		Lazy: make([]Update, 4*n),
		N:    len(data),
	}
	tr.Build(1, 1, tr.N)
	return tr
}

func (s *SegTree) Build(o int, l, r int) {
	if l > r {
		return
	}
	if l == r {
		s.Node[o] = s.Data[l-1]
		return
	}
	mid := l + (r-l)/2
	s.Build(o*2, l, mid)
	s.Build(o*2+1, mid+1, r)
	s.Node[o] = s.Node[o*2+1] + s.Node[o*2+2]
}

func (s *SegTree) Update(start, end int, update Update) {
	s.update(1, 1, s.N, start, end, update)
}

func (s *SegTree) update(no int, l, r int, start, end int, value Update) {
	if l > r || start > end {
		return
	}
	if start <= l && r <= end {
		s.do(no)
		return
	}

}

func (s *SegTree) Query(l, r int) int {
	return s.query(0, 0, s.N-1, l, r)
}

func (s *SegTree) query(o, l, r, start, end int) int {
	if end < l || r < start {
		return 0
	}
	// if l <= start && r >= end {
	// 	return s.Node[o]
	// }
	if start <= l && end >= r {
		return s.Node[o]
	}

	mid := l + (r-l)/2
	left := s.query(o*2+1, l, mid, start, end)
	right := s.query(o*2+2, mid+1, r, start, end)
	return left + right
}

// 更新左右子树
func (s *SegTree) do(no int) {

}
