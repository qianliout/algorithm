package main

func main() {

}

func sumCounts(nums []int) int {

}

// 1 <= nums.length <= 105
// 1 <= nums[i] <= 105

type node struct {
	left  int
	right int

	cnt  int
	todo int
	sum  int
}

type segmentTree struct {
	tree []*node
	last map[int]int // 这个数上一次存在的下标,如果不存在，那就是0，所以在计算下标时整体加1
}

func (st segmentTree) build(root int, l, r int) {

}

func (st segmentTree) update(root int, L, R, ADD int) {
	sg := st.tree[root]
	if L <= sg.left && R >= sg.right {

	}
}

func (st segmentTree) query(root int, L, R int) int {

}

func (st segmentTree) pushUp(root int) {

}

func (st segmentTree) PushDown(root int) {

}

func (st segmentTree) do(root int, add int) {
	sg := st.tree[root]
	sg.sum += add * (sg.right - sg.left + 1)
	sg.todo += add
}
