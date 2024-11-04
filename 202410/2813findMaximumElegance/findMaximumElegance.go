package main

import (
	"slices"
)

func main() {

}

// 反悔贪心的做法

func findMaximumElegance(items [][]int, k int) int64 {
	// 先按利润排序，最先选出利润最大的 k 个item
	slices.SortFunc(items, func(e []int, e2 []int) int { return e2[0] - e[0] })
	ans := 0
	visit := make(map[int]bool)
	// 因为是按利润排序了，dup 里最后一个数，一定是利润最小的
	dup := make([]int, 0)
	totalP := 0
	for i, ch := range items {
		p, c := ch[0], ch[1]
		if i < k {
			totalP += p
			if !visit[c] {
				visit[c] = true
			} else {
				// 如果已访问过了，就加入到重复队列中去
				dup = append(dup, p)
			}
			// 如果之前没有被加入过，那这个数加入后，类别数会增加，所以答案可能变大
		} else if len(dup) > 0 && !visit[c] {
			visit[c] = true
			totalP += p - dup[len(dup)-1]
			dup = dup[:len(dup)-1]
		} else {
			// 比前面的利润小，而且类别还重复了，选它只会让 totalProfit 变小，len(vis) 不变，优雅度不会变大
			// 所以可以不用处理
		}
		ans = max(ans, totalP+len(visit)*len(visit))
	}

	return int64(ans)
}
