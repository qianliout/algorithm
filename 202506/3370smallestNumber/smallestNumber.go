package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(smallestNumber(5))
}

func smallestNumber2(n int) int {
	i := n
	for {
		if check(i) {
			return i
		}
		i++
	}
}

func check(a int) bool {
	for a > 0 {
		if a&1 == 0 {
			return false
		}
		a = a >> 1
	}
	return true
}

func smallestNumber(n int) int {
	l := bits.Len(uint(n))
	ans := 1<<l - 1
	if ans < n {
		return 1<<(l+1) - 1
	}
	return ans
}

func minChanges2(n int, k int) int {
	ans := 0
	for n > 0 || k > 0 {
		if n&1 == 0 && k&1 == 1 {
			return -1
		}
		if n&1 == 1 && k&1 == 0 {
			ans++
		}
		n = n >> 1
		k = k >> 1
	}
	return ans
}

// minChanges 计算将 n 变成 k 所需的最少操作次数
// 只能将 1 改为 0，不能将 0 改为 1
func minChanges(n int, k int) int {
	// 如果 n 和 k 相等，不需要任何操作
	if n == k {
		return 0
	}

	// 关键判断：检查是否可能从 n 变成 k
	// 如果 k 的某一位是 1，而 n 的对应位是 0，则无法实现转换
	// 因为我们只能将 1 改为 0，不能将 0 改为 1
	// 等价于检查：k 是否是 n 的子集（在位运算意义下）
	if (n & k) != k {
		return -1
	}

	// 计算需要改变的位数
	// n ^ k 会得到所有不同的位
	// 由于我们已经确保 k 是 n 的子集，所以 n ^ k 中的 1 都是需要从 n 中移除的位
	return bits.OnesCount(uint(n ^ k))
}

/*
给你两个正整数 n 和 k。
你可以选择 n 的 二进制表示 中任意一个值为 1 的位，并将其改为 0。
返回使得 n 等于 k 所需要的更改次数。如果无法实现，返回 -1。
*/
