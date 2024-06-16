package main

import (
	"fmt"
)

func main() {
	fmt.Println(buildArray([]int{1, 3}, 3))
	fmt.Println(buildArray([]int{1, 2}, 4))
}

func buildArray(target []int, n int) []string {
	ans := make([]string, 0)
	start := 1
	for _, num := range target {
		ans = append(ans, h(start, num-1)...)
		ans = append(ans, "Push")
		start = num + 1
	}
	return ans
}

func h(start, end int) []string {
	ans := make([]string, 0)
	if start > end {
		return ans
	}
	n := end - start + 1
	for i := 0; i < n; i++ {
		ans = append(ans, "Push")
	}
	for i := 0; i < n; i++ {
		ans = append(ans, "Pop")
	}
	return ans
}
