package main

import "fmt"

func main() {
	// 测试用例
	nums1 := []int{1, 2, 3, 4}
	quantity1 := []int{2}
	fmt.Printf("测试1: nums=%v, quantity=%v, 结果=%v\n", nums1, quantity1, canDistribute(nums1, quantity1))

	nums2 := []int{1, 2, 3, 3}
	quantity2 := []int{2}
	fmt.Printf("测试2: nums=%v, quantity=%v, 结果=%v\n", nums2, quantity2, canDistribute(nums2, quantity2))

	nums3 := []int{1, 1, 2, 2}
	quantity3 := []int{2, 2}
	fmt.Printf("测试3: nums=%v, quantity=%v, 结果=%v\n", nums3, quantity3, canDistribute(nums3, quantity3))
}

/*
核心思路：状态压缩DP + 子集枚举

问题分析：
1. 每个顾客需要相同的数字，且数量固定
2. 需要判断是否能满足所有顾客的需求
3. 关键约束：quantity数组长度最多10，nums中最多50个不同值

解题技巧：
1. 状态压缩：用二进制表示顾客集合，1表示已满足，0表示未满足
2. 预处理：计算所有顾客子集的总需求量
3. DP状态：dp[i][mask] 表示前i种数字能否满足mask表示的顾客集合
4. 子集枚举：对于每种数字，枚举它能满足的顾客子集
*/
func canDistribute(nums []int, quantity []int) bool {
	// 第一步：预处理 - 计算所有顾客子集的总需求量
	// m = 2^len(quantity)，表示所有可能的顾客组合
	m := 1 << len(quantity)
	sum := make([]int, m) // sum[mask] = mask对应顾客集合的总需求量

	// 使用动态规划思想计算所有子集的和
	// 对于每个顾客i，将其加入到所有不包含i的子集中
	for i, ch := range quantity {
		bit := 1 << i // 第i个顾客对应的位
		// 遍历所有不包含第i个顾客的子集
		for j := 0; j < bit; j++ {
			// bit|j 表示在子集j的基础上加入第i个顾客
			sum[bit|j] = sum[j] + ch
		}
	}

	// 第二步：统计每种数字的出现次数
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}

	// 将计数转换为数组，方便后续处理
	res := make([]int, 0)
	for _, v := range cnt {
		res = append(res, v)
	}

	// 第三步：动态规划
	n := len(cnt) // 不同数字的种类数
	// dp[i][mask] 表示前i种数字能否满足mask表示的顾客集合
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m)
		// 初始状态：不满足任何顾客（mask=0）总是可能的
		dp[i][0] = true
	}

	// 第四步：状态转移
	i := 0
	for _, c := range res { // c是当前数字的出现次数
		for j, ok := range dp[i] { // j是顾客集合的状态
			if ok { // 如果前i种数字能满足状态j
				// 不使用当前数字，直接继承状态
				dp[i+1][j] = true
				continue
			}

			// 关键：子集枚举技巧
			// 枚举j的所有子集sub，尝试用当前数字满足sub中的顾客
			for sub := j; sub > 0; sub = (sub - 1) & j {
				// 检查两个条件：
				// 1. sum[sub] <= c：当前数字的数量足够满足sub中的顾客
				// 2. dp[i][j^sub]：前i种数字能满足j-sub的顾客
				if sum[sub] <= c && dp[i][j^sub] {
					dp[i+1][j] = true
					break // 找到一种方案即可
				}
			}
		}
		i++
	}

	// 第五步：返回结果
	// dp[n][m-1] 表示所有数字能否满足所有顾客（m-1的二进制全为1）
	return dp[n][m-1]
}

/*
给你一个长度为 n 的整数数组 nums ，这个数组中至多有 50 个不同的值。同时你有 m 个顾客的订单 quantity ，其中，
整数 quantity[i] 是第 i 位顾客订单的数目。请你判断是否能将 nums 中的整数分配给这些顾客，且满足：
第 i 位顾客 恰好 有 quantity[i] 个整数。
第 i 位顾客拿到的整数都是 相同的 。
每位顾客都满足上述两个要求。
如果你可以分配 nums 中的整数满足上面的要求，那么请返回 true ，否则返回 false 。
*/
