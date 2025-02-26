package main

import (
	"fmt"
)

func main() {
	fmt.Println(crackNumber(216612))
}

func crackNumber(ciphertext int) int {
	cip := fmt.Sprintf("%d", ciphertext)
	n := len(cip)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = f[i-1] // 新组一成
		if (cip[i-1] <= '5' && (i-2 >= 0 && cip[i-2] >= '1' && cip[i-2] <= '2')) || (cip[i-1] > '5' && (i-2 >= 0 && cip[i-2] == '1')) {
			f[i] += f[i-2]
		}
	}
	return f[n]
}

// 题目中说了是合法的

func crackNumber2(ciphertext int) int {
	return 0
}
