package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumGain("cdbcbbaaabab", 4, 5))
}

func maximumGain(s string, x int, y int) int {
	if x < y {
		return maximumGain(reverse(s), y, x)
	}
	ss := []byte(s)
	n := len(ss)
	ans := 0
	i := 0
	for i < n {
		for i < n && ss[i] != 'a' && ss[i] != 'b' {
			i++
		}
		ca, cb := 0, 0
		for i < n && (ss[i] == 'a' || ss[i] == 'b') {
			if ss[i] == 'a' {
				ca++
			} else if ss[i] == 'b' {
				if ca > 0 {
					ans += x
					ca--
				} else {
					cb++
				}
			}
			i++
		}
		ans += min(ca, cb) * y
	}

	return ans
}

func reverse(s string) string {
	ss := []byte(s)
	le, ri := 0, len(s)-1
	for le < ri {
		ss[le], ss[ri] = ss[ri], ss[le]
		le++
		ri--
	}
	return string(ss)
}

// https://leetcode.cn/problems/maximum-score-from-removing-substrings/
