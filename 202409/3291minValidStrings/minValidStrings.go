package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(minValidStrings([]string{"abc", "aaaaa", "bcdef"}, "aabcdabc"))
}

func minValidStrings(words []string, target string) int {
	ans := dfs(words, target)
	if ans >= math.MaxInt/10 {
		return -1
	}
	return ans
}

// 这样写是不可以的，因为匹配时可以是任意顺序匹配 words
func dfs(words []string, target string) int {
	inf := math.MaxInt / 10
	if target == "" {
		return 0
	}
	if len(words) == 0 {
		return inf
	}
	first := words[0]
	pre := findPrefix(first, target)
	ans1, ans2 := inf, inf
	if pre != "" {
		ans1 = dfs(words[1:], strings.TrimPrefix(target, pre)) + 1
	}
	ans2 = dfs(words[1:], target)
	return min(ans1, ans2)
}

func findPrefix(a, b string) string {
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return a[:i]
}
