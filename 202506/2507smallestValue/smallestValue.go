package main

import (
	"fmt"
)

func main() {
	n := 5
	for i := 0; i < n; i++ {
		n--
		fmt.Println("n", n)
	}

}

func smallestValue(n int) int {
	for {
		x := n
		ans := count(x)
		if ans == x {
			return ans
		}
		n = ans
	}
}

// count 计算一个数的所有质因数之和（包含重复的质因数）
// 例如：count(12) = 2 + 2 + 3 = 7 (因为 12 = 2² × 3)
// 例如：count(15) = 3 + 5 = 8 (因为 15 = 3 × 5)
func count(n int) int {
	ans := 0 // 存储质因数之和

	// 从2开始检查所有可能的质因数，只需要检查到 sqrt(n)
	// 因为如果 n 有大于 sqrt(n) 的因数，那么必然有一个小于 sqrt(n) 的对应因数
	for i := 2; i*i <= n; i++ {
		// 内层循环：如果 i 是 n 的因数，就不断除以 i，直到不能整除为止
		// 这样可以找到 i 这个质因数在 n 中出现的所有次数
		for ; n%i == 0; n = n / i {
			ans += i // 每找到一个质因数 i，就加到答案中
		}
		// 循环结束后，n 中所有的因数 i 都被除掉了
	}

	// 如果经过上面的处理后 n > 1，说明 n 本身就是一个大于 sqrt(原始n) 的质数
	// 这个质数也需要加到答案中
	if n > 1 {
		ans += n
	}

	return ans
}

/*
给你一个正整数 n 。
请你将 n 的值替换为 n 的 质因数 之和，重复这一过程。
注意，如果 n 能够被某个质因数多次整除，则在求和时，应当包含这个质因数同样次数。
返回 n 可以取到的最小值。
*/

// 判断一个数是否是质数
func check(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
