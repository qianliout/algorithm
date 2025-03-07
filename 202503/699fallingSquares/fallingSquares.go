package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(fallingSquares([][]int{{1, 2}, {2, 3}, {6, 1}}))
}

func fallingSquares(positions [][]int) []int {
	ans := make([]int, 0)
	inf := int(math.Pow10(10))
	root := &node{rx: inf}
	tr := segmentTree{tree: root}
	for _, ch := range positions {
		x := ch[0]
		h := ch[1]
		cur := tr.query(root, x, x+h-1)
		tr.update(root, x, x+h-1, cur+h)
		ans = append(ans, tr.query(root, 0, inf))
	}
	return ans
}

type node struct {
	left  *node
	right *node
	lx    int // 左边 x
	rx    int // 右边 x
	mxh   int // 这个区间里最大的高度
	todo  int
}

type segmentTree struct {
	tree *node
}

// 注意这里的 mxh 是把这个区间都更新为mxh,而不是新增
func (st segmentTree) update(root *node, L, R, mxh int) {
	le, ri := root.lx, root.rx
	if L <= le && R >= ri {
		st.do(root, mxh)
		return
	}
	mid := (le + ri) >> 1

	if root.left == nil {
		root.left = &node{lx: le, rx: mid}
		root.right = &node{lx: mid + 1, rx: ri}
	}

	st.pushDown(root)
	if L <= mid {
		st.update(root.left, L, R, mxh)
	}
	if R >= mid+1 {
		st.update(root.right, L, R, mxh)
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
	root.mxh = max(le.mxh, ri.mxh)
}

func (st segmentTree) do(root *node, mxh int) {
	root.mxh = max(root.mxh, mxh)
	root.todo = mxh
}

func (st segmentTree) query(root *node, L, R int) int {
	if L <= root.lx && R >= root.rx {
		return root.mxh
	}

	mid := (root.lx + root.rx) >> 1

	if root.left == nil {
		root.left = &node{lx: root.lx, rx: mid}
		root.right = &node{lx: mid + 1, rx: root.rx}
	}
	st.pushDown(root)
	res := 0
	if L <= mid {
		res = max(res, st.query(root.left, L, R))
	}
	if R >= mid+1 {
		res = max(res, st.query(root.right, L, R))
	}
	return res
}
