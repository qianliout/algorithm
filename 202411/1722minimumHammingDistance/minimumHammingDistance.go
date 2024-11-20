package main

import (
	. "outback/algorithm/common/unionfind"
)

func main() {

}

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	uf := NewRankUnionFind(n + 1)
	for _, ch := range allowedSwaps {
		x, y := ch[0], ch[1]
		uf.Union(x, y)
	}
	mp := make(map[int][]int)
	for i := range source {
		root := uf.Find(i)
		mp[root] = append(mp[root], i)
	}
	ans := 0

	for _, lst := range mp {
		cnt1 := make(map[int]int)
		cnt2 := make(map[int]int)
		for _, i := range lst {
			cnt1[source[i]]++
			cnt2[target[i]]++
		}
		for k, v := range cnt2 {
			// 还是没有能理解
			// 按道理cnt2[k]-v 的也是不同的，也要统计啊，但是为啥没有统计呢
			// 这是因为，如果也统计的话，就会重复
			ans += max(0, v-cnt1[k])
		}
	}
	return ans
}
