package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumOfFlooredPairs([]int{34912, 57940, 45747}))
	fmt.Println(sumOfFlooredPairs([]int{2, 5, 9}))
}

func sumOfFlooredPairs(nums []int) int {
	n := 100010
	mod := 1000000007
	// 这里开两倍树组是本题目的关键
	cnt := make([]int, n*2)
	for _, ch := range nums {
		cnt[ch]++
	}
	sum := make([]int, n*2+1)
	for i := 0; i < 2*n; i++ {
		sum[i+1] = sum[i] + cnt[i]
	}
	ans := 0
	for i, ch := range cnt {
		if ch == 0 {
			continue
		}
		// 假如i=4,那么此时我们就需要找4*1，4*2，4*3，4*4 等等,比4小的数不用找，因为对答案没有影响
		// 假如n=13,那么找4*4的时候就会越界，但是如果只找到 d*i<=n,那最后一个13又不能加入到答案中去，
		// 所以最简单的办法是，cnt及sum 都扩大一倍，这样就不用考虑边界的问题
		for d := 1; d*i <= n; d++ {

			ans += ch * d * (sum[(d+1)*i] - sum[d*i])
			ans = ans % mod
		}
	}
	return ans % mod
}

// 统计每个数的出现次数 cnt，并计算次数的前缀和 sum。
// 然后枚举分母以及商，用前缀和计算对应的分子个数，将分母个数 c、商 d、分子个数三者相乘即为每部分的结果，累加每部分的结果即为答案。
// nums = [2,5,6]
