package main

import (
	"math/bits"
)

func main() {

}

func maximumRequests(n int, requests [][]int) int {
	m := len(requests)
	ans := 0
	for i := 0; i < 1<<m; i++ {
		cnt := bits.OnesCount(uint(i))
		if cnt < ans {
			continue
		}
		if check(requests, i) {
			ans = max(ans, bits.OnesCount(uint(i)))
		}
	}
	return ans
}

func check(requests [][]int, j int) bool {
	cnt := make([]int, 20) // 题目中规定的数据大小
	sum := 0
	for i := 0; i < len(requests); i++ {
		f, t := requests[i][0], requests[i][1]
		if j>>i&1 == 1 {
			cnt[f]++
			if cnt[f] == 1 {
				sum++
			}
			cnt[t]--
			if cnt[t] == 0 {
				sum--
			}
		}
	}
	return sum == 0
}
