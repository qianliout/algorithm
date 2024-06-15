package main

import (
	"fmt"
)

func main() {
	fmt.Println(doesValidArrayExist([]int{1, 1, 0}))
	fmt.Println(doesValidArrayExist([]int{1, 1}))
	fmt.Println(doesValidArrayExist([]int{1, 0}))
}

func circularGameLosers(n int, k int) []int {
	d := k
	visit := make([]bool, n)
	start := 0
	for !visit[start] {
		visit[start] = true
		start = (start + d) % n
		d = d + k
	}

	ans := make([]int, 0)
	for i, v := range visit {
		if !v {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func doesValidArrayExist(derived []int) bool {
	ans := 0
	for i := 0; i < len(derived); i++ {
		ans = ans ^ derived[i]
	}
	return ans == 0
}
