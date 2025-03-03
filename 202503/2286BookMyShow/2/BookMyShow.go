package main

import (
	"fmt"
)

func main() {
	b := Constructor(4, 5)
	fmt.Println(b.Scatter(6, 2))
	fmt.Println(b.Gather(6, 3))
	fmt.Println(b.Scatter(9, 1))
}

type BookMyShow struct {
	st   SegmentTree
	M, N int
}

func Constructor(n int, m int) BookMyShow {
	st := make(SegmentTree, n*4)
	st.build(1, 0, n-1)
	return BookMyShow{st: st, M: m, N: n}
}

func (s *BookMyShow) Gather(k int, maxRow int) []int {
	r := s.st.findFirst(1, maxRow, s.M-k)
	if r < 0 {
		return nil
	}
	c := s.st.querySum(1, r, r)
	s.st.update(1, r, k)
	return []int{r, c}
}

func (s *BookMyShow) Scatter(k int, maxRow int) bool {
	r := s.st.querySum(1, 0, maxRow)
	if s.M*(maxRow+1)-r < k {
		// 座位不够用了
		return false
	}
	idx := s.st.findFirst(1, maxRow, s.M-1)
	for k > 0 {
		left := min(s.M-s.st.querySum(1, idx, idx), k)
		s.st.update(1, idx, left)
		k -= left
		idx++
	}
	return true
}

type SegmentTree []Segment

// Segment  代表一个区间
type Segment struct {
	Left, Right int // 左右端点
	// 以下是需要维护的数据
	UsedMi int // 这个区间中各行已定座位的最小值
	Sum    int
}

func (st SegmentTree) build(o int, l, r int) {
	st[o].Left = l
	st[o].Right = r
	if l == r {
		return
	}
	mid := l + (r-l)/2
	st.build(o<<1, l, mid)
	st.build(o<<1|1, mid+1, r)
}

// 返回区间 [0,r] 中 <= val 的最靠左的位置，不存在时返回 -1
func (st SegmentTree) findFirst(o int, r, vl int) int {
	seg := st[o]
	if seg.UsedMi > vl {
		return -1
	}
	if seg.Left == seg.Right {
		return seg.Left
	}
	m := seg.Left + (seg.Right-seg.Left)/2
	// 找左边
	if st[o<<1].UsedMi <= vl {
		return st.findFirst(o<<1, r, vl)
	}
	// a := st.findFirst(o<<1, r, vl)
	// if a >= 0 {
	// 	在左边能找到就返回左边的，因为题目中说了返回最小的
	// return a
	// }
	if r >= m+1 {
		b := st.findFirst(o<<1|1, r, vl)
		if b >= 0 {
			return b
		}
	}
	// 都没有找到
	return -1
}

func (st SegmentTree) querySum(o int, l, r int) int {
	seg := st[o]
	if l <= seg.Left && r >= seg.Right {
		return seg.Sum
	}

	// m := seg.Left + (seg.Right-seg.Left)/2
	sum := 0
	m := (seg.Left + seg.Right) / 2
	if l <= m {
		// 有部分区间在左边
		sum += st.querySum(o<<1, l, r)
	}
	if r >= m+1 {
		sum += st.querySum(o<<1|1, l, r)
	}
	return sum
}

// 这里是单点更新
func (st SegmentTree) update(o int, idx, add int) {
	seg := st[o]
	if seg.Left == seg.Right {
		st[o].UsedMi += add
		st[o].Sum += add
		return
	}
	m := seg.Left + (seg.Right-seg.Left)/2
	if idx <= m {
		// 说明在左子树里
		st.update(o<<1, idx, add)
	}
	if idx >= m+1 {
		// 说明在右子树里
		st.update(o<<1+1, idx, add)
	}
	// 维护
	st[o].Sum = st[o<<1].Sum + st[o<<1|1].Sum
	st[o].UsedMi = min(st[o<<1].UsedMi, st[o<<1|1].UsedMi)
}
