package main

import (
	"sort"
)

func main() {

}

func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	ans, n := 0, len(beans)
	sum := 0
	for i := 0; i < n; i++ {
		sum += beans[i]
	}
	for i := 0; i < n; i++ {
		ans = max(ans, beans[i]*(n-i))
	}
	return int64(sum) - int64(ans)
}
