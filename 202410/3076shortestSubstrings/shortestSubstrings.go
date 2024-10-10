package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(shortestSubstrings([]string{"gfnt", "xn", "mdz", "yfmr", "fi", "wwncn", "hkdy"}))
}

// 数据量小，直接暴力

func shortestSubstrings(arr []string) []string {
	res := make([]string, len(arr))
	for i, ch := range arr {
		ans := ""
		n := len(ch)
		for j := 0; j < n; j++ {
			for k := j + 1; k <= n; k++ {
				sub := ch[j:k]
				if help(sub, i, arr) && (ans == "" || (len(ans) == len(sub) && ans > sub) || (len(sub) < len(ans))) {
					ans = sub
				}
			}
		}
		res[i] = ans
	}

	return res
}

func help(str string, idx int, arr []string) bool {
	for i, ch := range arr {
		if i == idx {
			continue
		}
		if strings.Contains(ch, str) {
			return false
		}
	}
	return true
}
