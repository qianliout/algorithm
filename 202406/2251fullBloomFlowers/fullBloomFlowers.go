package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fullBloomFlowers([][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}, []int{2, 3, 7, 11}))
	fmt.Println(fullBloomFlowers([][]int{{19, 37}, {19, 38}, {19, 35}}, []int{6, 7, 21, 1, 13, 37, 5, 37, 46, 43}))
}

// 直接使用数组，会打爆内存，只能使用 hash 表或平衡树
func fullBloomFlowers2(flowers [][]int, people []int) []int {
	mx := 0
	for _, ch := range flowers {
		mx = max(mx, ch[1])
	}
	n := len(people)
	m := mx + 1
	d := make([]int, mx+1)
	for _, ch := range flowers {
		x, y := ch[0], ch[1]+1
		if x >= 0 && x < m {
			d[x]++
		}
		if y >= 0 && y < m {
			d[y]--
		}
	}
	ans := make([]int, n)
	pre := make([]int, m)
	pre[0] = d[0]
	for i := 1; i < m; i++ {
		pre[i] = pre[i-1] + d[i]
	}
	for i, x := range people {
		if x > mx {
			ans[i] = 0
			continue
		}
		ans[i] = pre[x]
	}
	return ans
}

func fullBloomFlowers(flowers [][]int, people []int) []int {
	n := len(flowers)
	start, end := make([]int, n), make([]int, n)
	for i, ch := range flowers {
		start[i] = ch[0]
		end[i] = ch[1]
	}
	sort.Ints(start)
	sort.Ints(end)

	// 二分查找
	ans := make([]int, len(people))
	for i, ch := range people {
		ans[i] = sort.SearchInts(start, ch+1) - sort.SearchInts(end, ch)
	}
	return ans
}
