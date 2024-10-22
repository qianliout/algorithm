package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findReplaceString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	n := len(s)
	ans := make([]string, n)
	replaceStr := make([]string, n)
	replaceL := make([]int, n)
	for i, idx := range indices {
		if strings.HasPrefix(s[idx:], sources[i]) {
			replaceStr[idx] = targets[i]
			replaceL[idx] = len(sources[i])
		}
	}
	for i := 0; i < n; {
		if replaceStr[i] == "" {
			ans = append(ans, string(s[i]))
			i++
		} else {
			ans = append(ans, replaceStr[i])
			i += replaceL[i]
		}
	}
	return strings.Join(ans, "")
}

type Pair struct {
	Idx    int
	Source string
	Target string
}
