package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numberOfStableArrays(3, 3, 2))
	fmt.Println(numberOfStableArrays(13, 20, 93))
	fmt.Println(numberOfStableArrays(31, 36, 60))
	fmt.Println(numberOfStableArrays(35, 35, 22))
	fmt.Println(numberOfStableArrays(677, 458, 614))
}

func numberOfStableArrays(zero int, one int, limit int) int {
	mod := int(math.Pow10(9)) + 7
	// 成定义 dfs(i,j,k) 表示用 i 个 0 和 j 个 1 构造稳定数组的方案数，其中第 i+j 个位置要填 k，其中 k 为 0 或 1。
	var dfs func(i, j, k int) int
	// 考虑 dfs(i,j,0) 怎么算。现在，第 i+j 个位置填的是 0，考虑第 i+j−1 个位置要填什么：
	//    填 0，方案数就是 dfs(i−1,j,0)。 // 因为i+j 的位置已经填了0了，所以还剩下 i-1个0
	//    填 1，方案数就是 dfs(i−1,j,1)。
	// limit的意思是最多有limit 个连续的0或连续的1, 但是加上 Limit 的影响，上面会有部分不合法的数
	// 假如 limit=3 那么最后最多有3个连续的0。
	// dfs(i−4,j,1),这个表示：在 i-4+j 这个位置填1，那么 i-4+j 后面的的4个位置就只能填0，这样是不合法的，要减去
	// 所以 dfs(i,j,0) = dfs(i-1,j,0)+dfs(i-1,j,1)-dfs(i-limit-1,j,1)
	// 对于最后一个位置写1是同理的
	// dfs(i,j,1) = dfs(i,j-1,0)+dfs(i,j-1,1)-dfs(i,j-limit-1,0)
	mem := make(map[string]int)
	dfs = func(i, j, k int) int {
		// 这种方式效率不高，会超时
		key := fmt.Sprintf("%d-%d-%d", i, j, k)
		if va, ok := mem[key]; ok {
			return va
		}
		if i == 0 {
			ans := 0
			if k == 1 && j <= limit {
				ans = 1
			}
			mem[key] = ans
			return ans
		}
		if j == 0 {
			ans := 0
			if k == 0 && i <= limit {
				ans = 1
			}
			mem[key] = ans
			return ans
		}
		if k == 0 {
			ans := dfs(i-1, j, 0) + dfs(i-1, j, 1)
			if i > limit {
				ans -= dfs(i-limit-1, j, 1)
			}
			ans = (ans%mod + mod) % mod
			mem[key] = ans
			return ans
		}
		if k == 1 {
			ans := dfs(i, j-1, 0) + dfs(i, j-1, 1)
			if j > limit {
				ans -= dfs(i, j-limit-1, 0)
			}
			ans = (ans%mod + mod) % mod
			mem[key] = ans
			return ans
		}
		return 0
	}
	ans := dfs(zero, one, 0) % mod
	ans += dfs(zero, one, 1) % mod

	ans = (ans%mod + mod) % mod

	return ans
}
