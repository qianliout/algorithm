package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(findRightInterval([][]int{{1, 2}}))
	fmt.Println(findRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}}))
}

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	pairs := make([]pair, n)
	for i := range intervals {
		pairs[i] = pair{i, intervals[i][0], intervals[i][1]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].start != pairs[j].start {
			return pairs[i].start < pairs[j].start
		}
		return pairs[i].end < pairs[j].end
	})

	ans := make([]int, n)

	for i, pa := range pairs {
		le, ri := i, n
		for le < ri {
			mid := le + (ri-le)/2
			if mid >= 0 && mid < n && pairs[mid].start >= pa.end {
				ri = mid
			} else {
				le = mid + 1
			}
		}
		if le >= n || le < 0 || pairs[le].start < pa.end {
			ans[pa.idx] = -1
		} else {
			ans[pa.idx] = pairs[le].idx
		}
	}
	return ans
}

type pair struct {
	idx, start, end int
}
