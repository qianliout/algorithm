package main

import (
	"math/bits"
	"slices"
	"sort"
)

func main() {

}

func separateSquares(squares [][]int) float64 {
	xs := make([]int, 0)
	events := make([]event, 0)
	for _, ch := range squares {
		lx, rx := ch[0], ch[0]+ch[2]
		xs = append(xs, lx, rx)
		events = append(events, event{y: ch[1], lx: lx, rx: rx, add: 1})
		events = append(events, event{y: ch[1] + ch[2], lx: lx, rx: rx, add: -1})
	}
	sort.Ints(xs)
	xs = slices.Compact(xs)

	sort.Slice(events, func(i, j int) bool { return events[i].y < events[j].y })
	tr := newSegmentTree(xs)
	all := 0
	records := make([]pair, 0)
	for i, ch := range events[:len(events)-1] {
		l := sort.SearchInts(xs, ch.lx)
		r := sort.SearchInts(xs, ch.rx) - 1
		// 为啥不这样写呢，这样才是 x 的左右区间嘛，因为我们做了离散化处理
		// tr.update(1,ch.lx,ch.rx,ch.add)
		tr.update(1, l, r, ch.add)
		sum := xs[len(xs)-1] - xs[0]
		if tr.tree[1].yMinCoverCnt == 0 {
			sum -= tr.tree[1].xMinCoverLen
		}
		records = append(records, pair{area: all, sumLen: sum})
		all += sum * (events[i+1].y - events[i].y)
	}
	// 二分查找 二分找最后一个 < totArea / 2 的面积
	m := len(squares) * 2
	i := sort.Search(m-1, func(i int) bool { return records[i].area*2 >= all }) - 1
	ans := float64(events[i].y) + float64(all-records[i].area*2)/float64(records[i].sumLen*2)

	return ans
}

type pair struct {
	sumLen int
	area   int
}

func newSegmentTree(xs []int) segmentTree {
	n := len(xs)
	tree := make([]*node, 2<<bits.Len(uint(n-1)))
	for i := range tree {
		tree[i] = &node{}
	}
	st := segmentTree{tree: tree, xs: xs}
	st.build(1, 0, n-2)
	return st
}

type node struct {
	left  int
	right int

	// 维护的数据
	xMinCoverLen int // x方向最小覆盖次数的区间所对应的长度
	yMinCoverCnt int // y方向所有区间的最小覆盖次数
	todo         int // xMinCoverL
}

type segmentTree struct {
	tree []*node
	xs   []int // x 反向的坐标
}

type event struct {
	y   int
	lx  int
	rx  int
	add int
}

func (st segmentTree) build(root int, le, ri int) {
	sg := st.tree[root]
	sg.left = le
	sg.right = ri
	if sg.left == sg.right {
		sg.xMinCoverLen = st.xs[le+1] - st.xs[le]
		return
	}
	st.pushDown(root)
	mid := (le + ri) / 2
	st.build(root*2, le, mid)
	st.build(root*2+1, mid+1, ri)
	st.pushUp(root)
}

func (st segmentTree) pushUp(root int) {
	sg, le, ri := st.tree[root], st.tree[root*2], st.tree[root*2+1]
	sg.yMinCoverCnt = min(le.yMinCoverCnt, ri.yMinCoverCnt)
	sg.xMinCoverLen = 0
	if sg.yMinCoverCnt == le.yMinCoverCnt {
		sg.xMinCoverLen += le.xMinCoverLen
	}
	if sg.yMinCoverCnt == ri.yMinCoverCnt {
		sg.xMinCoverLen += ri.xMinCoverLen
	}
}

func (st segmentTree) do(root int, add int) {
	sg := st.tree[root]
	sg.todo += add
	sg.yMinCoverCnt += add
}

func (st segmentTree) update(root int, L, R int, add int) {
	sg := st.tree[root]
	if L <= sg.left && R >= sg.right {
		st.do(root, add)
		return
	}
	mid := (sg.left + sg.right) / 2
	st.pushDown(root)
	if L <= mid {
		st.update(root*2, L, R, add)
	}
	if R >= mid+1 {
		st.update(root*2+1, L, R, add)
	}
	st.pushUp(root)
}

func (st segmentTree) pushDown(root int) {
	sg := st.tree[root]
	if sg.todo != 0 {
		st.do(root*2, sg.todo)
		st.do(root*2+1, sg.todo)
		sg.todo = 0
	}
}
