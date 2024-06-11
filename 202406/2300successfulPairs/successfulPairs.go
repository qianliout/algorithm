package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	pn := len(potions)
	ans := make([]int, len(spells))
	for i, ch := range spells {
		le, ri := 0, pn
		for le < ri {
			mid := le + (ri-le)/2
			// 寻找大于等于的左边界
			if mid >= 0 && mid < pn && int64(ch*potions[mid]) >= success {
				ri = mid
			} else {
				le = mid + 1
			}
		}
		ans[i] = pn - le
	}
	return ans
}
