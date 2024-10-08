package main

import "sort"

func main() {

}

func maxTotalReward2(rewardValues []int) int {
	n := len(rewardValues)
	sort.Ints(rewardValues)
	k := rewardValues[n-1]
	// 定义 f[i][j] 表示能否从前 i 个数中得到总奖励 j。
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, k+1)
	}

	for i := 0; i <= k; i++ {
		f[0][i] = min(rewardValues[0], i)
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			f[i][j] = f[i-1][j]
			if j >= rewardValues[i] {
				f[i][j] = max(f[i][j], f[i-1][j-rewardValues[i]]+rewardValues[i])
			}
		}
	}

	return f[n-1][k]
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
