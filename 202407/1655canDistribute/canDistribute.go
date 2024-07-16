package main

import (
	"fmt"
)

func main() {
	fmt.Println(canDistribute([]int{1, 2, 3, 4}, []int{2}))
}

func canDistribute(nums []int, quantity []int) bool {
	m := 1 << len(quantity)
	sum := make([]int, m)
	for i, ch := range quantity {
		bit := 1 << i
		for j := 0; j < bit; j++ {
			sum[bit|j] = sum[j] + ch
		}
	}
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}
	res := make([]int, 0)
	for _, v := range cnt {
		res = append(res, v)
	}
	n := len(cnt)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m)
		// 初值
		dp[i][0] = true
	}
	i := 0
	for _, c := range res {
		for j, ok := range dp[i] {
			if ok {
				dp[i+1][j] = true
				continue
			}
			// sub 是 j 的子集
			for sub := j; sub > 0; sub = (sub - 1) & j {
				if sum[sub] <= c && dp[i][j^sub] {
					dp[i+1][j] = true
					// break
				}
			}
		}
		i++
	}

	return dp[n][m-1]
}
