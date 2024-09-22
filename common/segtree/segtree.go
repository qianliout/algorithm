package segtree

/*
根节点node[0]
对于任务节点 i，他的父节点是(i-1)/2
左孩子节点：i*2+1
右孩子节点：i*2+2
*/

type SegTree struct {
	Data []int
	Node []int
	N    int
	Mod  int
}

func NewSegTree(data []int, mod int) *SegTree {
	n := len(data)
	node := make([]int, 4*n)
	return &SegTree{
		Data: data,
		Node: node,
		N:    n,
		Mod:  mod,
	}
}

func (s *SegTree) Build() {
	s.build(1, 1, s.N)
}

func (s *SegTree) build(no int, start, end int) {
	if start == end {
		s.Node[no] = s.Data[start]
		return
	}
	mid := start + (end-start)/2
	s.build(2*no+1, start, mid)
	s.build(2*no+2, mid, end)
	s.Node[no] = s.Node[2*no+1] + s.Node[2*no+2]
}

func (s *SegTree) Update(idx int, value int) {
	s.update(1, 1, s.N, idx, value)
}
func (s *SegTree) update(no int, start, end int, idx int, value int) {
	if start == end {
		s.Data[idx] = value
		s.Node[no] = value
		return
	}
	mid := start + (end-start)/2
	if start <= idx+1 && idx+1 <= mid {
		// 左边
		s.update(2*no, start, mid, idx, value)
	} else {
		// 右边
		s.update(2*no+1, mid+1, end, idx, value)
	}
	// 更新上级
	s.Node[no] = (s.Node[2*no] + s.Node[2*no]) % s.Mod
}

func (s *SegTree) UpdateAdd(idx int, value int) {
	s.updateAdd(1, 1, s.N, idx, value)
}
func (s *SegTree) updateAdd(no int, start, end int, idx int, value int) {
	if start == end {
		s.Data[idx] += value
		s.Node[no] += value
		return
	}
	mid := start + (end-start)/2
	if start <= idx && idx <= mid {
		// 左边
		s.updateAdd(2*no, start, mid, idx, value)
	} else {
		// 右边
		s.updateAdd(2*no+1, mid+1, end, idx, value)
	}
	// 更新上级
	s.Node[no] = (s.Node[2*no] + s.Node[2*no+1]) % s.Mod
}

func (s *SegTree) UpdateMul(idx int, value int) {
	s.updateMul(1, 0, s.N-1, idx, value)
}
func (s *SegTree) updateMul(no int, start, end int, idx int, value int) {
	if start == end {
		s.Data[idx] = (s.Data[idx] * value) % s.Mod
		s.Node[no] = (s.Node[idx] * value) % s.Mod

		return
	}
	mid := start + (end-start)/2
	if start <= idx && idx <= mid {
		// 左边
		s.updateMul(2*no, start, mid, idx, value)
	} else {
		// 右边
		s.updateMul(2*no+1, mid+1, end, idx, value)
	}
	// 更新上级
	s.Node[no] = (s.Node[2*no] + s.Node[2*no+1]) % s.Mod
}

func (s *SegTree) Query(left, right int) int {
	return s.query(1, 0, s.N-1, left, right)
}

func (s *SegTree) query(no int, start, end int, left, right int) int {
	// 没有交集
	if left > end || right < start {
		return 0
	}
	// 完全包含
	if start <= left && end >= right {
		return s.Node[no]
	}
	mid := start + (end-left)/2
	le := s.query(no*2, start, mid, left, right)
	ri := s.query(no*2+1, mid+1, end, left, right)
	return le + ri
}

func (s *SegTree) RangeAdd(l, r, value int) {
	for i := l; i <= r; i++ {
		s.UpdateAdd(i, value)
	}
}

func (s *SegTree) RangeMul(l, r, value int) {
	for i := l; i <= r; i++ {
		s.UpdateMul(i, value)
	}
}
