package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(minimumXORSum([]int{1, 2}, []int{2, 3}))
}

func minimumXORSum(nums1 []int, nums2 []int) int {
	n := len(nums1)
	mask := 1 << n
	var inf int = 0x3f3f3f3f
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, mask)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for s := 0; s < mask; s++ {
			// if getCnt(s, n) != i {
			// 	continue
			// }
			// 由于总共考虑了前 i个成员，因此 s中 111 的数量必然为 i否则 f[i] 就不是一个合法状态，跳过转移
			if bits.OnesCount(uint(s)) != i {
				continue
			}
			for j := 1; j <= n; j++ {
				if (s>>(j-1))&1 == 0 {
					continue
				}
				dp[i][s] = min(dp[i][s], dp[i-1][s^(1<<(j-1))]+(nums1[i-1]^nums2[j-1]))
			}
		}
	}
	return dp[n][mask-1]
}

/*

为了方便，我们令下标从 111 开始。
定义 f[i][s]f[i][s]f[i][s] 为考虑前 iii 个元素，且对 nums2 的使用情况为 sss 时的最小异或值。其中 sss 是一个长度为 n的二进制数：若 s 中的第 k 位为 1，说明 nums2[k] 已被使用；若 s 中的第 k 位为 0说明 nums2[k] 未被使用。
起始时，只有 f[0]=0，其余均为无穷大 INF。f[0]含义为在不考虑任何数，对 nums2 没有任何占用情况时，最小异或值为 0。最终 f[n][2n−1] 即为答案。
不失一般性考虑 f[i][s]f[i][s]f[i][s] 该如何转移，可以以 nums1[i] 是与哪个 nums2[j] 进行配对作为切入点：
    由于总共考虑了前 i个成员，因此 s中 111 的数量必然为 i否则 f[i] 就不是一个合法状态，跳过转移
    枚举 nums1[i] 是与哪一个 nums2[j] 进行配对的，且枚举的 j 需满足在 sss 中的第 j 位值为 1，若满足则有

f[i][s]=min(f[i][s],f[i−1][prev]+nums1[i]^nums2[j])
其中 prev 为将 s 中的第 j 位进行置零后的二进制数，即 prev = s ^ (1 << j)，符号 ^ 代表异或操作
*/

func getCnt(s, n int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans += (s >> i) & 1
	}
	return ans
}
