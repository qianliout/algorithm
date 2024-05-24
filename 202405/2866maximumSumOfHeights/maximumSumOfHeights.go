package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumSumOfHeights([]int{5, 3, 4, 1, 1}))
	fmt.Println(maximumSumOfHeights([]int{6, 5, 3, 9, 2, 7}))
}

func maximumSumOfHeights(maxHeights []int) int64 {
	h := maxHeights

	st := make([]int, 0)
	st = append(st, len(h))
	sufs := make([]int, len(h))

	// 美丽山脉右边部分是递减的，所以从后向前就应该是单调递增的
	// 也就是说x是做最大值加入栈的
	// 先维护后缀
	suf := 0
	for i := len(h) - 1; i >= 0; i-- {
		// maintain
		for len(st) > 1 && h[st[len(st)-1]] >= h[i] {
			le := st[len(st)-1]
			top := h[le]
			st = st[:len(st)-1]
			ri := st[len(st)-1]
			suf -= top * (ri - le)
		}
		// add
		suf += h[i] * (st[len(st)-1] - i)
		sufs[i] = suf
		// push
		st = append(st, i)
	}

	// 再维护前缀
	st = make([]int, 0)
	st = append(st, -1)

	res := 0
	pre := 0
	for i := 0; i < len(h); i++ {
		// maintain
		for len(st) > 1 && h[st[len(st)-1]] >= h[i] {
			ri := st[len(st)-1]
			top := h[ri]
			st = st[:len(st)-1]
			le := st[len(st)-1]
			pre -= top * (ri - le)
		}
		// add
		pre += h[i] * (i - st[len(st)-1])
		res = max(res, pre+sufs[i]-h[i])
		// push
		st = append(st, i)
	}

	return int64(res)
}

// https://leetcode.cn/problems/beautiful-towers-i/solutions/2617548/javapython3cqian-hou-zhui-he-dan-diao-zh-vcvo/
