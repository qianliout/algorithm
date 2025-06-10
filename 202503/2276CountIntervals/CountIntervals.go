package main

import (
	"math"
)

func main() {

}

type CountIntervals struct {
	tr   segmentTree
	root *node
	inf  int
}

func Constructor() CountIntervals {
	inf := int(math.Pow10(9)) + 10
	root := &node{lx: 1, rx: inf}
	tr := segmentTree{tree: root}
	return CountIntervals{tr: tr, root: root, inf: inf}
}

func (s *CountIntervals) Add(left int, right int) {
	s.tr.update(s.root, left, right, 1)
}

func (s *CountIntervals) Count() int {
	ans := s.tr.query(s.root, 1, s.inf)
	return ans
}

/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */

type node struct {
	left  *node
	right *node
	lx    int // 左边 x
	rx    int // 右边 x
	cnt   int // 区间中已跟中的实数
	todo  int
}

type segmentTree struct {
	tree *node
}

// 新增
func (st segmentTree) update(root *node, L, R, add int) {
	le, ri := root.lx, root.rx
	if L <= le && R >= ri {
		st.do(root, add)
		return
	}
	mid := (le + ri) >> 1

	if root.left == nil {
		root.left = &node{lx: le, rx: mid}
		root.right = &node{lx: mid + 1, rx: ri}
	}

	st.pushDown(root)
	if L <= mid {
		st.update(root.left, L, R, add)
	}
	if R >= mid+1 {
		st.update(root.right, L, R, add)
	}

	st.pushUp(root)
}

func (st segmentTree) pushDown(root *node) {
	if root.todo == 0 {
		return
	}
	st.do(root.left, root.todo)
	st.do(root.right, root.todo)
	root.todo = 0
}

func (st segmentTree) pushUp(root *node) {
	le, ri := root.left, root.right
	root.cnt = min(root.rx-root.lx+1, le.cnt+ri.cnt)
}

func (st segmentTree) do(root *node, add int) {
	le, ri := root.lx, root.rx
	root.cnt = ri - le + 1
	root.todo = add
}

func (st segmentTree) query(root *node, L, R int) int {
	if L <= root.lx && R >= root.rx {
		return root.cnt
	}

	mid := (root.lx + root.rx) >> 1

	if root.left == nil {
		root.left = &node{lx: root.lx, rx: mid}
		root.right = &node{lx: mid + 1, rx: root.rx}
	}
	st.pushDown(root)
	res := 0
	if L <= mid {
		res += st.query(root.left, L, R)
	}
	if R >= mid+1 {
		res += st.query(root.right, L, R)
	}
	return res
}
