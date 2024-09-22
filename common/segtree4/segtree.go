package segtree

import (
	"math"
)

type SegTree struct {
	Data []int
	Node []int
	Lazy []Lazy
	N    int
	MOD  int
}

type Lazy struct {
	AddValue int
	MulValue int
}

func NewSegTree(data []int) *SegTree {
	n := len(data)
	tr := &SegTree{
		Data: data,
		Node: make([]int, 4*len(data)),
		Lazy: make([]Lazy, 4*len(data)),
		N:    n,
		MOD:  int(math.Pow10(9) + 7),
	}
	tr.build(0, 0, n-1)
	return tr
}

func (s *SegTree) build(no int, l, r int) {
	if l > r {
		return
	}
	if l == r {
		s.Node[no] = s.Data[l]
		return
	}

	mid := l + (r-l)/2
	s.build(no*2+1, l, mid)
	s.build(no*2+2, mid+1, r)
	s.Node[no] = s.Node[no*2+1] + s.Node[no*2+2]
}

func (s *SegTree) Query(start, end int) int {
	return s.query(0, 0, s.N-1, start, end)
}

func (s *SegTree) Update(start, end int, up Lazy) {
	s.update(0, 0, s.N-1, start, end, up)
}

// 要着重理解这几个参数的意义
// o:本次查询的区间起点，也是递归的入口：树状数组的根节点
// l,r,是次所查询的树状数组的左右边界，在递归过程中，会不断缩小查询范围
// start,end 是本次要查询的原数据的左右范围，闭区间
func (s *SegTree) query(no, l, r, start, end int) int {
	if l > end || r < start {
		return 0
	}

	if start <= l && end >= r {
		s.do(no, l, r, s.Lazy[no])
		return s.Node[no]
	}

	mid := l + (r-l)/2
	left, right := 0, 0
	if l <= mid {
		left = s.query(2*no+1, l, mid, start, end)
	}
	if r >= mid+1 {
		right = s.query(2*no+2, mid+1, r, start, end)
	}
	return (left + right) % s.MOD
}

func (s *SegTree) update(no, l, r int, start, end int, p Lazy) {
	if l > end || r < start {
		return
	}
	// 要更新的区间[start:end+1]已经完全包含[l:r+1]，就不用递归了,注意这里都是闭区间
	// 比如要更新 [1:7],此时的区间是[2:6],此时就不用递归，只把需要更新的操作记录到 lazy 里
	if start <= l && r >= end {
		s.do(no, l, r, p)
		return
	}
	// 走到这一步说明要[l:r]区间不完全在[start:end]的区间中
	// 比如: 要更新 [3:7],此时的区间是[1:4]或[4:8],这情况下就需要递归了
	// 第一步，把父节记录的 lazy 信息向子节点传递
	// 先做乘法，再做加法
	if s.Lazy[no].MulValue != 1 {
		s.Lazy[no*2+1].MulValue *= s.Lazy[no].MulValue
		s.Lazy[no*2+2].MulValue *= s.Lazy[no].MulValue
	}

	if s.Lazy[no].AddValue != 0 {
		s.Lazy[no*2+1].AddValue += s.Lazy[no].AddValue
		s.Lazy[no*2+2].AddValue += s.Lazy[no].AddValue
	}

	// 第二步:递归更新左右子节点
	// 左右两个节点
	mid := l + (r-l)/2
	// 有交集了，就需要继续递归了
	// 上面判断的 l>end || r<start 的情况，这里其实可以省略判断
	if l <= mid {
		s.update(no*2+1, l, mid, start, end, p)
	}
	if r >= mid+1 {
		s.update(no*2+2, mid+1, r, start, end, p)
	}
	// 维护操作：
	s.Node[no] = s.Node[2*no+1] + s.Node[2*no+2]
}

func (s *SegTree) do(no int, l, r int, p Lazy) {
	if l > r {
		return
	}
	sz := r - l + 1
	// 乘法暂时不懂
	if p.MulValue != 1 {
		s.Node[no] = s.Node[no] * p.MulValue % s.MOD
		s.Lazy[no].AddValue = s.Lazy[no].AddValue * p.MulValue % s.MOD
		s.Lazy[no].MulValue = s.Lazy[no].MulValue * p.MulValue % s.MOD
	}

	if p.AddValue != 0 {
		s.Node[no] = (s.Node[no] + sz*p.AddValue) % s.MOD
		s.Lazy[no].AddValue += p.AddValue
	}

	if l == r {
		s.Data[l] = s.Node[l]
	}
}

func (s *SegTree) spread(no int, l, r int, lazy Lazy) {
	if lazy.AddValue == 0 && lazy.MulValue == 1 {
		return
	}
}
