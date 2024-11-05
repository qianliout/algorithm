package main

import (
	"math"
	"slices"
)

func main() {

}

// 类似合并区间的写法
func numberOfGoodPartitions(nums []int) int {
	m := make(map[int]pair)
	for i, ch := range nums {
		if p, ok := m[ch]; ok {
			p.r = i
			m[ch] = p
		} else {
			m[ch] = pair{l: i, r: i}
		}
	}
	pairs := make([]pair, 0)
	for _, p := range m {
		pairs = append(pairs, p)
	}
	slices.SortFunc(pairs, func(a, b pair) int {
		if a.l != b.l {
			return a.l - b.l
		}
		return a.r - b.r
	})
	ans := 1
	mod := int(math.Pow10(9)) + 7
	mx := pairs[0].r
	for i := 1; i < len(pairs); i++ {
		if pairs[i].l > mx {
			ans = ans * 2 % mod
		}
		mx = max(mx, pairs[i].r)
	}
	return ans
}

type pair struct {
	l, r int
}
