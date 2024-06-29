package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumNumbers(58, 9))
	fmt.Println(minimumNumbers(0, 7))
	fmt.Println(minimumNumbers(10, 1))
}

var base = math.MaxInt / 100

func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	if k == 0 {
		if num%10 != 0 {
			return -1
		}
		return 1
	}

	// if !check(num, k) {
	// 	return -1
	// }
	var dfs func(n int) int

	mem := make([]int, num+10)
	for i := range mem {
		mem[i] = -1
	}

	dfs = func(n int) int {
		if n == 0 {
			return 0
		}
		if mem[n] != -1 {
			return mem[n]
		}
		res := base
		for i := k; i <= n; i = i + 10 {
			nex := dfs(n - i)
			if nex <= base {
				res = min(res, dfs(n-i)+1)
			}
		}
		mem[n] = res
		return res
	}
	a := dfs(num)
	if a >= base {
		return -1
	}
	return a
}

func check(nums int, k int) bool {

	n := nums % 10
	for i := 1; i <= 9; i++ {
		if (k*i)%10 == n {
			return true
		}
	}
	return false
}
