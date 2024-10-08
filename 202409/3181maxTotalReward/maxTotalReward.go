package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxTotalReward([]int{1, 1, 3, 3}))
}

func maxTotalReward2(rewardValues []int) int {
	n := len(rewardValues)
	sort.Ints(rewardValues)
	k := rewardValues[n-1] * 2
	// 定义 f[i][j] 表示能否从前 i 个数中得到总奖励 j。
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, k)
	}
	// 初始值 f[0][0]=true
	f[0][0] = true

	for i := 1; i <= n; i++ {
		v := rewardValues[i-1]
		for j := 0; j < k; j++ {
			f[i][j] = f[i-1][j]
			if j >= v && j-v < v {
				f[i][j] = f[i][j] || f[i-1][j-v]
			}
		}
	}

	// 答案为最大的满足 f[n][j]=true 的 j。
	for j := k - 1; j >= 0; j-- {
		if f[n][j] {
			return j
		}
	}
	return 0
}

// 会超时
func maxTotalReward(rewardValues []int) int {
	n := len(rewardValues)
	sort.Ints(rewardValues)
	m := rewardValues[n-1]
	f := make([]int, m)
	for i := 0; i < n; i++ {
		for j := m - 1; j >= rewardValues[i]; j-- {
			x := j - rewardValues[i]
			if rewardValues[i] <= x {
				x = rewardValues[i] - 1
			}
			f[j] = max(f[j], rewardValues[i]+f[x])
		}
	}
	return f[m-1] + m
}
