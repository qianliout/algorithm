package main

import (
	"math"
)

func main() {

}

func maximumSumSubsequence(nums []int, queries [][]int) int {
	st := newSeg(nums)
	n := len(nums)
	ans := 0
	mod := int(math.Pow10(9)) + 7
	for _, ch := range queries {
		pa := updateParam{idx: ch[0], val: ch[1]}
		st.update(1, 0, n-1, pa)
		ans += st.tree[1].data11
		ans = ans % mod
	}
	return ans
}

// 1 <= nums.length <= 5 * 104
// -105 <= nums[i] <= 105
// 1 <= queries.length <= 5 * 104
// queries[i] == [posi, xi]
// 0 <= posi <= nums.length - 1
// -105 <= xi <= 105

type node struct {
	left  int
	right int

	// 维护的数据
	data00 int // 第一个不选，最后一个也不选
	data01 int // 第一个不选，最后一个可选可不选
	data10 int // 第一个可选可不选，最后一个不选
	data11 int // 第一个可选可不选，最后一个可选可不选
}

func newSeg(nums []int) segmentTree {
	n := len(nums)
	tree := make([]*node, 4*n)
	for i := range tree {
		tree[i] = &node{}
	}
	st := segmentTree{tree: tree, nums: nums}
	st.build(1, 0, n-1)
	return st
}

type segmentTree struct {
	tree []*node
	nums []int
}

func (st segmentTree) pushUp(root int) {
	sg, le, ri := st.tree[root], st.tree[root<<1], st.tree[root<<1|1]
	sg.data00 = max(le.data00+ri.data10, le.data01+ri.data00)
	sg.data01 = max(le.data00+ri.data11, le.data01+ri.data01)
	sg.data10 = max(le.data10+ri.data10, le.data11+ri.data00)
	sg.data11 = max(le.data10+ri.data11, le.data11+ri.data01)
}

func (st segmentTree) build(root int, l, r int) {
	sg := st.tree[root]
	sg.left = l
	sg.right = r
	if sg.left == sg.right {
		sg.data11 = max(0, st.nums[l])
		return
	}
	mid := (l + r) / 2
	st.build(root*2, l, mid)
	st.build(root*2+1, mid+1, r)
	st.pushUp(root)
}

type updateParam struct {
	idx int
	val int
}

func (st segmentTree) update(root int, l, r int, pa updateParam) {
	sg := st.tree[root]
	if l == r {
		sg.data11 = max(0, pa.val)
		return
	}
	m := (l + r) / 2
	if pa.idx <= m {
		st.update(root*2, l, m, pa)
	}
	if pa.idx >= m+1 {
		st.update(root*2+1, m+1, r, pa)
	}
	st.pushUp(root)
}
