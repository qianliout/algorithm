package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(strangePrinter("aaabbb"))
	fmt.Println(strangePrinter("aba"))
}

/*
定义 f[l][r]f[l][r]f[l][r] 为将 [l,r][l, r][l,r] 这一段打印成目标结果所消耗的最小打印次数。

不失一般性考虑 f[l][r]f[l][r]f[l][r] 该如何转移：

只打印 lll 这个位置，此时 f[l][r]=f[l+1][r]+1f[l][r] = f[l + 1][r] + 1f[l][r]=f[l+1][r]+1
不只打印 lll 这个位置，而是从 lll 打印到 kkk（需要确保首位相同 s[l]=s[k]s[l] = s[k]s[l]=s[k]）：f[l][r]=f[l][k−1]+f[k+1][r],l<k<=rf[l][r] = f[l][k - 1] + f[k + 1][r], l < k <= rf[l][r]=f[l][k−1]+f[k+1][r],l<k<=r
其中状态转移方程中的情况 222 需要说明一下：由于我们只确保 s[l]=s[k]s[l] = s[k]s[l]=s[k]，并不确保 [l,k][l, k][l,k] 之间的字符相同，根据我们基本分析可知，s[k]s[k]s[k] 这个点可由打印 s[l]s[l]s[l] 的时候一同打印，因此本身 s[k]s[k]s[k] 并不独立消耗打印次数，所以这时候 [l,k][l, k][l,k] 这一段的最小打印次数应该取 f[l][k−1]f[l][k - 1]f[l][k−1]，而不是 f[l][k]f[l][k]f[l][k]。

最终的 f[l][r]f[l][r]f[l][r] 为上述所有方案中取 minminmin。

作者：宫水三叶
链接：https://leetcode.cn/problems/strange-printer/solutions/792628/gong-shui-san-xie-noxiang-xin-ke-xue-xi-xqeo9/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func strangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for l := n - 1; l >= 0; l-- {
		dp[l][l] = 1 // 只打印自己
		for r := l + 1; r < n; r++ {
			if s[l] == s[r] {
				dp[l][r] = dp[l][r-1]
				continue
			}
			dp[l][r] = math.MaxInt
			for k := l; k < r; k++ {
				dp[l][r] = min(dp[l][r], dp[l][k]+dp[k+1][r])
			}
		}
	}
	return dp[0][n-1]
}
