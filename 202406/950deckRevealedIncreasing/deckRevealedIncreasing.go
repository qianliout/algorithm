package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
}

// 每个值都不一样
// 模拟题目
func deckRevealedIncreasing(deck []int) []int {
	// deck 里只是牌的值，是什么顺序不重要
	tem := make([]int, 0)
	dq := make([]int, 0)
	tem = append(tem, deck...)
	dq = append(dq, deck...)
	n := len(deck)
	i := 0
	// 表示每个值的显示顺序
	idx := make(map[int]int)
	for len(dq) > 0 {
		fir := dq[0]
		idx[fir] = i
		i++
		dq = dq[1:]
		if len(dq) > 0 {
			se := dq[0]
			dq = dq[1:]
			dq = append(dq, se)
		}
	}
	// 题目要求递增
	// 这里是最难理解的，
	// 要求的返回是 ans, ans 按题目要求的展示方法展示之后的数组是排序后的 tmm（也就是题目中说的递增）
	// 那么ans[j]，表示第j 个元素，那他在 tmm 中是第几个元素呢：那就是 idx[deck[j]],那这个元素的值就是tem[idx[deck[j]]]
	sort.Ints(tem)
	ans := make([]int, n)
	for j := 0; j < n; j++ {
		ans[j] = tem[idx[deck[j]]]
	}
	return ans
}
