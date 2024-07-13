package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsPattern([]int{1, 2, 3, 1, 2}, 2, 2))
}

// 连续 重复多次但 不重叠
func containsPattern(arr []int, m int, k int) bool {
	n := len(arr)

	for i := 0; i <= n-m; i++ {
		if check(arr, i, k, arr[i:i+m]) {
			return true
		}
	}
	return false
}

// 连续 重复多次但 不重叠
func check(arr []int, start, k int, sub []int) bool {
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, sub...)
	}
	end := min(len(arr), start+len(ans))
	return same(arr[start:end], ans)
}

func same(a, b []int) bool {
	for len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
