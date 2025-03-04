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
	tree segmentTree
	d    map[int][]int
}

func Constructor(arr []int) MajorityChecker {
	d := make(map[int][]int)
	for i, ch := range arr {
		if d[ch] == nil {
			d[ch] = make([]int, 0)
		}
		d[ch] = append(d[ch], i)
	}
	tree := newSegmentTree(arr)
	return MajorityChecker{
		tree: tree,
		d:    d,
	}
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	ans := this.tree.query(1, left, right)
	if ans == nil {
		return -1
	}
	l := sort.SearchInts(this.d[ans.x], left)
	r := sort.SearchInts(this.d[ans.x], right+1)
	if r-l >= threshold {
		return ans.x
	}
	return -1
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */

type node struct {
	left  int
	right int
	x     int // 可能的 众数
	cnt   int // 摩尔投票结果，注意这里不是这个众数的个数
}

type segmentTree struct {
	tree []*node
	nums []int
}

func newSegmentTree(nums []int) segmentTree {
	n := len(nums)
	st := make([]*node, 4*n)
	for i := range st {
		st[i] = &node{}
	}
	seg := segmentTree{tree: st, nums: nums}
	seg.build(1, 0, n-1)
	return seg
}

// 找可能的众数
func (st segmentTree) query(root int, L, R int) *node {
	sg := st.tree[root]
	if L <= sg.left && R >= sg.right {
		return &node{x: sg.x, cnt: sg.cnt}
	}
	var (
		le *node
		ri *node
	)
	mid := (sg.left + sg.right) / 2
	if L <= mid {
		le = st.query(root*2, L, R)
	}
	if R >= mid+1 {
		ri = st.query(root*2+1, L, R)
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
		return &node{x: le.x, cnt: le.cnt + ri.cnt}
	}
	if le.cnt >= ri.cnt {
		return &node{x: le.x, cnt: le.cnt - ri.cnt}
	}
	return &node{x: ri.x, cnt: ri.cnt - le.cnt}
}

func (st segmentTree) build(root int, le, ri int) {
	sg := st.tree[root]
	sg.left = le
	sg.right = ri
	if le == ri {
		sg.x = st.nums[le]
		sg.cnt = 1
		return
	}
	mid := (le + ri) / 2
	st.build(root*2, le, mid)
	st.build(root*2+1, mid+1, ri)
	st.pushUp(root)
}

func (st segmentTree) pushUp(root int) {
	sg, le, ri := st.tree[root], st.tree[root*2], st.tree[root*2+1]
	if le.x == ri.x {
		sg.x = le.x
		sg.cnt = le.cnt + ri.cnt
		return
	}

	if le.cnt >= ri.cnt {
		sg.x = le.x
		sg.cnt = le.cnt - ri.cnt
		return
	}

	sg.x = ri.x
	sg.cnt = ri.cnt - le.cnt
}
