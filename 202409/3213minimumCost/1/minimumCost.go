package main

import (
	"index/suffixarray"
	"math"
)

func main() {

}

// minimumCost 计算达到目标字符串target的最小成本。
// 参数:
//
//	target - 目标字符串
//	words - 一组单词
//	costs - 对应单词的成本数组
//
// 返回值:
//
//	最小成本，如果无法达到目标字符串则返回-1
func minimumCost(target string, words []string, costs []int) int {
	// minCost用于存储每个单词的最小成本
	minCost := map[string]uint16{}
	// 遍历words数组，初始化minCost映射
	for i, w := range words {
		c := uint16(costs[i])
		// 如果单词w不在minCost中，或者已有成本更高，则更新成本为当前成本c
		if minCost[w] == 0 {
			minCost[w] = c
		} else {
			minCost[w] = min(minCost[w], c)
		}
	}

	n := len(target)
	// from数组用于存储每个位置可以由哪些单词到达，以及到达的成本
	from := make([][]pair, n+1)
	// 构建目标字符串的后缀数组，用于高效查找单词在目标字符串中的位置
	sa := suffixarray.New([]byte(target))
	// 遍历minCost中的每个单词，查找它们在目标字符串中的位置
	for w, c := range minCost {
		// 对于每个单词w，在目标字符串中查找所有出现位置
		for _, l := range sa.Lookup([]byte(w), -1) {
			r := l + len(w)
			// 将单词到达位置和成本添加到from数组中
			from[r] = append(from[r], pair{uint16(l), c})
		}
	}

	// f数组用于动态规划存储到达每个位置的最小成本
	f := make([]int, n+1)
	// 初始化起始位置成本为无穷大，表示不可达
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		// 遍历所有可以到达当前位置的单词和成本
		for _, p := range from[i] {
			// 更新到达当前位置的最小成本
			f[i] = min(f[i], f[p.l]+int(p.cost))
		}
	}
	// 如果最终位置的最小成本仍然是无穷大，则表示目标字符串不可达，返回-1
	if f[n] == math.MaxInt/2 {
		return -1
	}
	// 返回到达最终位置的最小成本
	return f[n]
}

type pair struct {
	l    uint16
	cost uint16
}
