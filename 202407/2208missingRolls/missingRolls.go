package main

import (
	"fmt"
)

func main() {
	fmt.Println(missingRolls([]int{1, 2, 3, 4}, 6, 4))
}

func missingRolls(rolls []int, mean int, n int) []int {

	m := len(rolls)
	all := mean * (m + n)
	sum := 0
	for _, ch := range rolls {
		sum += ch
	}
	sub := all - sum
	if sub < 1*n || sub > 6*n {
		return []int{}
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = sub / n
	}
	for i := 0; i < sub%n; i++ {
		ans[i]++
	}
	return ans
}
