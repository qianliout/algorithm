package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestMerge("abcabc", "abdcaba")) // abdcabcabcaba
}

func largestMerge(word1 string, word2 string) string {
	ans := make([]byte, 0)
	m, n := len(word1), len(word2)
	i, j := 0, 0
	for i < m && j < n {
		// 这样判断是不可以的，因为会出现相同的字符
		// if word1[i] >= word2[j] {
		if word1[i:] >= word2[j:] {
			ans = append(ans, byte(word1[i]))
			i++
		} else {
			ans = append(ans, byte(word2[j]))
			j++
		}
	}

	for k := i; k < m; k++ {
		ans = append(ans, byte(word1[k]))
	}
	for k := j; k < n; k++ {
		ans = append(ans, byte(word2[k]))
	}
	return string(ans)
}
