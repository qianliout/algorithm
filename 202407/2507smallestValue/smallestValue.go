package main

import (
	"fmt"
)

func main() {
	// fmt.Println(smallestValue(15))
	fmt.Println(smallestValue(12))
	fmt.Println(smallestValue(4))
}

// smallestValue 寻找小于或等于 n 的最小自然数，该数的质因数分解之和等于 n。
// 如果不存在这样的数，则返回 n 本身。
// 参数:
//
//	n - 输入的整数
//
// 返回值:
//
//	最小满足条件的自然数
func smallestValue(n int) int {
	// 不断迭代直到找到满足条件的最小自然数
	for {
		// 初始化 ans 为 0，用于累加质因数的和
		ans := 0
		// x 用于临时存储 n，以便在循环中进行操作
		x := n
		// 从 2 开始遍历到 sqrt(x)，寻找 x 的质因数
		for i := 2; i*i <= x; i++ {
			// 当 x 能被 i 整除时，累加 i 到 ans，并更新 x
			for ; x%i == 0; x = x / i {
				ans += i
			}
		}
		// 如果 x 大于 1，说明 x 是个质数,没有除尽，将其加到 ans
		if x > 1 {
			ans += x
		}
		// 如果 ans 的质因数和等于 n，返回 n
		if ans == n {
			return n
		}
		// 更新 n 为 ans，继续下一轮迭代
		n = ans
	}
}
