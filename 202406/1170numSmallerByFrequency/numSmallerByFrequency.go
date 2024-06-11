package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
	fmt.Println(numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
}

func numSmallerByFrequency(queries []string, words []string) []int {
	qn := len(queries)
	wn := len(words)
	q := make([]int, qn)
	for i := range queries {
		q[i] = f(queries[i])
	}
	w := make([]int, len(words))
	for i := range w {
		w[i] = f(words[i])
	}
	sort.Ints(w)
	ans := make([]int, qn)
	for i, x := range q {
		// 需统计 words 中满足 f(queries[i]) < f(W) 的 词的数目
		// 这里使用 w[k]>=x+1是一个技巧，使用 w[k]>=x 会出错
		// j := sort.Search(wn, func(k int) bool { return w[k] >= x+1 })
		// 这样写也是可以的
		j := sort.Search(wn, func(k int) bool { return w[k] > x })
		ans[i] = wn - j
		// ans[i] = wn - sort.SearchInts(w, x+1)
	}
	return ans
}

func f(s string) int {
	var aa byte
	cnt := make(map[byte]int)
	for _, c := range s {
		cnt[byte(c)]++
		if aa == byte(0) || byte(c) < aa {
			aa = byte(c)
		}
	}
	return cnt[aa]
}
