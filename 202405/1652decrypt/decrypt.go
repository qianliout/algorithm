package main

import (
	"fmt"
)

func main() {
	fmt.Println(decrypt([]int{5, 7, 1, 4}, 3))
	fmt.Println(decrypt([]int{2, 4, 9, 3}, -2)) // [12,5,6,13]
}

func decrypt(code []int, k int) []int {
	n := len(code)
	sum := make([]int, len(code)*2+1)
	for i := 0; i < len(code); i++ {
		sum[i+1] = sum[i] + code[i]
	}

	for i := 0; i < len(code); i++ {
		sum[i+n+1] = sum[i+n] + code[i]
	}
	ans := make([]int, n)
	if k == 0 {
		return ans
	}

	for i := 0; i < n; i++ {
		le := i + 1
		if k < 0 {
			le = i + n
		}
		ri := le + k
		if k < 0 {
			le, ri = ri, le
		}

		ans[i] = sum[ri] - sum[le]
	}
	return ans
}
