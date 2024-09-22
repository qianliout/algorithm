package segtree

/*
懒惰标记的基本实现步骤：
定义 Lazy 结构：用于存储当前节点需要的操作（如加法或乘法）及其值。
更新时检查 Lazy：在执行更新时，先检查当前节点的懒惰标记。如果存在标记，先处理它，然后再进行更新。
推送 Lazy 到子节点：在更新父节点时，如果子节点需要更新，将懒惰标记传递给子节点。
查询时处理 Lazy：在查询时，如果遇到有懒惰标记的节点，先更新节点值，然后再返回查询结果。
*/

type Lazy struct {
	NeedAdd bool
	Add     int
	NeedMul bool
	Mul     int
}

type SegTree struct {
	Data []int
	Node []int
	Lazy []Lazy
	Mod  int
	N    int
}

func NewSegTree(data []int, mod int) *SegTree {
	n := len(data)
	node := make([]int, 4*n)
	lazy := make([]Lazy, 4*n)
	return &SegTree{
		Data: data,
		Node: node,
		Lazy: lazy,
		N:    n,
		Mod:  mod,
	}
}

func (s *SegTree) Build() {
	s.build(1, 0, s.N-1)
}

func (s *SegTree) build(no, start, end int) {
	if start == end {
		s.Node[no] = s.Data[start]
		return
	}
	mid := start + (end-start)/2
	s.build(2*no, start, mid)
	s.build(2*no+1, mid+1, end)
	s.Node[no] = s.Node[2*no+1] + s.Node[2*no+2]
}

func (s *SegTree) Update(idx int, value int) {
	s.update(1, 0, s.N-1, idx, value)
}

func (s *SegTree) RangeAdd(left, right, value int) {
	s.updateRangeAdd(0, 0, s.N-1, left, right, value)
}
func (s *SegTree) RangeMul(left, right, value int) {
	s.updateRangeMul(0, 0, s.N-1, left, right, value)
}

func (s *SegTree) updateRangeMul(no, start, end, left, right, value int) {
	if s.Lazy[no].NeedMul {
		s.applyLazyMul(no, start, end)
	}
	if s.Lazy[no].NeedAdd {
		s.applyLazyAdd(no, start, end)
	}

	if start > right || end < left {
		return
	}
	if start >= left && end <= right {
		s.Node[no] *= value
		if start != end {
			s.Lazy[2*no+1].NeedMul = true
			s.Lazy[2*no+1].Mul *= value
			s.Lazy[2*no+2].NeedMul = true
			s.Lazy[2*no+2].Mul *= value
		}
		return
	}

	mid := start + (end-start)/2
	s.updateRangeMul(2*no+1, start, mid, left, right, value)
	s.updateRangeMul(2*no+2, mid+1, end, left, right, value)
	s.Node[no] = s.Node[2*no+1] + s.Node[2*no+2]
}

func (s *SegTree) applyLazyAdd(no, start, end int) {
	s.Node[no] += (end - start + 1) * s.Lazy[no].Add
	if start != end {
		s.Lazy[2*no+1].NeedAdd = true
		s.Lazy[2*no+1].Add += s.Lazy[no].Add
		s.Lazy[2*no+2].NeedAdd = true
		s.Lazy[2*no+2].Add += s.Lazy[no].Add
	}
	s.Lazy[no].NeedAdd = false
}

func (s *SegTree) Query(left, right int) int {
	return s.query(0, 0, s.N-1, left, right)
}

func (s *SegTree) query(no, start, end, left, right int) int {
	if s.Lazy[no].NeedMul {
		s.applyLazyMul(no, start, end)
	}
	if s.Lazy[no].NeedAdd {
		s.applyLazyAdd(no, start, end)
	}

	if start > right || end < left {
		return 0
	}
	if start >= left && end <= right {
		return s.Node[no]
	}

	mid := start + (end-start)/2
	le := s.query(2*no+1, start, mid, left, right)
	ri := s.query(2*no+2, mid+1, end, left, right)
	return le + ri
}

func (s *SegTree) update(no int, start, end int, idx int, value int) {
	if start == end {
		s.Data[idx] = value
		s.Node[no] = value
		return
	}
	mid := start + (end-start)/2
	if start <= idx && idx <= mid {
		// 左边
		s.update(2*no+1, start, mid, idx, value)
	} else {
		// 右边
		s.update(2*no+2, mid+1, end, idx, value)
	}
	// 更新上级
	s.Node[no] = (s.Node[2*no+1] + s.Node[2*no+2]) % s.Mod
}

func (s *SegTree) updateRangeAdd(no, start, end, left, right, value int) {
	if s.Lazy[no].NeedMul {
		s.applyLazyMul(no, start, end)
	}
	if s.Lazy[no].NeedAdd {
		s.applyLazyAdd(no, start, end)
	}

	if start > right || end < left {
		return
	}
	if start >= left && end <= right {
		s.Node[no] += (end - start + 1) * value
		if start != end {
			s.Lazy[2*no+1].NeedAdd = true
			s.Lazy[2*no+1].Add += value
			s.Lazy[2*no+2].NeedAdd = true
			s.Lazy[2*no+2].Add += value
		}
		return
	}

	mid := start + (end-start)/2
	s.updateRangeAdd(2*no+1, start, mid, left, right, value)
	s.updateRangeAdd(2*no+2, mid+1, end, left, right, value)
	s.Node[no] = (s.Node[2*no+1] + s.Node[2*no+2]) % s.Mod
}
func (s *SegTree) applyLazyMul(no, start, end int) {
	s.Node[no] = (s.Node[no] * s.Lazy[no].Mul) % s.Mod
	if start != end {
		s.Lazy[2*no+1].NeedMul = true
		s.Lazy[2*no+1].Mul *= s.Lazy[no].Mul
		s.Lazy[2*no+1].Mul %= s.Mod
		s.Lazy[2*no+2].NeedMul = true
		s.Lazy[2*no+2].Mul *= s.Lazy[no].Mul
		s.Lazy[2*no+2].Mul %= s.Mod
	}
	s.Lazy[no].NeedMul = false
}
