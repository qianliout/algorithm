package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(squareFreeSubsets([]int{3, 4, 4, 5}))
	fmt.Println(squareFreeSubsets([]int{1}))
}

func squareFreeSubsets(nums []int) int {
	nsq := genMSQ()
	m := 1 << 10
	f := make([]int, m)

	mod := int(math.Pow10(9)) + 7
	f[0] = 1
	for _, ch := range nums {
		mask := nsq[ch]
		if mask < 0 { // 不是nsg，也就是说这个数，不是由平方数构成
			continue
		}
		// 这里是从 m-1 开始，因为 f 数组最大的 index 就是 m-1
		for j := m - 1; j >= mask; j-- {
			// mask是 j 的子集合
			if j|mask == j {
				f[j] = f[j] + f[mask^j]
				f[j] = f[j] % mod
			}
		}
	}
	ans := 0
	for _, ch := range f {
		ans = (ans + ch) % mod
	}

	// 最后要减去空集合
	ans = (ans - 1 + mod) % mod
	return ans
}

func genMSQ() []int {
	pri := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	nsq := make([]int, 31)
	for i := 2; i <= 30; i++ {
		for j, p := range pri {
			if i%p == 0 {
				// 有平方因子
				if i%(p*p) == 0 {
					nsq[i] = -1
					break
				}
				nsq[i] = nsq[i] | (1 << j)
			}
		}
	}
	return nsq
}
