package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(maxAbsValExpr([]int{1, -2}, []int{8, 8}))
	fmt.Println(mctFromLeafValues([]int{7, 12, 8, 10}))
}

func mctFromLeafValues(arr []int) int {
	ans := 0
	st := make([]int, 0)
	st = append(st, math.MaxInt/10)
	for _, ch := range arr {
		for ch >= st[len(st)-1] {
			pop := st[len(st)-1]
			st = st[:len(st)-1]
			ans += pop * min(st[len(st)-1], ch)
		}
		st = append(st, ch)
	}

	// for len(st) > 2 {
	// 	pop := st[len(st)-1]
	// 	st = st[:len(st)-1]
	// 	ans += pop * st[len(st)-1]
	// }

	for i := len(st) - 1; i > 1; i-- {
		ans += st[i] * st[i-1]
	}
	return ans
}

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	dirs := [][]int{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}}
	n := len(arr1)
	ans := 0
	for _, dir := range dirs {
		mx, mi := math.MinInt/100, math.MaxInt/100
		for i := 0; i < n; i++ {
			mx = max(mx, arr1[i]*dir[0]+arr2[i]*dir[1]+i)
			mi = min(mi, arr1[i]*dir[0]+arr2[i]*dir[1]+i)
		}
		ans = max(ans, mx-mi)
	}

	return ans
}
