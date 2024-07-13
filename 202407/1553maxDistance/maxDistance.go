package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println(maxDistance([]int{5, 4, 3, 2, 1, 1000000000}, 2))
	fmt.Println(maxDistance([]int{1, 2, 3, 4, 7}, 3))
}

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	if len(position) < m {
		return -1
	}
	mx := slices.Max(position)
	le, ri := 0, mx
	for le < ri {
		mid := le + (ri-le+1)/2
		if le < mx && check(position, mid) >= m {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	return le
}

func check(pos []int, mid int) int {
	n := len(pos)
	start := pos[0]
	cnt := 1
	for i := 1; i < n; i++ {
		if pos[i]-start >= mid {
			cnt++
			start = pos[i]
		}
	}
	return cnt
}
