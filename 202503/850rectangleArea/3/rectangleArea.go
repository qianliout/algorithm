package main

import (
	"fmt"
	"math/bits"
	"slices"
	"sort"
)

func main() {
	fmt.Println(rectangleArea([][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}}))
	fmt.Println(rectangleArea([][]int{{0, 0, 3, 3}, {2, 0, 5, 3}, {1, 1, 4, 4}}))
}

func rectangleArea(rectangles [][]int) (ans int) {
	xs := make([]int, 0) // 所有 x 坐标
	events := make([]event, 0)
	for _, ch := range rectangles {
		lx, rx := ch[0], ch[2]
		xs = append(xs, lx, rx)
		events = append(events, event{y: ch[1], lx: lx, rx: rx, add: 1})
		events = append(events, event{y: ch[3], lx: lx, rx: rx, add: -1})
	}
	sort.Ints(xs)
	xs = slices.Compact(xs)
	// 以y坐标排序
	sort.Slice(events, func(i, j int) bool { return events[i].y < events[j].y })
	tr := newSegmentTree(xs)
	sumLen := xs[len(xs)-1] - xs[0]
	for i, e := range events[:len(events)-1] {
		l := sort.SearchInts(xs, e.lx)
		r := sort.SearchInts(xs, e.rx) - 1
		tr.update(1, l, r, e.add)
		a := sumLen
		if tr.tree[1].yMinCover == 0 {
			a -= tr.tree[1].xMinCoverLen
		}
		ans += a * (events[i+1].y - e.y)
	}

	mod := 1000000007
	return ans % mod
}

type event struct {
	y   int
	lx  int
	rx  int
	add int
}

type node struct {
	left         int
	right        int
	yMinCover    int // 扫描线进行扫描时的最小覆盖次数,本次解法是y作为扫描线，也就是 y 方向的最小覆盖次数
	xMinCoverLen int // 在扫描过程中x 方向最小覆盖次数的最小边长
	todo         int // 子树内的所有节点的 yMinCover 需要增加的量，注意y走出区间时就是负数
}

func newSegmentTree(xs []int) *segmentTree {
	n := len(xs)
	tree := make([]*node, 2<<bits.Len(uint(n)))
	for i := range tree {
		tree[i] = &node{}
	}
	tr := &segmentTree{tree: tree, xs: xs}
	// 有n个值，只有 n-1个区间，这里 n-2 是一个想不到的地方
	tr.build(1, 0, n-2)
	return tr
}

type segmentTree struct {
	tree []*node
	xs   []int // 矩形的x坐标
}

func (st *segmentTree) build(root int, le, ri int) {
	sg := st.tree[root]
	sg.left = le
	sg.right = ri
	if sg.left == sg.right {
		// 此时只是求的 x 方向上区间，y方向还没有开始扫描
		sg.yMinCover = 0 // 这里为啥是0，没有能理解
		sg.xMinCoverLen = st.xs[le+1] - st.xs[le]
		return
	}
	mid := (le + ri) / 2
	st.build(root<<1, le, mid)
	st.build(root<<1|1, mid+1, ri)
	st.pushUp(root)
}

// 下传懒标记 todo
func (st *segmentTree) pushDown(root int) {
	sg := st.tree[root]
	if sg.todo != 0 {
		st.do(root<<1, sg.todo)
		st.do(root<<1|1, sg.todo)
		sg.todo = 0
	}
}

// 仅更新节点信息，不下传懒标记 todo
func (st *segmentTree) do(root, v int) {
	sg := st.tree[root]
	sg.yMinCover += v
	sg.todo += v
}

// 根据左右子节点更新 root 节点
func (st *segmentTree) pushUp(root int) {
	sg, le, ri := st.tree[root], st.tree[root<<1], st.tree[root<<1|1]
	sg.yMinCover = min(le.yMinCover, ri.yMinCover)
	sg.xMinCoverLen = 0 // 这里不赋值成0是不可以的
	if le.yMinCover == sg.yMinCover {
		sg.xMinCoverLen += le.xMinCoverLen
	}
	if ri.yMinCover == sg.yMinCover {
		sg.xMinCoverLen += ri.xMinCoverLen
	}
}

func (st *segmentTree) update(root int, le, ri int, v int) {
	sg := st.tree[root]
	if le <= sg.left && ri >= sg.right {
		st.do(root, v)
		return
	}
	st.pushDown(root)
	mid := (sg.left + sg.right) / 2
	if le <= mid {
		st.update(root<<1, le, ri, v)
	}
	if ri >= mid+1 {
		st.update(root<<1|1, le, ri, v)
	}
	st.pushUp(root)
}
