package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(shortestSuperstring([]string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"}))
	fmt.Println(shortestSuperstring([]string{"alex", "loves", "leetcode"}))
}

func shortestSuperstring(words []string) string {
	n := len(words)
	pairs := make([]pair, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			pairs = append(pairs, gen(words[i], words[j]))
		}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].c > pairs[j].c })
	ans := make([]byte, 0)
	visit := make(map[string]bool)
	for _, pa := range pairs {
		if visit[pa.a] {
			continue
		}
		if visit[pa.b] {
			continue
		}
		ans = append(ans, []byte(pa.a)...)
		ans = append(ans, []byte(pa.b[pa.c:])...)
		visit[pa.a] = true
		visit[pa.b] = true
	}
	return string(ans)
}

type pair struct {
	a, b string
	c    int
}

func gen(a, b string) pair {
	x, y := len(a)-1, 0
	ans := 0
	for x >= 0 && y < len(b) {
		if a[x] == b[y] {
			ans++
			x--
			y++
		} else {
			break
		}
	}
	return pair{
		a: a,
		b: b,
		c: ans,
	}
}
