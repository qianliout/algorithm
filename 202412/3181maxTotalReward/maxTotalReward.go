package main

import (
	"fmt"
	"math/big"
	"slices"
	"sort"
)

func main() {
	fmt.Println(maxTotalReward([]int{1, 1, 3, 3}))
}

func maxTotalReward1(rewardValues []int) int {
	n := len(rewardValues)
	sort.Ints(rewardValues)
	k := rewardValues[n-1] * 2
	// 定义 f[i][j] 表示能否从前 i 个数中得到总奖励 j。
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, k)
	}
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
	for j := k - 1; j >= 0; j-- {
		if f[n][j] {
			return j
		}
	}
	return 0
}
func maxTotalReward2(rewardValues []int) int {
	n := len(rewardValues)
	sort.Ints(rewardValues)
	var dfs func(i, s int) int
	mem := make([]map[int]int, n+1)
	for i := range mem {
		mem[i] = make(map[int]int)
	}
	dfs = func(i, pre int) int {

		if i >= n {
			return pre
		}
		if i > 0 || rewardValues[i-1] == rewardValues[i] {
			return pre
		}

		if rewardValues[i] <= pre {
			return pre
		}

		if v, ok := mem[i][pre]; ok {
			return v
		}

		ans := dfs(i+1, pre)
		if rewardValues[i] > pre {
			a := dfs(i+1, rewardValues[i]+pre)
			ans = max(a, ans)
		}
		mem[i][pre] = ans
		return ans
	}
	ans := dfs(0, 0)
	return ans
}

// bitset 优化
// maxTotalReward 计算最大的总奖励值。
// 参数 rewardValues 是一个整数切片，表示各个奖励值。
// 返回值是所有奖励值的最大总和。
func maxTotalReward(rewardValues []int) int {
	// 对奖励值进行排序，以便后续处理。
	slices.Sort(rewardValues)
	// 去除重复的奖励值，因为重复的奖励值不会影响最大总奖励值。
	rewardValues = slices.Compact(rewardValues) // 去重

	// 初始化大整数变量one为1，用于位运算。
	one := big.NewInt(1)
	// 初始化大整数变量f为1，用于记录总奖励值的位表示。
	f := big.NewInt(1)
	// 初始化大整数变量p用于辅助位运算。
	p := new(big.Int)

	// 遍历排序后的奖励值。
	for _, v := range rewardValues {
		// 计算当前奖励值的掩码，用于清除位。
		mask := p.Sub(p.Lsh(one, uint(v)), one)
		// 更新总奖励值的位表示，考虑当前奖励值。
		f.Or(f, p.Lsh(p.And(f, mask), uint(v)))
	}

	// 返回最大总奖励值，减1是因为计算的是位长度。
	return f.BitLen() - 1
}
