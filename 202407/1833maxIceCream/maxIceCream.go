package main

import (
	"sort"
)

func main() {

}

// 你必须使用计数排序解决此问题。
// 这种做法没有使用计数排序，只使用了贪心的做法
func maxIceCream(costs []int, coins int) int {
	n := len(costs)
	sort.Ints(costs)
	ans := 0
	for i := 0; i < n; i++ {
		if coins < costs[i] {
			break
		}
		ans++
		coins -= costs[i]
	}
	return ans
}
