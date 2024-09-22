package segtree

type Update struct {
	Add int
	Mul int
}

type SegTree struct {
	Data []int
	Node []int
	Lazy []Update
	N    int
}

func NewSegTree(data []int) *SegTree {
	n := len(data)
	st := &SegTree{
		Data: data,
		Node: make([]int, 4*n),
		Lazy: make([]Update, 4*n),
		N:    len(data),
	}

	return st
}

func (s *SegTree) Build() {
	s.build(1, 1, s.N)
}
func (s *SegTree) build(rootId int, l, r int) {
	if l > r {
		return
	}
	if l == r {
		s.Node[rootId] = s.Data[l-1]
		return
	}
	mid := l + (r-l)/2
	s.build(2*rootId, l, mid)
	s.build(2*rootId+1, mid+1, r)
	s.PushUp(rootId)
}
func (s *SegTree) PushUp(rootId int) {
	s.Node[rootId] = s.Node[2*rootId] + s.Node[2*rootId+1]
}

func (s *SegTree) Update(start, end int, up Update) {
	s.update(1, 1, s.N, start, end, up)
}

func (s *SegTree) update(rootId, l, r int, start, end int, up Update) {
	// 不合法
	if l > r || start > end {
		return
	}
	// 无交集
	if end < l || start > r {
		return
	}
	// 全部包含
	if start <= l && r <= end {
		// 更新数字和，向上保持正确,也就是说 root节点的值是正确的,lazy[root]的值是还没有向下调整的值
		s.Node[rootId] += (end - start + 1) * up.Add
		// 增加 add 标记，表示本区间的 Sum 正确，子区间的 Sum 仍需要根据 add 的值来调整
		s.Lazy[rootId].Add += up.Add
		return
	}
	// 部分包含
	mid := l + (r-l)/2
	// 把lazy[root]的标计下推
	s.PushDown(rootId, mid-l+1, r-(mid+1)+1)
	s.update(rootId*2, l, mid, start, end, up)
	s.update(rootId*2+1, mid+1, r, start, end, up)
	s.PushUp(rootId)

	if l == r {
		s.Data[l-1] = s.Node[l]
	}
}

func (s *SegTree) PushDown(rootId int, leN, riN int) {
	if s.Lazy[rootId].Add == 0 {
		return
	}
	// 下推标计
	s.Lazy[2*rootId].Add += s.Lazy[rootId].Add
	s.Lazy[2*rootId+1].Add += s.Lazy[rootId].Add

	// 保持左右两节点值正确
	s.Node[2*rootId] += leN * s.Lazy[rootId].Add
	s.Node[2*rootId+1] += riN * s.Lazy[rootId].Add

	// 清除 root 的 标计
	s.Lazy[rootId] = Update{}
}
