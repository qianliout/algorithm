package main

import "fmt"

func main() {
	b := Constructor(4, 5)
	fmt.Println(b.Scatter(6, 2))
	fmt.Println(b.Gather(6, 3))
	fmt.Println(b.Scatter(9, 1))

	b2 := Constructor(2, 5)
	fmt.Println(b2.Gather(4, 0))
	fmt.Println(b2.Gather(2, 0))
	fmt.Println(b2.Scatter(5, 1))
	fmt.Println(b2.Scatter(5, 1))

	b3 := Constructor(5, 9)
	b3.Gather(10, 1)
	b3.Scatter(3, 3)
	b3.Gather(9, 1)
	b3.Gather(10, 2)
	fmt.Println(b3.Gather(2, 2))

}

type BookMyShow struct {
	st   SegmentTree
	m, n int
}

func Constructor(n int, m int) BookMyShow {
	st := make(SegmentTree, 4*n)
	for i := range st {
		st[i] = &Segment{}
	}
	st.build(1, 0, n-1)

	return BookMyShow{st, m, n}
}

func (s *BookMyShow) Gather(k int, maxRow int) []int {
	r := s.st.findFirst(1, maxRow, s.m-k)
	if r < 0 {
		return nil
	}
	c := s.st.querySum(1, r, r)
	s.st.update(1, r, k)
	return []int{r, c}
}

func (s *BookMyShow) Scatter(k int, maxRow int) bool {
	sum := s.st.querySum(1, 0, maxRow)
	if s.m*(maxRow+1)-sum < k {
		return false
	}
	idx := s.st.findFirst(1, maxRow, s.m-1)
	for k > 0 {
		left := min(k, s.m-s.st.querySum(1, idx, idx))
		s.st.update(1, idx, left)
		k -= left
		idx++
	}
	return true
}

/**
 * Your BookMyShow object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Gather(k,maxRow);
 * param_2 := obj.Scatter(k,maxRow);
 */

type Segment struct {
	left, right int
	Mi          int
	Sum         int
}

type SegmentTree []*Segment

func (s SegmentTree) build(o int, l, r int) {
	s[o].left = l
	s[o].right = r

	if l == r {
		return
	}
	m := (l + r) / 2
	s.build(o<<1, l, m)
	s.build(o*2+1, m+1, r)
}

// 单点更新
func (s SegmentTree) update(o int, idx, add int) {
	seg := s[o]
	if seg.left == seg.right {
		// 如果这样更新的话，那么数组里一定得存指针
		seg.Mi += add
		seg.Sum += add
		return
	}
	m := (seg.left + seg.right) / 2
	if idx <= m {
		s.update(o<<1, idx, add)
	}
	if idx >= m+1 {
		s.update(o*2+1, idx, add)
	}

	le, ri := s[o<<1], s[o*2+1]
	s[o].Mi = min(le.Mi, ri.Mi)
	s[o].Sum = le.Sum + ri.Sum
}

// 找 [0:r]内小于等于vl 的最小下标 (包括 r)
func (s SegmentTree) findFirst(o int, r int, al int) int {
	seg := s[o]
	if seg.Mi > al {
		return -1
	}
	if seg.left == seg.right {
		return seg.left
	}

	m := (seg.left + seg.right) / 2
	// 先找左边
	// 因为规范是从0开始，所以左边一定在区间内，可以不用判断
	// if s[o<<1].Mi <= al {
	// 	return s.findFirst(o<<1, r, al)
	// }
	a := s.findFirst(o<<1, r, al)
	if a >= 0 {
		return a
	}
	if r >= m+1 {
		b := s.findFirst(o*2+1, r, al)
		if b >= 0 {
			return b
		}
	}
	return -1
}

func (s SegmentTree) querySum(o int, l, r int) int {
	seg := s[o]
	if l <= seg.left && r >= seg.right {
		return seg.Sum
	}
	m := (seg.left + seg.right) / 2
	sum := 0
	if l <= m {
		sum += s.querySum(o<<1, l, r)
	}
	if r >= m+1 {
		sum += s.querySum(o*2+1, l, r)
	}
	return sum
}
