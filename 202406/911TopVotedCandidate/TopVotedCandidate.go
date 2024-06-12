package main

import (
	"sort"
)

func main() {

}

type TopVotedCandidate struct {
	win []int // 记录每个时刻的获胜者
	tm  []int // 记录投票时间，共下Q方法查询
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n, mxC, mxP := len(persons), 0, 0
	win, sum := make([]int, n), make([]int, n+1)
	for i, p := range persons {
		sum[p]++
		if sum[p] >= mxC {
			mxC, mxP = sum[p], p
		}
		win[i] = mxP
	}
	return TopVotedCandidate{win: win, tm: times}
}

func (this *TopVotedCandidate) Q(t int) int {
	/*
		// 找<=t 的最大值，也就是左端点,这种写法也是对的
		j := sort.SearchInts(this.tm, t+1)
		if j <= 0 {
			return -1
		}
		return this.win[j-1]
	*/
	j := sort.Search(len(this.tm), func(k int) bool { return this.tm[k] > t })
	if j <= 0 {
		return -1
	}
	return this.win[j-1]
}
