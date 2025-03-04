package main

import (
	"fmt"
	"sort"
)

func main() {
	mc := Constructor([]int{1, 1, 2, 2, 1, 1})
	fmt.Println(mc.Query(0, 5, 4))
	fmt.Println(mc.Query(0, 3, 3))
	fmt.Println(mc.Query(2, 3, 2))
}

type MajorityChecker struct {
	tree SegmentTree
	d    map[int][]int
}

func Constructor(arr []int) MajorityChecker {
	tr := NewSegmentTree(arr)
	d := make(map[int][]int)
	for i, ch := range arr {
		if d[ch] == nil {
			d[ch] = make([]int, 0)
		}
		d[ch] = append(d[ch], i)
	}
	return MajorityChecker{tree: tr, d: d}
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	node := this.tree.query(1, left, right)
	if node == nil {
		return -1
	}
	// 为啥使用二分，没有想明白
	l := sort.SearchInts(this.d[node.x], left)
	r := sort.SearchInts(this.d[node.x], right+1)
	if r-l >= threshold {
		return node.x
	}
	return -1
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */

type segment struct {
	left  int
	right int
	x     int //  这个区间的众数
	cnt   int // 上面这个众数的个数
}

type SegmentTree struct {
	tree []*segment
	nums []int
}

func NewSegmentTree(nums []int) SegmentTree {
	n := len(nums)
	tr := make([]*segment, n*4)
	for i := range tr {
		tr[i] = &segment{}
	}
	s := SegmentTree{tree: tr, nums: nums}
	s.build(1, 0, n-1)
	return s
}

func (st *SegmentTree) build(o int, l, r int) {
	sg := st.tree[o]
	sg.left = l
	sg.right = r
	if sg.left == sg.right {
		sg.x = st.nums[l]
		sg.cnt = 1
		return
	}
	mid := (l + r) >> 1
	st.build(o<<1, l, mid)
	st.build(o<<1|1, mid+1, r)
	st.pushUp(o)
}

func (st *SegmentTree) pushUp(o int) {
	// 因为传的是指针，所以可以这样更新
	sg, le, ri := st.tree[o], st.tree[o<<1], st.tree[o<<1|1]

	if le.x == ri.x {
		sg.x = le.x
		sg.cnt = le.cnt + ri.cnt
		return
	}
	// 不相等，摩尔投票
	if le.cnt >= ri.cnt {
		sg.x = le.x
		sg.cnt = le.cnt - ri.cnt
	} else {
		sg.x = ri.x
		sg.cnt = ri.cnt - le.cnt
	}
}

func (st *SegmentTree) query(o int, L, R int) *segment {
	sg := st.tree[o]
	if L <= sg.left && R >= sg.right {
		// 这里一定要重新赋值，因为下面要做更改，这种错很难发现，所以对于返回值还是不用使用 segment 的好
		return &segment{x: sg.x, cnt: sg.cnt}
	}
	m := (sg.left + sg.right) / 2
	var (
		le *segment
		ri *segment
	)
	if L <= m {
		le = st.query(o<<1, L, R)
	}
	if R >= m+1 {
		ri = st.query(o<<1|1, L, R)
	}
	if le == nil && ri == nil {
		return nil
	}
	if le == nil {
		return ri
	}
	if ri == nil {
		return le
	}

	if le.x == ri.x {
		le.cnt += ri.cnt
		return le
	}
	if le.cnt >= ri.cnt {
		le.cnt -= ri.cnt
		return le
	}
	ri.cnt -= le.cnt
	return ri
}
