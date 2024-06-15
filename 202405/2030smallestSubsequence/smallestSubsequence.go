package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(smallestSubsequence("leetcode", 4, 'e', 2))
	// fmt.Println(smallestSubsequence("leet", 3, 'e', 1))
	fmt.Println(smallestSubsequence("aaabbbcccddd", 3, 'b', 2)) // abb
	// fmt.Println(smallestSubsequence("adffhjfmmmmorsfff", 6, 'f', 5))
	// fmt.Println(smallestSubsequence("hjjhhhmhhwhz", 6, 'h', 5))
}

func smallestSubsequence(s string, k int, letter byte, repetition int) string {
	ans := make([]byte, 0)
	unread := strings.Count(s, string(letter)) // 未遍历到的 letter 个数
	inQ := 0                                   // 在单调队列中的 letter 个数
	stark := make([]byte, 0)

	n := len(s)

	ri := 0
	for len(ans) < k {
		// 且 字典序最小 的子序列
		for ri < n && len(stark) > 0 && stark[len(stark)-1] > byte(s[ri]) {
			if stark[len(stark)-1] == letter {
				if inQ+unread <= repetition {
					break
				}
				inQ--
			}
			stark = stark[:len(stark)-1]
		}
		if ri < n {
			if s[ri] == letter {
				inQ++
				unread--
			}
			stark = append(stark, s[ri])
		}

		if ri >= n-k {
			if stark[0] == letter {
				inQ--
				repetition--
				ans = append(ans, stark[0])
			} else if len(ans)+repetition < k { // 还有足够空间可以放 letter
				ans = append(ans, stark[0])
			}
			stark = stark[1:]
		}
		ri++
	}

	return string(ans)
}
