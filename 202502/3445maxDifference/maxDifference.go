package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(maxDifference("12233", 4))
}

// s 仅由数字 '0' 到 '4' 组成。
// 输入保证至少存在一个子字符串是由一个出现奇数次的字符和一个出现偶数次的字符组成。
func maxDifference(s string, k int) int {
	cnt := make([]int, 5)
	n := len(s)
	for i := 0; i < n; i++ {
		cnt[s[i]-'0']++
		if i < k-1 {
			continue
		}

	}

	// todo(not finnish)
	return 0
}

func maxDifference2(s string) int {
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}
	n := len(s)
	a := 0 // 出现奇数次的最大
	b := n // 出现偶数次的最小
	for i := 0; i < 26; i++ {
		if cnt[i] == 0 {
			continue
		}
		if cnt[i]%2 == 0 {
			b = min(b, cnt[i])
		} else {
			a = max(a, cnt[i])
		}
	}
	return a - b
}

func findTheLongestSubstring(s string) int {
	pos := make([]int, 1<<5+1)
	for i := range pos {
		pos[i] = -1
	}
	status := 0
	ans := 0
	y := []byte("aeiou")
	pos[0] = 0

	for i := range s {
		idx := bytes.LastIndexByte(y, s[i])
		if idx != -1 {
			status = status ^ (1 << idx)
		}

		if pos[status] >= 0 {
			ans = max(ans, i+1-pos[status])
		} else {
			pos[status] = i + 1 // 这里为啥是 i+1呢，因为 pos[0] 表示空集的状态
		}
	}
	return ans
}

// 将 5 个元音字母出现次数的奇偶视为一种状态，一共有 32 种状态，不妨使用一个整数代表状态，第 0 位为 1 表示 a 出现奇数次，
// 第一位为 1 表示 e 出现奇数次……以此类推。仅有状态 0 符合题意。
// 而如果子串 [0，i] 与字串 [0,j] 状态相同，那么字串 [i+1,j] 的状态一定是 0，因此可以记录每个状态第一次出现的位置，此后再出现该状态时相减即可。
// 需要注意状态 0 首次出现的位置应该设定为 -1。
