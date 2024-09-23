package main

func main() {

}

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	n := len(nums1)
	tree := &SegTree{Node: make([]Node, n*4)}
	tree.build(nums1, 1, 1, n)
	ans := make([]int64, 0)
	sum := 0
	for _, x := range nums2 {
		sum += x
	}
	for _, ch := range queries {
		if ch[0] == 1 {
			tree.update(1, ch[1]+1, ch[2]+1)
		} else if ch[0] == 2 {
			sum += tree.Query(1, 1, n) * ch[1]
		} else {
			ans = append(ans, int64(sum))
		}
	}
	return ans
}

type Node struct {
	Left  int
	Right int
	Cnt1  int  // 区间内1的个数
	Flip  bool // 执行反转,相当于 lazy
}

type SegTree struct {
	Node []Node
}

func (s *SegTree) build(nums []int, rootId, l, r int) {
	s.Node[rootId].Left = l
	s.Node[rootId].Right = r
	if l == r {
		s.Node[rootId].Cnt1 = nums[l-1]
		return
	}
	mid := (l + r) >> 1
	s.build(nums, rootId<<1, l, mid)
	s.build(nums, rootId<<1|1, mid+1, r)
	s.maintain(rootId)
}

func (s *SegTree) maintain(rootId int) {
	s.Node[rootId].Cnt1 = s.Node[rootId<<1].Cnt1 + s.Node[rootId<<1|1].Cnt1
}
func (s *SegTree) update(rootId int, l, r int) {
	root := s.Node[rootId]
	if l <= root.Left && r >= root.Right {
		s.do(rootId)
		return
	}
	s.pushDown(rootId)

	mid := (root.Left + root.Right) >> 1

	if l <= mid {
		s.update(rootId<<1, l, r)
	}
	if mid+1 <= r {
		s.update(rootId<<1|1, l, r)
	}
	s.maintain(rootId)
}

func (s *SegTree) Query(rootId, l, r int) int {
	root := s.Node[rootId]
	if l <= root.Left && r >= root.Right {
		return s.Node[rootId].Cnt1
	}
	s.pushDown(rootId)
	mid := (root.Left + root.Right) >> 1
	ans := 0
	if l <= mid {
		ans += s.Query(rootId<<1, l, r)
	}
	if mid+1 <= r {
		ans += s.Query(rootId, l, r)
	}
	return ans
}

func (s *SegTree) pushDown(rootId int) {
	root := s.Node[rootId]
	if root.Flip {
		s.do(rootId << 1)
		s.do(rootId<<1 | 1)
		s.Node[rootId].Flip = false
	}
}

func (s *SegTree) do(rootId int) {
	root := s.Node[rootId]
	s.Node[rootId].Cnt1 = root.Right - root.Left + 1 - s.Node[rootId].Cnt1
	s.Node[rootId].Flip = !s.Node[rootId].Flip
}

// 设 nums1 中总共有 c 个 1，那么操作 2 相当于把 nums2的元素和增加了 c⋅p。所以只需要维护 nums1 中 1 的个数。
// 如何实现操作 1？用 Lazy 线段树维护区间内 1 的个数 cnt1，以及整个区间是否需要反转的 Lazy 标记 flip
