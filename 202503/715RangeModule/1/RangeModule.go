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
	return this.tr.query(this.root, left, right-1)
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
	lx    int  // 左边 x
	rx    int  // 右边 x
	all   bool // 这个区间是否都覆盖了
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
	root.all = le.all && ri.all
}

func (st segmentTree) do(root *node, mxh int) {
	if mxh == -1 {
		root.all = false
	} else if mxh == 1 {
		root.all = true
	}
	root.todo = mxh
}

func (st segmentTree) query(root *node, L, R int) bool {
	if L <= root.lx && R >= root.rx {
		return root.all
	}

	mid := (root.lx + root.rx) >> 1

	if root.left == nil {
		root.left = &node{lx: root.lx, rx: mid}
		root.right = &node{lx: mid + 1, rx: root.rx}
	}
	st.pushDown(root)
	res := true
	if L <= mid {
		res = res && st.query(root.left, L, R)
	}
	if R >= mid+1 {
		res = res && st.query(root.right, L, R)
	}
	return res
}
