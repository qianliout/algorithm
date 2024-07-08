package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getKth(7, 11, 4))
}

func getKth(lo int, hi int, k int) int {
	if k > (hi - lo + 1) {
		return 0
	}
	mem := make(map[int]int)
	ans := make([]int, 0)
	for i := lo; i <= hi; i++ {
		ans = append(ans, i)
	}
	sort.Slice(ans, func(i, j int) bool {
		a := cal(ans[i], mem)
		b := cal(ans[j], mem)
		if a < b {
			return true
		} else if a > b {
			return false
		}
		return ans[i] < ans[j]
	})

	return ans[k-1]
}
func cal(n int, mem map[int]int) int {
	if n == 1 {
		return 0
	}
	if mem[n] > 0 {
		return mem[n]
	}
	if n&1 == 0 {
		res := cal(n/2, mem) + 1
		mem[n] = res
		return res
	}

	res := cal(3*n+1, mem) + 1
	mem[n] = res
	return res
}
