package main

import (
	"fmt"
	"math/bits"
	"slices"
)

func main() {
	// fmt.Println(findKthSmallest([]int{3, 6, 9}, 3))
	fmt.Println(findKthSmallest([]int{5, 2}, 7))
}

func findKthSmallest(coins []int, k int) int64 {
	var check func(mid int) bool
	n := len(coins)
	check = func(mid int) bool {
		cnt := 0
		for i := 1; i < 1<<n; i++ {
			lcmRes := 1
			for j := 0; j < n; j++ {
				if (i>>j)&1 != 0 {
					lcmRes = lcm(lcmRes, coins[j])
				}
			}
			c := mid / lcmRes
			if bits.OnesCount(uint(i))%2 == 0 {
				c = -c
			}
			cnt += c
		}

		return cnt >= k
	}

	mi := slices.Min(coins)
	mx := mi*k + 1
	le, ri := mi, mx
	for le < ri {
		mid := le + (ri-le)/2
		if le >= mi && le < mx && check(mid) {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return int64(le)
}

func lcm(a, b int) int {
	return a * b / gcb(a, b)
}

func gcb(a, b int) int {
	if b == 0 {
		return a
	}
	return gcb(b, a%b)
}
