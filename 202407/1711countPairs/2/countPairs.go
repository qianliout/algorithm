package main

import (
	"math"
)

func main() {

}

func countPairs(deliciousness []int) int {
	cnt := make(map[int]int)
	ans := 0
	mod := int(math.Pow10(9)) + 7
	mx := 1 << 22 // 题目中给定的范围
	for _, ch := range deliciousness {
		cnt[ch]++
	}
	for k, v := range cnt {
		for x := 1; x < mx; x = x << 1 {
			t := x - k

			b := cnt[t]

			if k == t {
				ans += (v - 1) * b
			} else {
				ans += v * b
			}
		}
	}
	// 这种方式会对同一对计算两次，所以要除以2
	ans >>= 1

	return ans % mod
}

// 检测是否是 n 的幂
func check(n int) bool {
	// 判断 n 是否大于0，并且 n 的二进制表示中只有一位为1
	return n > 0 && (n&(n-1)) == 0
}
