package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxFreq("aababcaab", 2, 3, 4))
}

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	le, ri, n := 0, 0, len(s)
	win := make(map[byte]int)
	cnt := make(map[string]int)
	aa := 0
	for le <= ri && ri < n {
		if win[s[ri]] == 0 {
			aa++
		}
		win[s[ri]]++
		ri++

		// 一定得是先缩小窗口，再更新答案，是为啥呢
		if ri-le >= minSize && ri-le <= maxSize && aa <= maxLetters {
			cnt[s[le:ri]]++
		}

		// 不再满足条件时才缩小窗口
		// 首先我们需要明确一个结论：如果长度长的串重复，那么他的子串重复次数一定大于等于他。如abcabc，
		// 其中abc出现了两次，他的子串a，ab，bc同样出现两次，也有可能更多。所以题目给我们的maxSize是干扰项，
		// 我们只需要找到大于minSize的符合要求的字符串就好了。
		for le <= ri && (aa > maxLetters || ri-le > minSize) {
			if win[s[le]] == 1 {
				aa--
			}
			win[s[le]]--
			le++
		}
	}
	ans := 0
	for _, v := range cnt {
		ans = max(ans, v)
	}
	return ans
}

// 这样写为啥不对,没有想明白
func maxFreq2(s string, maxLetters int, minSize int, maxSize int) int {
	le, ri, n := 0, 0, len(s)
	win := make(map[byte]int)
	cnt := make(map[string]int)
	aa := 0
	for le <= ri && ri < n {

		win[s[ri]]++
		if win[s[ri]] == 1 {
			aa++
		}
		ri++
		// 问题在于：
		// 这个逻辑会导致窗口在满足条件时立即缩小，而不是等到窗口不再满足条件时才缩小。
		// 这样会导致很多符合条件的子字符串没有被正确记录，因为窗口在记录一次后就被缩小了。
		for ri-le >= minSize && ri-le <= maxSize && aa <= maxLetters {
			cnt[s[le:ri]]++
			win[s[le]]--
			if win[s[le]] == 0 {
				aa--
			}
			le++
		}
	}
	ans := 0
	for _, v := range cnt {
		ans = max(ans, v)
	}
	return ans
}
