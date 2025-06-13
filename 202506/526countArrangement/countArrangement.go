package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(countArrangement(2))
}

func countArrangement(n int) int {
	u := (1 << n) - 1
	var dfs func(s int) int

	// 按照上面的讨论，定义 dfs(S) 表示在之前选过的数的集合为 S 的情况下，剩余数字可以构造的优美排列的数量。
	dfs = func(s int) int {
		if s == u {
			return 1
		}
		ans := 0
		// i 表示上一次选的数
		// 如果是上一次没有选过，那么
		i := bits.OnesCount(uint(s)) + 1

		for j := 1; j <= n; j++ {
			if s>>(j-1)&1 == 0 && (i%j == 0 || j%i == 0) {
				ans += dfs(s | (1 << (j - 1)))
			}
		}
		return ans
	}
	ans := dfs(0)
	return ans
}

// 1 <= n <= 15
// 对于本题来说，二进制的最低位表示 1，次低位表示 2，依此类推。
