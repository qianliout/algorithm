package main

import (
	"sort"
)

func main() {

}

func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	n := len(arr)
	mid := arr[(n-1)/2]
	ans := make([]int, 0)
	l, r := 0, n-1
	for k > 0 && l <= r {
		a := abs(mid - arr[l])
		b := abs(mid - arr[r])
		if a > b {
			ans = append(ans, arr[l])
			l++
		} else {
			ans = append(ans, arr[r])
			r--
		}
		k--
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
