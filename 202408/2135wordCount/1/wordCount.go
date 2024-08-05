package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordCount([]string{"ant", "act", "tack"}, []string{"tack", "act", "acti"}))
	fmt.Println(wordCount([]string{"g", "vf", "ylpuk", "nyf", "gdj", "j", "fyqzg", "sizec"}, []string{"r", "am", "jg", "umhjo", "fov", "lujy", "b", "uz", "y"}))
}

// 在 startWords 或 targetWords 的任一字符串中，每个字母至多出现一次
func wordCount(startWords []string, targetWords []string) int {
	target := make([]pair, 0)
	for _, ch := range targetWords {
		cnt1 := make(map[byte]int)
		for _, by := range ch {
			cnt1[byte(by)]++
		}
		target = append(target, cnt1)
	}
	ans := 0
	find := make([]bool, len(target))
	for _, st := range startWords {
		ret := cal(st)
		for _, ch1 := range ret {

			for i, ch2 := range target {
				if find[i] {
					continue
				}
				if equal(ch1, ch2) {
					ans++
					find[i] = true
				}
			}
		}
	}

	return ans
}

func cal(s string) []pair {
	// 增加一个
	cnt1 := make(map[byte]int)
	for i := range s {
		cnt1[s[i]]++
	}
	ans := make([]pair, 0)
	for i := 'a'; i <= 'z'; i++ {
		no := copyPair(cnt1)
		no[byte(i)]++
		ans = append(ans, no)
	}
	return ans
}

type pair map[byte]int

func copyPair(cnt1 pair) pair {
	ans := make(map[byte]int)
	for k, v := range cnt1 {
		ans[k] = v
	}
	return ans
}

func equal(cnt1, cnt2 pair) bool {
	if len(cnt1) != len(cnt2) {
		return false
	}
	for k, v := range cnt1 {
		if cnt2[k] != v {
			return false
		}
	}
	return true
}
