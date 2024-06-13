package main

import (
	"fmt"
	"slices"
)

func main() {
	com := [][]int{{2, 1}, {1, 2}, {1, 1}}
	stok := []int{1, 1}
	cost := []int{5, 5}
	fmt.Println(maxNumberOfAlloys(2, 3, 10, com, stok, cost))
}

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	var check func(mx int) bool

	check = func(m int) bool {
		for j := range composition {
			com := composition[j]
			ans := 0
			for i := 0; i < len(com); i++ {
				ans += max(0, com[i]*m-stock[i]) * cost[i]
			}
			if ans <= budget {
				return true
			}
		}
		return false
	}

	mx := slices.Min(stock) + budget
	// j := sort.Search(mx, check)
	// return j
	// 查右端点逻辑，最好自已实现
	le := 0
	ri := mx
	for le < ri {
		mid := le + (ri-le+1)/2
		if mid >= 0 && mid < mx && check(mid) {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	return le
}
