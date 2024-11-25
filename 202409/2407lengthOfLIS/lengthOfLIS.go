package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{4, 2, 1, 4, 3, 4, 5, 8, 15}, 3))
}

type Node struct {
	Left  int
	Right int
	Max   int
}

type SegTree struct {
	Node []Node
}

func (s *SegTree) Update(rootId int, idx int, value int) {

	root := s.Node[rootId]
	if root.Left == root.Right {
		s.Node[rootId].Max = value
		return
	}
	mid := (root.Left + root.Right) >> 1

	// 这样写也是可以的
	// if idx <= mid {
	// 	s.Update(rootId<<1, idx, value)
	// } else {
	// 	s.Update(rootId<<1|1, idx, value)
	// }

	if idx <= mid {
		s.Update(rootId<<1, idx, value)
	}

	if idx >= mid+1 {
		s.Update(rootId<<1|1, idx, value)
	}

	s.Node[rootId].Max = max(s.Node[rootId<<1].Max, s.Node[rootId<<1|1].Max)
}

func (s *SegTree) Build(rootId, l, r int) {
	if l > r {
		return
	}
	s.Node[rootId].Left = l
	s.Node[rootId].Right = r
	if l == r {
		// 初值不能是1，只能是0，因为这是值域dp,如果 nums 中没有 rootId,那么初值就是0
		// s.Head[rootId].Max = 1 // 初始值是
		return
	}
	mid := (l + r) >> 1
	s.Build(rootId<<1, l, mid)
	s.Build(rootId<<1|1, mid+1, r)
}
func (s *SegTree) Query(rootId int, l, r int) int {
	root := s.Node[rootId]
	if l <= root.Left && r >= root.Right {
		return root.Max
	}
	mid := (root.Left + root.Right) >> 1
	// 为啥初值是0而不是1呢，因为这是值域dp,如果[l,r]不在 nums 中那么值就只能是0
	ans := 0 // 这里的初值一定是0
	if l <= mid {
		ans = max(ans, s.Query(rootId<<1, l, r))
	}
	if r >= mid+1 {
		ans = max(ans, s.Query(rootId<<1|1, l, r))
	}
	return ans

	// 这样写也可以
	// if r <= mid {
	// 	return s.Query(rootId<<1, l, r)
	// }
	// if l >= mid+1 {
	// 	return s.Query(rootId<<1|1, l, r)
	// }
	// return max(s.Query(rootId<<1, l, r), s.Query(rootId<<1|1, l, r))
}

func lengthOfLIS(nums []int, k int) int {
	mx := 0
	for _, ch := range nums {
		mx = max(mx, ch)
	}
	tree := &SegTree{Node: make([]Node, mx*4)}
	tree.Build(1, 1, mx)

	for _, ch := range nums {
		if ch == 1 {
			tree.Update(1, 1, 1) // 初值
		} else {
			pre := tree.Query(1, max(ch-k, 1), ch-1)
			tree.Update(1, ch, pre+1)
		}
	}
	return tree.Query(1, 1, mx)
}
