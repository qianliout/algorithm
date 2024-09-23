package segtree

import (
	"math"
	"math/bits"
)

type Todo struct {
	Add int
	Mul int
}

type SegTree struct {
	Data []int
	Node []*Node
	Lazy []*Todo
	N    int
	MOD  int
}
type Node struct {
	Left  int
	Right int
	Value int
	Todo  *Todo
}

func NewSegTree(data []int) *SegTree {
	n := len(data)
	st := &SegTree{
		Data: data,
		Node: make([]*Node, 2<<bits.Len(uint(n-1))),
		Lazy: make([]*Todo, 2<<bits.Len(uint(n-1))),
		N:    n,
		MOD:  int(math.Pow(10, 9) + 7),
	}

	st.Build()

	return st
}

func (s *SegTree) Build() {
	s.build(1, 1, s.N)
}
func (s *SegTree) build(rootId int, l, r int) {
	// 不合法
	if l > r {
		return
	}
	if s.Node[rootId] == nil {
		s.Node[rootId] = &Node{}
	}
	if s.Lazy[rootId] == nil {
		s.Lazy[rootId] = &Todo{Mul: 1}
	}
	s.Node[rootId].Left = l
	s.Node[rootId].Right = r
	s.Lazy[rootId] = &Todo{Mul: 1}

	if l == r {
		s.Node[rootId].Value = s.Data[l-1]
		return
	}
	mid := l + (r-l)/2
	s.build(2*rootId, l, mid)
	s.build(2*rootId+1, mid+1, r)
	s.PushUp(rootId)
}
func (s *SegTree) PushUp(rootId int) {
	s.Node[rootId].Value = (s.Node[rootId<<1].Value + s.Node[rootId<<1|1].Value) % s.MOD
}

func (s *SegTree) Update(start, end int, up *Todo) {
	s.update(1, start, end, up)
}

func (s *SegTree) update(rootId, start, end int, up *Todo) {
	// 不合法
	if start > end {
		return
	}
	// 无交集
	root := s.Node[rootId]
	l, r := root.Left, root.Right
	if end < l || start > r {
		return
	}
	// 全部包含
	if start <= l && r <= end {
		s.Do(rootId, up)
		return
	}
	s.PushDown(rootId)

	// 部分包含
	mid := l + (r-l)/2

	// 把lazy[root]的标计下推
	if l <= mid {
		s.update(rootId<<1, start, end, up)
		// s.update(root.Left, start, end, up)
	}
	if mid < r {
		s.update(rootId<<1+1, start, end, up)
		// s.update(root.Right, start, end, up)
	}
	s.PushUp(rootId)
}

func (s *SegTree) PushDown(rootId int) {
	lz := s.Lazy[rootId]
	if lz.Add == 0 && lz.Mul == 1 {
		return
	}

	s.Do(rootId<<1, lz)
	s.Do(rootId<<1|1, lz)
	s.Lazy[rootId] = &Todo{Mul: 1}
}

func (s *SegTree) Do(rootId int, up *Todo) {

	root := s.Node[rootId]

	lz := s.Lazy[rootId]
	sz := root.Right - root.Left + 1

	if up.Mul != 1 {
		root.Value = (root.Value * up.Mul) % s.MOD
		lz.Add = (lz.Add * up.Mul) % s.MOD
		lz.Mul = (lz.Mul * up.Mul) % s.MOD
	}
	if up.Add != 0 {
		root.Value = (root.Value + sz*up.Add) % s.MOD
		lz.Add = (lz.Add + up.Add) % s.MOD
	}

	s.Node[rootId] = root
	s.Lazy[rootId] = lz
}

func (s *SegTree) Query(start, end int) int {
	return s.query(1, start, end)
}

func (s *SegTree) query(rootId int, start, end int) int {
	// 不合法
	if start > end {
		return 0
	}
	// 无交集
	root := s.Node[rootId]
	l, r := root.Left, root.Right
	if end < l || start > r {
		return 0
	}
	// 全部包含
	if start <= l && r <= end {
		return s.Node[rootId].Value
	}
	// 部分包含
	// 把lazy[root]的标计下推
	s.PushDown(rootId)

	mid := l + (r-l)/2
	ans := 0
	if end <= mid {
		return s.query(rootId<<1, start, end)
	}
	if start >= mid+1 {
		return s.query(rootId<<1|1, start, end)
	}
	ans += s.query(rootId<<1, start, end)
	ans += s.query(rootId<<1+1, start, end)
	return ans % s.MOD
}
