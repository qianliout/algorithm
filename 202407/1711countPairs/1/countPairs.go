package main

import (
	"math"
)

func main() {

}

// 会超时
func countPairs(deliciousness []int) int {
	cnt := make(map[int]int)
	ans := 0
	mod := int(math.Pow10(9)) + 7
	for _, ch := range deliciousness {
		for k, v := range cnt {
			if check(k + ch) {
				ans += v
			}
		}
		cnt[ch]++
	}
	return ans % mod
}

// 检测是否是 n 的幂
func check(n int) bool {
	// 判断 n 是否大于0，并且 n 的二进制表示中只有一位为1
	return n > 0 && (n&(n-1)) == 0
}
