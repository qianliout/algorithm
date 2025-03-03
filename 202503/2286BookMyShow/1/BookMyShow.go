package main

import (
	"fmt"
)

func main() {
	b := Constructor(2, 5)
	fmt.Println(b.Gather(4, 0))
	fmt.Println(b.Gather(2, 0))
	fmt.Println(b.Scatter(5, 1))
}

type BookMyShow struct {
	ST   SegmentTree
	N, M int
}

func Constructor(n int, m int) BookMyShow {
	t := SegmentTree{Data: make([]Node, 4*n)}
	t.build(1, 0, n-1)

	return BookMyShow{
		ST: t,
		M:  m,
		N:  n,
	}
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	r := this.ST.FindFirst(1, maxRow, this.M-k)
	if r < 0 {
		return nil
	}
	c := this.ST.querySum(1, r, r)
	this.ST.update(1, r, k)
	return []int{r, c}
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	s := this.ST.querySum(1, 0, maxRow)
	if s > this.M*(maxRow+1)-k {
		return false
	}
	i := this.ST.FindFirst(1, maxRow, this.M-1)
	for k > 0 {
		left := min(this.M-this.ST.querySum(1, i, i), k)
		this.ST.update(1, i, left)
		k -= left
		i++
	}
	return true
}

/**
 * Your BookMyShow object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Gather(k,maxRow);
 * param_2 := obj.Scatter(k,maxRow);
 */

type Node struct {
	Left  int
	Right int
	// 这些都可以用线段树解决。线段树维护每个区间的接水量的最小值 min，以及每个区间的接水量之和 sum。
	Mi  int
	Sum int
}

type SegmentTree struct {
	Data []Node
}

func (st *SegmentTree) build(o int, l, r int) {
	st.Data[o].Left = l
	st.Data[o].Right = r
	if l == r {
		return
	}
	// m := l + (r-l)/2
	m := (l + r) >> 1

	st.build(o<<1, l, m)
	st.build(o*2+1, m+1, r)
}

// 把下标 i 的元素增加 vl,
// o 表示根节点
func (st *SegmentTree) update(o int, i, vl int) {
	if st.Data[o].Left == st.Data[o].Right {
		st.Data[o].Mi += vl
		st.Data[o].Sum += vl
		return
	}
	m := (st.Data[o].Left + st.Data[o].Right) >> 1
	if i <= m {
		st.update(o<<1, i, vl)
	} else {
		st.update(o<<1+1, i, vl)
	}
	lo, ro := st.Data[o<<1], st.Data[o<<1|1]
	st.Data[o].Mi = min(lo.Mi, ro.Mi)
	st.Data[o].Sum = lo.Sum + ro.Sum
}

func (st *SegmentTree) querySum(o int, l, r int) int {
	if l <= st.Data[o].Left && r >= st.Data[o].Right {
		return st.Data[o].Sum
	}
	m := (st.Data[o].Left + st.Data[o].Right) >> 1
	sum := 0
	if l <= m {
		sum += st.querySum(o<<1, l, r)
	}
	if r >= m+1 {
		sum += st.querySum(o<<1|1, l, r)
	}
	return sum
}

// 返回区间 [0,r] 中 <= val 的最靠左的位置，不存在时返回 -1
func (st *SegmentTree) FindFirst(o int, r int, vl int) int {
	if st.Data[o].Mi > vl {
		return -1
	}
	if st.Data[o].Left == st.Data[o].Right {
		return st.Data[o].Left
	}
	m := (st.Data[o].Left + st.Data[o].Right) >> 1
	if st.Data[o<<1].Mi <= vl {
		return st.FindFirst(o<<1, r, vl)
	}
	if r > m {
		return st.FindFirst(o*2+1, r, vl)
	}
	return -1
}
