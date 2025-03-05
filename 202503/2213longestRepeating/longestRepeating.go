package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestRepeating("babacc", "bcb", []int{1, 3, 3}))
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	ss := []byte(s)

	st := make([]*node, len(ss)*4)
	for i := range st {
		st[i] = &node{}
	}
	seg := segmentTree{tree: st, ss: ss}
	seg.build(1, 0, len(ss)-1)

	ans := make([]int, len(queryIndices))
	for i, idx := range queryIndices {
		seg.ss[idx] = byte(queryCharacters[i])
		seg.update(1, idx)
		ans[i] = seg.tree[1].mx
	}
	return ans
}

type node struct {
	left  int // 理解：left	表示的是这个节点的区间的 left，这就是 ss的left
	right int
	pre   int
	suf   int
	mx    int
}

type segmentTree struct {
	tree []*node
	ss   []byte
}

func (st segmentTree) build(root int, l, r int) {
	sg := st.tree[root]
	sg.left = l
	sg.right = r
	if sg.left == sg.right {
		sg.pre = 1
		sg.suf = 1
		sg.mx = 1
		return
	}
	mid := (l + r) / 2
	st.build(root*2, l, mid)
	st.build(root*2+1, mid+1, r)
	st.pushUp(root)
}

func (st segmentTree) pushUp(root int) {
	sg, le, ri := st.tree[root], st.tree[root*2], st.tree[root*2+1]
	sg.suf = ri.suf
	sg.pre = le.pre
	sg.mx = max(le.mx, ri.mx)
	// 如果中间字符相同
	m := (sg.left + sg.right) / 2
	if st.ss[m] == st.ss[m+1] {
		if le.suf == m-sg.left+1 {
			sg.pre += ri.pre
		}
		if ri.pre == sg.right-m {
			sg.suf += le.suf
		}
		sg.mx = max(sg.mx, le.suf+ri.pre)
	}
}

func (st segmentTree) update(root int, idx int) {
	sg := st.tree[root]
	// 为啥是直接返回呢
	if sg.left == sg.right {
		return
	}
	m := (sg.left + sg.right) / 2
	if idx <= m {
		st.update(root*2, idx)
	}
	if idx >= m+1 {
		st.update(root*2+1, idx)
	}
	st.pushUp(root)
}
