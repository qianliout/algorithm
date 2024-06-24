package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumTastiness([]int{13, 5, 1, 8, 21, 2}, 3))
	fmt.Println(maximumTastiness([]int{1, 3, 1}, 2))
}

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	n := len(price)
	mi, mx := 0, price[n-1]
	le, ri := mi, mx
	for le < ri {
		mid := le + (ri-le+1)/2
		if mid >= mi && mid < mx && check(price, mid) >= k {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	if check(price, le) >= k {
		return le
	}
	return 0
}

func check(price []int, mx int) int {
	res := 1
	start := price[0]
	for i := 1; i < len(price); i++ {
		if price[i]-start >= mx {
			start = price[i]
			res++
		}
	}
	return res
}
