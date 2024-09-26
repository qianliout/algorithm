package main

import (
	"fmt"
)

func main() {
	// fmt.Println(countKConstraintSubstrings("0001111", 2, [][]int{{0, 6}}))
	fmt.Println(countKConstraintSubstrings("010101", 1, [][]int{{0, 5}, {1, 4}, {2, 3}})) // [15,9,3]
}

func countKConstraintSubstrings2(s string, k int, queries [][]int) []int64 {
	n := len(s)
	cnt := make([]int, 2)

	left := make([]int, n)
	pre := make([]int, n+1)
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[(s[l])&1]--
			l++
		}
		left[i] = l
		pre[i+1] = pre[i] + (i - l + 1)
	}

	check := func(j, l int) bool {
		return left[j] >= l
	}

	ans := make([]int64, len(queries))

	for i, ch := range queries {
		le, ri := ch[0], ch[1]
		if left[ri] <= le {
			ans[i] = int64((ri - le + 2) * (ri - le + 1) / 2)
		} else {
			// 二分找第一个满足 left[j]≥le 的下标 j,相当于找左端点
			j, r0 := 0, n
			for j < r0 {
				mid := (j + r0) >> 1
				if mid < n && check(mid, le) {
					r0 = mid
				} else {
					j = mid + 1
				}
			}

			ans[i] = int64((j-l+1)*(j-l)/2) + int64(pre[ri+1]-pre[j])
		}
	}
	return ans
}

func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	cnt := make([]int, 2)

	left := make([]int, n)
	pre := make([]int, n+1)
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[(s[l])&1]--
			l++
		}
		left[i] = l
		pre[i+1] = pre[i] + (i - l + 1)
	}

	ans := make([]int64, len(queries))

	// ans1 := make([]int64, len(queries))
	// for i, q := range queries {
	// 	l, r := q[0], q[1]
	// 	j := l + sort.SearchInts(left[l:r+1], l)
	// 	ans1[i] = int64(pre[r+1] - pre[j] + (j-l+1)*(j-l)/2)
	// }

	for i, ch := range queries {
		le, ri := ch[0], ch[1]
		if left[ri] <= le {
			ans[i] = int64((ri - le + 2) * (ri - le + 1) / 2)
		} else {
			j := bitSearch(left, le)
			ans[i] = int64((j-le+1)*(j-le)/2) + int64(pre[ri+1]-pre[j])
		}
	}
	return ans
}

func bitSearch(left []int, le int) int {
	n := len(left)
	l, r := 0, n
	for l < r {
		// 左端点
		mid := l + (r-l)/2
		if mid < n && left[mid] >= le {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
