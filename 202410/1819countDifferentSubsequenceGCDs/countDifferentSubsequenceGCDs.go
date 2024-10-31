package main

import (
	"slices"
)

func main() {

}

// countDifferentSubsequenceGCDs 计算数组 nums 中不同子序列的最大公约数 (GCD) 的数量。
// 它通过遍历从 1 到数组中的最大值，检查每个可能的 GCD 值是否可以通过数组中的某个子序列获得。
func countDifferentSubsequenceGCDs(nums []int) int {
	// 使用 map 记录数组中出现过的数字，以快速检查某个数字是否存在于数组中。
	exit := make(map[int]bool)
	for _, ch := range nums {
		exit[ch] = true
	}

	// ans 用于记录不同子序列的最大公约数的数量。
	ans := 0
	// 找出数组中的最大值，以确定检查的上限。
	mx := slices.Max(nums)

	// 遍历从 1 到最大值，检查每个数字是否可以作为某个子序列的最大公约数。
	for i := 1; i <= mx; i++ {
		// g 用于记录当前子序列的最大公约数。
		g := 0

		// 遍历数组中的每个数字，步长为 i，以检查是否可以形成最大公约数为 i 的子序列。
		for j := i; j <= mx; j += i {
			// 如果当前数字不在数组中，则跳过。
			if !exit[j] {
				continue
			}

			// 更新当前子序列的最大公约数。
			g = gcd(g, j)

			// 如果当前子序列的最大公约数等于 i，则说明找到了一个新的不同子序列的最大公约数。
			if g == i {
				ans++
				// 找到后即可停止当前子序列的检查，继续检查下一个可能的最大公约数。
				break
			}
		}
	}

	// 返回不同子序列的最大公约数的数量。
	return ans
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
