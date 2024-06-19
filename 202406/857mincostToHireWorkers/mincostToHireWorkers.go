package main

import (
	"container/heap"
	"fmt"
	"sort"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	fmt.Println(mincostToHireWorkers([]int{10, 20, 5}, []int{70, 50, 30}, 2))
	fmt.Println(mincostToHireWorkers([]int{3, 1, 10, 10, 1}, []int{4, 8, 2, 2, 7}, 3))
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	n := len(quality)
	poeple := make([]pair, n)
	for i := range quality {
		poeple[i] = pair{qu: quality[i], wa: wage[i], r: float64(wage[i]) / float64(quality[i])}
	}
	sort.Slice(poeple, func(i, j int) bool { return poeple[i].r < poeple[j].r })
	sumQ := 0
	mh := make(MaxHeap, 0)
	for i := 0; i < k; i++ {
		heap.Push(&mh, poeple[i].qu)
		sumQ += poeple[i].qu
	}
	ans := float64(sumQ) * poeple[k-1].r
	for j := k; j < n; j++ {
		if poeple[j].qu < mh[0] {
			sumQ = sumQ - mh[0] + poeple[j].qu
			mh[0] = poeple[j].qu
			heap.Fix(&mh, 0)
			ans = min(ans, float64(sumQ)*poeple[j].r)
		}
	}
	return ans
}

type pair struct {
	qu int
	wa int
	r  float64
}
