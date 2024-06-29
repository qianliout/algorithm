package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(distinctSequences(4))
	fmt.Println(distinctSequences(2))
	fmt.Println(distinctSequences(20))
}

func distinctSequences(n int) int {
	var dfs func(i int, last1, last2 int) int
	mod := int(math.Pow10(9)) + 7
	mem := make([][][]int, n+1)
	for i := range mem {
		mem[i] = make([][]int, 8)
		for j := range mem[i] {
			mem[i][j] = make([]int, 8)
			for k := range mem[i][j] {
				mem[i][j][k] = -1
			}
		}
	}
	// last2,last1,x 求 此时 x的方案数
	// i表示后面还剩下多秒个数，当i==0时，说明抵达终点了，
	dfs = func(i int, last1, last2 int) int {
		if i == 0 { // 剩下的长度是0了，表示找到了一个合法的方案
			return 1
		}
		if mem[i][last1][last2] != -1 {
			return mem[i][last1][last2]
		}
		res := 0
		for j := 1; j <= 6; j++ {
			if j != last1 && j != last2 && gcd(j, last1) == 1 {
				res = (res + dfs(i-1, j, last1)) % mod
			}

			// // 条件1
			// if gcd(j, last1) != 1 {
			// 	continue
			// }
			// // 条件2
			// if j == last1 || j == last2 {
			// 	continue
			// }
			// res = (res + dfs(i-1, j, last1)) % mod
		}
		mem[i][last1][last2] = res % mod
		return res
	}
	// 为啥可以填呢,因为，7 不在1-6里，且7和1-6之间的每个数的 gcd 都是1，也就是满足:
	// res = (res + dfs(i-1, j, last1)) % mod

	return dfs(n, 7, 7) % mod
}

// // 最大公约数
// func gcd(a, b int) int {
// 	for a != 0 {
// 		a, b = b%a, a
// 	}
// 	return b
//
// }

// 求最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return 0
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
