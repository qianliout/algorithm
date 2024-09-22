package segtree

import "math"

type SegTree struct {
	Data []int
	Node []int
	Lazy []Lazy
	N    int
	MOD  int
}

type Lazy struct {
	AddValue int
	NeedAdd  bool
	NeedMul  bool
	MulValue int
}

func NewSegTree(data []int) *SegTree {
	tr := &SegTree{
		Data: data,
		Node: make([]int, 4*len(data)),
		Lazy: make([]Lazy, 4*len(data)),
		N:    len(data),
		MOD:  int(math.Pow10(9) + 7),
	}
	tr.Build(0, 0, tr.N-1)
	return tr
}

func (s *SegTree) Build(no int, l, r int) {
	if l == r {
		s.Node[no] = s.Data[l]
		return
	}
	mid := l + (r-l)/2
	s.Build(no*2+1, l, mid)
	s.Build(no*2+2, mid+1, r)
	s.Node[no] = s.Node[no*2+1] + s.Node[no*2+2]
}

// 直接单值更新时，可以不用 lazy 做法,只有区间更新
func (s *SegTree) Update(idx int, value int) {
	s.update(0, 0, s.N-1, idx, value)
}

func (s *SegTree) update(no int, l, r int, idx, value int) {
	if l == r {
		s.Data[idx] = value
		s.Node[no] = value
		return
	}
	mid := l + (r-l)/2
	if idx <= mid {
		s.update(2*no+1, l, mid, idx, value)
	} else {
		s.update(2*no+2, mid+1, r, idx, value)
	}
	s.Node[no] = s.Node[no*2+1] + s.Node[no*2+2]
}

func (s *SegTree) Query(l, r int) int {
	return s.query(0, 0, s.N-1, l, r)
}

// 要着重理解这几个参数的意义
// o:本次查询的区间起点，也是递归的入口：树状数组的根节点
// l,r,是次所查询的树状数组的左右边界，在递归过程中，会不断缩小查询范围
// start,end 是本次要查询的原数据的左右范围，闭区间
func (s *SegTree) query(no, l, r, start, end int) int {
	// 本次所查询的树状数组的范围没有与查询范围有交集
	if end < l || r < start {
		return 0
	}
	if s.Lazy[no].NeedAdd || s.Lazy[no].NeedMul {
		s.Lazy[no*2+1].AddValue += s.Lazy[no].AddValue
		s.Lazy[no*2+2].AddValue += s.Lazy[no].AddValue
		s.Lazy[no].AddValue = 0
		s.Lazy[no].NeedAdd = false

		s.Lazy[no*2+1].MulValue *= s.Lazy[no].MulValue
		s.Lazy[no*2+2].MulValue *= s.Lazy[no].MulValue
		s.Lazy[no].NeedMul = false
		s.Lazy[no].MulValue = 1
	}

	// 树装数组的范围已经完全在查询范围内了，就没有必要再递归下次了
	if start <= l && end >= r {
		return s.Node[no]
	}

	mid := l + (r-l)/2
	left := s.query(no*2+1, l, mid, start, end)
	right := s.query(no*2+2, mid+1, r, start, end)
	return (left + right) % s.MOD
}

func (s *SegTree) RangeAdd(start, end int, add int) {
	if start > end {
		return
	}
	s.rangeAdd(0, 0, s.N-1, start, end, add)
}

func (s *SegTree) rangeAdd(no, l, r int, start, end, value int) {
	// 没有破坏区间
	if start <= l && end >= r {
		s.Lazy[no].AddValue += value
		s.Lazy[no].NeedAdd = true
		return
	}

	// 破坏了区间，那么此时就需要把 lazy 值传递下去
	if s.Lazy[no].NeedAdd {
		s.Lazy[no*2+1].AddValue += s.Lazy[no].AddValue
		s.Lazy[no*2+2].AddValue += s.Lazy[no].AddValue
		s.Lazy[no].AddValue = 0
		s.Lazy[no].NeedAdd = false

		s.Lazy[no*2+1].MulValue *= s.Lazy[no].MulValue
		s.Lazy[no*2+2].MulValue *= s.Lazy[no].MulValue
		s.Lazy[no].NeedMul = false
		s.Lazy[no].MulValue = 1
	}
	mid := l + (r-l)/2

	// 有交集了就要更新
	if start <= mid {
		s.rangeAdd(no*2+1, l, mid, start, end, value)
	}
	if end >= mid+1 {
		s.rangeAdd(no*2+2, mid+1, r, start, end, value)
	}
	// 维护根节点
	s.Node[no] = (s.Node[no*2+1] + s.Node[no*2+2]) % s.MOD
}

func (s *SegTree) RangeMul(start, end int, value int) {
	if start > end {
		return
	}
	s.rangeAdd(0, 0, s.N-1, start, end, value)
}

func (s *SegTree) rangeMul(no, l, r int, start, end, value int) {
	// 没有破坏区间
	if start <= l && end >= r {
		s.Lazy[no].MulValue *= value
		s.Lazy[no].NeedMul = true
		return
	}
	// 破坏了区间，那么此时就需要把 lazy 值传递下去
	if s.Lazy[no].NeedMul {
		s.Lazy[no*2+1].AddValue += s.Lazy[no].AddValue
		s.Lazy[no*2+2].AddValue += s.Lazy[no].AddValue
		s.Lazy[no].AddValue = 0
		s.Lazy[no].NeedAdd = false

		s.Lazy[no*2+1].MulValue *= s.Lazy[no].MulValue
		s.Lazy[no*2+2].MulValue *= s.Lazy[no].MulValue
		s.Lazy[no].NeedMul = false
		s.Lazy[no].MulValue = 1
	}
	mid := l + (r-l)/2

	// 有交集了就要更新
	if start <= mid {
		s.rangeMul(no*2+1, l, mid, start, end, value)
	}
	if end >= mid+1 {
		s.rangeAdd(no*2+2, mid+1, r, start, end, value)
	}
	// 维护根节点
	s.Node[no] = (s.Node[no*2+1] + s.Node[no*2+2]) % s.MOD
}
