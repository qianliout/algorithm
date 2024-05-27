package main

import (
	"fmt"
)

func main() {
	fmt.Println(expressiveWords("heeellooo", []string{"hello", "hi", "helo"}))
}

func expressiveWords(s string, words []string) int {
	ans := make([]string, 0)
	for i := range words {
		if check(s, words[i]) {
			ans = append(ans, words[i])
		}
	}
	return len(ans)
}

func check(s, word string) bool {
	ai, bi := 0, 0

	for ai < len(s) && bi < len(word) {
		start := bi
		ac, bc := 0, 0
		for ai < len(s) && s[ai] == word[start] {
			ac++
			ai++
		}
		for bi < len(word) && word[bi] == word[start] {
			bc++
			bi++
		}
		/*
			【条件1】由于“e”和“o”在s中连续出现的次数是3（即：满足>=3），所以要求e和o在word中连续出现的次数不能超过3次
			【条件2】由于“h”和“l”在s中连续出现的次数不满足>=3，所以要求h和l在word中连续出现的次数要与s中出现的次数相同
		*/
		if (ac != bc && ac < 3) || (ac < bc && ac >= 3) {
			return false
		}

		// if !((ac != bc && ac >= 3 && bc < 3) || (ac < 3 && ac == bc)) {
		// 	return false
		// }
	}
	return ai == len(s) && bi == len(word)
}
