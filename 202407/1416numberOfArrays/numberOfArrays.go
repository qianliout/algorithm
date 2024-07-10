package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numberOfArrays("1000", 1000))

}

func numberOfArrays(s string, k int) int {
	n := len(s)
	// dp[i] 表示前 i 个数字进行恢复的方案数，那么可以很容易地写出递推式：
	dp := make([]int, n+1)
	dp[0] = 1
	mod := int(math.Pow10(9)) + 7

	for i := 1; i <= n; i++ {
		num, base := 0, 1
		// 题目中说了 k 的数据范围
		for j := i - 1; j >= 0 && i-j <= 10; j-- {
			num = int(s[j]-'0')*base + num
			if num > k {
				break
			}
			// 不能这样写，如果碰到0，还是可以向前走的
			// if s[j]=='0'{
			// 	break
			// }

			// 不能有前导0
			if s[j] != '0' {
				dp[i] += dp[j]
			}
			base = base * 10
		}
		dp[i] = dp[i] % mod
	}
	return dp[n]
}
