package main

import (
	"sort"
)

func main() {

}

// 超时
func successfulPairs2(spells []int, potions []int, success int64) []int {
	m := len(spells)
	ans := make([]int, m)
	// 在这里排序了，
	sort.Ints(potions)
	for i, c := range spells {
		// 生成这个数组很耗时
		po := make([]int, len(potions))
		for j := range potions {
			po[j] = potions[j] * c
		}
		// 不需要再对 po 排序，不然会超时
		ans[i] = search2(po, int(success))
	}
	return ans
}

func search2(potions []int, success int) int {
	n := len(potions)
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if mid < n && potions[mid] >= success {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return n - left
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	m := len(spells)
	ans := make([]int, m)
	// 在这里排序了，
	sort.Ints(potions)
	for i, c := range spells {
		// 不需要再对 po 排序，不然会超时
		ans[i] = search(potions, c, int(success))
	}
	return ans
}

func search(potions []int, c, success int) int {
	n := len(potions)
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if mid < n && potions[mid]*c >= success {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return n - left
}
