package main

import (
	"sort"
)

func main() {

}

func minCost(nums []int, cost []int) int64 {
	n := len(nums)
	pp := make([]pair, n)
	sumCost := 0
	for i := 0; i < n; i++ {
		pp[i] = pair{x: nums[i], cost: cost[i]}
		sumCost += cost[i]
	}
	sort.Slice(pp, func(i, j int) bool { return pp[i].x < pp[j].x })
	// 先计算把所有的数都改成nums[0]
	total := 0
	for i := 1; i < n; i++ {
		total += pp[i].x - pp[0].x*pp[i].cost
	}
	// 再计算把所有数改成其他数的情况
	// 然后考虑要使所有元素都等于 nums[1]，sumCost 的变化量是多少：
	//    有 cost[0] 这么多的开销要增加 nums[1]−nums[0]；
	//    有 sumCost−cost[0] 这么多的开销要减少 nums[1]−nums[0]。
	mi := total
	for i := 1; i < n; i++ {
		total = total + pp[i-1].cost*(pp[i].x-pp[i-1].x)
		total = total - (sumCost-pp[i-1].cost)*(pp[i].x-pp[i-1].x)
		sumCost = sumCost - pp[i-1].cost*2
		mi = min(mi, total)
	}
	return int64(mi)
}

type pair struct {
	x, cost int
}
