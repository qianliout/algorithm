package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countSpecialSubsequences([]int{0, 1, 2, 2}))
	fmt.Println(countSpecialSubsequences([]int{0, 1, 2, 0, 1, 2}))
}

func countSpecialSubsequences(nums []int) int {
	mod := int(math.Pow10(9)) + 7
	n := len(nums)
	// f[i][0] 表示前 i 项得到的全 0 子序列个数
	// f[i][1] 表示前 i 项得到的先 0 后 1 的子序列个数
	// f[i][2] 表示前 i 项得到的特殊子序列个数
	//  遍历数组 nums，对于 f[i][j]，若 j!=nums[i]，则直接从前一项转移过来，即 f[i][j]=f[i−1][j]。
	// 若 j=nums[i] 则需要分类计算：
	// 对于 f[i][0]，当遇到 0 时，有选或不选两种方案，不选 0 时有 f[i][0]=f[i−1][0]，选 0 时，可以单独组成一个子序列，也可以与前面的 0 组合，因此有 f[i][0]=f[i−1][0]+1，两者相加得 f[i][0]=2⋅f[i−1][0]+1。
	// 对于 f[i][1]，当遇到 1 时，有选或不选两种方案，不选 1 时有 f[i][1]=f[i−1][1]，选 1 时，可以单独与前面的 0 组成一个子序列，也可以与前面的 1 组合，因此有 f[i][1]=f[i−1][1]+f[i−1][0]，两者相加得 f[i][1]=2⋅f[i−1][1]+f[i−1][0]。
	// f[i][2] 和 f[i][1] 类似，有 f[i][2]=2⋅f[i−1][2]+f[i−1][1]。
	// 参考 链接：https://leetcode.cn/problems/count-number-of-special-subsequences/solutions/908427/dong-tai-gui-hua-by-endlesscheng-4onu/

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, 3)
	}
	for i, ch := range nums {
		for j := 0; j < 3; j++ {
			if ch != j {
				f[i+1][j] = f[i][j]
			} else {
				if j == 0 {
					f[i+1][0] = f[i][0] + (f[i][0] + 1)
				}
				if j == 1 {
					f[i+1][1] = f[i][1] + f[i][1] + f[i][0]
				}
				if j == 2 {
					f[i+1][2] = f[i][2] + f[i][2] + f[i][1]
				}
			}
			f[i+1][j] %= mod
		}
	}
	return f[n][2] % mod
}

func countSpecialSubsequences2(nums []int) int {
	mod := int(math.Pow10(9)) + 7
	n := len(nums)
	// f[i][0] 表示前 i 项得到的全 0 子序列个数
	// f[i][1] 表示前 i 项得到的先 0 后 1 的子序列个数
	// f[i][2] 表示前 i 项得到的特殊子序列个数
	//  遍历数组 nums，对于 f[i][j]，若 j!=nums[i]，则直接从前一项转移过来，即 f[i][j]=f[i−1][j]。
	// 若 j=nums[i] 则需要分类计算：
	// 对于 f[i][0]，当遇到 0 时，有选或不选两种方案，不选 0 时有 f[i][0]=f[i−1][0]，选 0 时，可以单独组成一个子序列，也可以与前面的 0 组合，因此有 f[i][0]=f[i−1][0]+1，两者相加得 f[i][0]=2⋅f[i−1][0]+1。
	// 对于 f[i][1]，当遇到 1 时，有选或不选两种方案，不选 1 时有 f[i][1]=f[i−1][1]，选 1 时，可以单独与前面的 0 组成一个子序列，也可以与前面的 1 组合，因此有 f[i][1]=f[i−1][1]+f[i−1][0]，两者相加得 f[i][1]=2⋅f[i−1][1]+f[i−1][0]。
	// f[i][2] 和 f[i][1] 类似，有 f[i][2]=2⋅f[i−1][2]+f[i−1][1]。
	// 参考 链接：https://leetcode.cn/problems/count-number-of-special-subsequences/solutions/908427/dong-tai-gui-hua-by-endlesscheng-4onu/

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, 3)
	}
	for i, ch := range nums {
		for j := 0; j < 3; j++ {
			if ch != j {
				f[i+1][j] = f[i][j]
			} else {
				if j == 0 {
					f[i+1][0] = f[i][0] + (f[i][0] + 1)
				}
				if j == 1 {
					f[i+1][1] = f[i][1] + f[i][1] + f[i][0]
				}
				if j == 2 {
					f[i+1][2] = f[i][2] + f[i][2] + f[i][1]
				}
			}
			f[i+1][j] %= mod
		}
	}
	return f[n][2] % mod
}
