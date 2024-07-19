package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(minCharacters("acac", "bd"))
}

/*
操作的最终目标是满足下列三个条件 之一 ：
    a 中的 每个字母 在字母表中 严格小于 b 中的 每个字母 。
    b 中的 每个字母 在字母表中 严格小于 a 中的 每个字母 。
    a 和 b 都 由 同一个 字母组成。

*/

func minCharacters(a string, b string) int {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	return min(help1(a, b), help1(b, a), help2(a, b))
}

// a 严格小于 b
// 这里容易出错，严格小于，也可以把 a 改小，也可以把了改大
func help1(a, b string) int {
	cnt1 := make([]int, 26)
	cnt2 := make([]int, 26)
	for _, ch := range a {
		idx := int(ch) - int('a')
		cnt1[idx]++
	}
	for _, ch := range b {
		idx := int(ch) - int('a')
		cnt2[idx]++
	}
	mib := 0
	for i := 0; i < 26; i++ {
		if cnt2[i] > 0 {
			mib = i
			break
		}
	}
	ans := 0
	for i, ch := range cnt1 {
		if i >= mib {
			ans += ch
		}
	}
	return ans
}

// a 和 b 都 由 同一个 字母组成。
func help2(a, b string) int {
	cnt1 := make([]int, 26)
	for _, ch := range a {
		idx := int(ch) - int('a')
		cnt1[idx]++
	}
	for _, ch := range b {
		idx := int(ch) - int('a')
		cnt1[idx]++
	}
	sort.Ints(cnt1)

	return len(a) + len(b) - cnt1[25]
}
