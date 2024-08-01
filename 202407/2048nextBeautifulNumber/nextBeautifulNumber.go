package main

import (
	"math"
)

func main() {

}

func nextBeautifulNumber(n int) int {
	le := n + 1
	// 不能二分，没有两段性
	for le < math.MaxInt64/100 {
		if check(le) {
			return le
		}
		le++
	}
	return le
}

func check(n int) bool {
	cnt := make(map[int]int)
	for n > 0 {
		cnt[n%10]++
		n /= 10
	}
	for k, v := range cnt {
		if k != v {
			return false
		}
	}
	return true
}
