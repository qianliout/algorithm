package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAddedCoins([]int{1, 4, 10, 5, 7, 19}, 19))
}

func minimumAddedCoins(coins []int, target int) int {
	sort.Ints(coins)
	ans := 0
	i, s := 0, 1
	for s <= target {
		if i < len(coins) && coins[i] <= s {
			s += coins[i]
			i++
		} else {
			// 此时就必须添加数了
			s *= 2
			ans++
		}
	}

	return ans
}
