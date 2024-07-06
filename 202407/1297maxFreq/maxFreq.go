package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxFreq("edbbdaaaacbbdbdacaccecacaeceaaedccdededbdbabebebce", 1, 4, 13))
}

/*
子串中不同字母的数目必须小于等于 maxLetters 。
子串的长度必须大于等于 minSize 且小于等于 maxSize 。
*/
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	wind := make(map[byte]int)
	ans, n := 0, len(s)
	le, ri := 0, 0
	cnt := 0
	m := make(map[string]int)
	for le <= ri && ri < n {
		ch := byte(s[ri])
		if wind[ch] == 0 {
			cnt++
		}
		wind[ch]++
		ri++

		for le <= ri && (cnt > maxLetters || ri-le > minSize) {
			if wind[byte(s[le])] == 1 {
				cnt--
			}
			wind[byte(s[le])]--
			le++
		}

		if cnt <= maxLetters && ri-le == minSize {
			m[s[le:ri]]++
		}
	}
	// 为啥要这样统计，因为子串可以重复，所以 不能用：	strings.Count(s,sub) 这种方式
	for _, v := range m {
		ans = max(ans, v)
	}
	return ans
}
