package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(longestAwesome("3242415"))
	fmt.Println(longestAwesome("12345678"))
	fmt.Println(longestAwesome("213123"))
	fmt.Println(longestAwesome("373781"))

}

// 进行了状态压缩，但是时间复杂度还是n平方，所以还是会 timeout
func longestAwesome2(s string) int {
	n := len(s)
	sum := make([]uint32, n+1)
	for i, ch := range s {
		sum[i+1] = sum[i] ^ (1 << int(ch-'0'))
	}

	ans := 0
	for i := 1; i < len(sum); i++ {
		for j := 0; j < i; j++ {
			if (i-j)&1 == 0 {
				if sum[j]^sum[i] == 0 {
					ans = max(ans, i-j)
				}
			} else {
				if bits.OnesCount32(sum[i]^sum[j]) == 1 {
					ans = max(ans, i-j)
				}
			}
		}
	}
	return ans
}

// 只关注各个前缀第一次出现的位置，用 hash 实现 O（1）
func longestAwesome3(s string) int {
	n := len(s)
	sum := make([]uint32, n+1)
	for i, ch := range s {
		sum[i+1] = sum[i] ^ (1 << int(ch-'0'))
	}
	if (n&1 == 0 && bits.OnesCount32(sum[n]) == 0) || (n&1 == 1 && bits.OnesCount32(sum[n]) == 1) {
		return n
	}

	ans := 0
	last := make(map[uint32]int)

	for i := 0; i < len(sum); i++ {
		ch := sum[i]
		// 奇数
		for j := 0; j <= 9; j++ {
			pre := ch ^ (1 << j)
			if last[pre] > 0 {
				ans = max(ans, i-(last[pre]-1))
			}
		}

		// 偶数
		if last[ch] > 0 {
			ans = max(ans, i-(last[ch]-1))
		}

		if last[ch] == 0 {
			last[ch] = i + 1 // 这里为啥是 i+1,防止和初值0冲突
		}
	}
	return ans
}

// 只关注各个前缀第一次出现的位置，用 hash 实现 O（1）
//
//	可以不开 sum 数组，在更新时计算
func longestAwesome(s string) int {
	n := len(s)
	// sum := make([]uint32, n+1)
	// for i, ch := range s {
	// 	sum[i+1] = sum[i] ^ (1 << int(ch-'0'))
	// }
	// if (n&1 == 0 && bits.OnesCount32(sum[n]) == 0) || (n&1 == 1 && bits.OnesCount32(sum[n]) == 1) {
	// 	return n
	// }

	ans := 0
	last := make(map[int]int)
	last[0] = 0

	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum ^ (1 << int(s[i-1]-'0'))
		// 奇数
		for j := 0; j <= 9; j++ {
			pre := sum ^ (1 << j)
			if preI, ok := last[pre]; ok {
				ans = max(ans, i-(preI))
			}
		}

		// 偶数
		if preI, ok := last[sum]; ok {
			ans = max(ans, i-(preI))
		}

		if _, ok := last[sum]; !ok {
			last[sum] = i
		}
	}
	return ans
}
