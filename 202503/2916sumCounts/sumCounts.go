package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sumCounts([]int{1, 2, 1}))
}

// 不对，为啥呢
func sumCounts(nums []int) int {
	n := len(nums)
	st := newSegmentTree(n)
	ans := 0
	mod := int(math.Pow10(9)) + 7
	s := 0
	for i, x := range nums {
		i++
		j := st.last[x]
		b := st.query(1, j+1, i)
		s += b*2 + (i - j)
		ans += s
		st.update(1, j+1, i, 1)
		st.last[x] = i
	}
	return ans % mod
}

// 1 <= nums.length <= 105
// 1 <= nums[i] <= 105

type node struct {
	left  int
	right int

	todo int
	sum  int
}

func newSegmentTree(n int) segmentTree {
	st := segmentTree{tree: make([]*node, (n+1)*4), last: map[int]int{}}
	for i := range st.tree {
		st.tree[i] = &node{}
	}
	st.build(1, 1, n)
	return st
}

type segmentTree struct {
	tree []*node
	last map[int]int // 这个数上一次存在的下标,如果不存在，那就是0，所以在计算下标时整体加1
}

func (st segmentTree) build(root int, l, r int) {
	sg := st.tree[root]
	sg.left = l
	sg.right = r
	if sg.left == sg.right {
		return
	}
	mid := (l + r) >> 1
	st.PushDown(root)
	st.build(root*2, l, mid)
	st.build(root*2+1, mid+1, r)
	st.pushUp(root)
}

func (st segmentTree) update(root int, L, R, ADD int) {
	sg := st.tree[root]
	if L <= sg.left && R >= sg.right {
		st.do(root, ADD)
		return
	}

	st.PushDown(root)
	mid := (sg.left + sg.right) >> 1

	if L <= mid {
		st.update(root*2, L, R, ADD)
	}
	if R >= mid+1 {
		st.update(root*2+1, L, R, ADD)
	}
}

func (st segmentTree) query(root int, L, R int) int {
	sg := st.tree[root]
	if L <= sg.left && R >= sg.right {
		return sg.sum
	}
	st.PushDown(root)

	mid := (sg.left + sg.right) >> 1
	res := 0
	if L <= mid {
		res += st.query(root*2, L, R)
	}
	if R >= mid+1 {
		res += st.query(root*2+1, L, R)
	}
	return res
}

func (st segmentTree) pushUp(root int) {
	sg, le, ri := st.tree[root], st.tree[root*2], st.tree[root*2+1]
	sg.sum = le.sum + ri.sum
}

func (st segmentTree) PushDown(root int) {
	sg := st.tree[root]
	if sg.todo != 0 {
		st.do(root*2, sg.todo)
		st.do(root*2+1, sg.todo)
		sg.todo = 0
	}
}

func (st segmentTree) do(root int, add int) {
	sg := st.tree[root]
	sg.sum += add * (sg.right - sg.left + 1)
	sg.todo += add
}
