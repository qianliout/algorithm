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
	// 表示后面两个数分别是 last1和last2，前面还有 i 个数没有填
	dfs = func(i int, last1, last2 int) int {
		if i == n {
			return 1
		}
		if mem[i][last1][last2] != -1 {
			return mem[i][last1][last2]
		}
		res := 0
		for j := 1; j <= 6; j++ {
			// 条件1
			if gcd(j, last1) != 1 {
				continue
			}
			// 条件2
			if j == last1 || j == last2 {
				continue
			}
			res = (res + dfs(i+1, j, last1)) % mod
		}
		mem[i][last1][last2] = res % mod
		return res
	}
	return dfs(0, 7, 7) % mod
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
