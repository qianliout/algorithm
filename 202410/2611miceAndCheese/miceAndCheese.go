package main

import (
	"sort"
)

func main() {

}

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	// 先假设把所有的奶酪都给2吃，那么总得分为all
	// 现在假设把第i个奶酪给1吃，那么总得分为all+(reward1[i]-reward2[])
	// 由此可以知道，要使用得分最大，就需要选(reward1[i]-reward2[i]) 的前 k 个

	n := len(reward1)
	p := make([]pair, n)
	all := 0
	for i := range reward2 {
		p[i] = pair{idx: i, sub: reward1[i] - reward2[i]}
		all += reward2[i]
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].sub >= p[j].sub
	})
	for i := 0; i < k; i++ {
		all += p[i].sub
	}
	return all
}

type pair struct {
	idx int
	sub int
}
