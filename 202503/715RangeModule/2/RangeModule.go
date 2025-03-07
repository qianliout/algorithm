package main

import (
	"fmt"
	"math"
)

func main() {
	cm := Constructor()
	cm.AddRange(10, 20)
	cm.RemoveRange(14, 16)
	fmt.Println(cm.QueryRange(10, 14))
	fmt.Println(cm.QueryRange(13, 15))
	fmt.Println(cm.QueryRange(16, 17))
}

type RangeModule struct {
	tr   segmentTree
	root *node
}

func Constructor() RangeModule {
	inf := int(math.Pow10(10))
	root := &node{rx: inf}
	tr := segmentTree{tree: root}
	return RangeModule{tr: tr, root: root}
}

func (this *RangeModule) AddRange(left int, right int) {
	this.tr.update(this.root, left, right-1, 1) // 半开的
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	pre := this.tr.query(this.root, left, right-1)
	return pre == (right - left)
}

func (this *RangeModule) RemoveRange(left int, right int) {
	this.tr.update(this.root, left, right-1, -1)
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
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
	root.cnt = le.cnt + ri.cnt
}

func (st segmentTree) do(root *node, mxh int) {
	le, ri := root.lx, root.rx
	if mxh == -1 {
		root.cnt = 0
	} else if mxh == 1 {
		root.cnt = ri - le + 1
	}
	root.todo = mxh
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
