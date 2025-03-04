package main

import (
	"fmt"
)

func main() {
	// fmt.Println(leftmostBuildingQueries([]int{6, 4, 8, 5, 2, 7}, [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}})) // 2,5,-1,5,2
	fmt.Println(leftmostBuildingQueries([]int{1, 2, 1, 2, 1, 2}, [][]int{{0, 0}, {0, 1}, {0, 2}})) // 0,1,3
}

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	n := len(queries)
	st := newSegmentTree(heights)
	ans := make([]int, n)

	for i, ch := range queries {
		x, y := ch[0], ch[1]
		if x > y {
			x, y = y, x
		}
		// 这个判断是难点
		// 如果一个人在建筑 i ，且存在 i < j 的建筑 j 满足 heights[i] < heights[j] ，那么这个人可以移动到建筑 j
		// 所以这里只能是 heights[y] > heights[x]
		if x == y || heights[y] > heights[x] {
			ans[i] = y
			continue
		}
		ans[i] = st.queryFirst(1, param{left: y, right: n - 1, he: max(heights[x], heights[y])})
	}
	return ans
}

type node struct {
	left  int
	right int
	mx    int // 区间最大值
}

func newSegmentTree(heights []int) segmentTree {
	n := len(heights)
	tr := make([]*node, 4*n)
	for i := range tr {
		tr[i] = &node{mx: 0}
	}
	s := segmentTree{tree: tr, heights: heights}
	s.build(1, 0, n-1)
	return s
}

type segmentTree struct {
	tree    []*node
	heights []int
}

func (st segmentTree) build(root int, l, r int) {
	sg := st.tree[root]
	sg.left = l
	sg.right = r
	if sg.left == sg.right {
		sg.mx = st.heights[l]
		return
	}
	mid := (l + r) / 2
	st.build(root*2, l, mid)
	st.build(root*2+1, mid+1, r)
	st.pushUp(root)
}

func (st segmentTree) pushUp(root int) {
	sg, le, ri := st.tree[root], st.tree[root*2], st.tree[root<<1|1]
	sg.mx = max(le.mx, ri.mx)
}

type param struct {
	left  int
	right int
	he    int
}

// 返回 [L,n-1] 中第一个 > v 的值的 最小下标
func (st segmentTree) queryFirst(root int, pa param) int {
	sg := st.tree[root]
	// 这里找的是大于 v
	if sg.mx <= pa.he {
		return -1
	}
	if sg.left == sg.right {
		return sg.left
	}
	mid := (sg.left + sg.right) / 2
	// 先找左边
	if pa.left <= mid {
		a := st.queryFirst(root*2, pa)
		if a >= 0 {
			return a
		}
	}
	return st.queryFirst(root*2+1, pa)
}
