package main

import (
	"fmt"
)

func main() {
	// fmt.Println(countCompleteSubstrings("igigee", 2))
	fmt.Println(countCompleteSubstrings("aaabbbccc", 3))
}

func countCompleteSubstrings(word string, k int) int {
	words := split(word)
	ans := 0
	for _, ch := range words {
		ans += count(ch, k)
	}
	return ans
}

func split(word string) []string {
	ans := make([]string, 0)
	i, n := 0, len(word)
	for i < n {
		start := i
		i++
		for i < n && abs(int(word[i])-int(word[i-1])) <= 2 {
			i++
		}
		ans = append(ans, word[start:i])
	}

	return ans
}

func count(word string, k int) int {
	n, ans := len(word), 0
	for j := 1; j <= 26; j++ {
		size := j * k
		if size > n {
			break
		}

		cnt := make(map[int]int)
		for i := 0; i < size; i++ {
			c := int(word[i] - 'a')
			cnt[c]++
		}
		ans += check(cnt, k)

		le, ri := 0, size
		for le < ri && ri < n {
			cnt[int(word[ri]-'a')]++
			cnt[int(word[le]-'a')]--
			ans += check(cnt, k)
			ri++
			le++
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func check(cnt map[int]int, k int) int {
	if len(cnt) == 0 {
		return 0
	}
	for _, v := range cnt {
		if v == 0 {
			continue
		}
		if v != k {
			return 0
		}
	}
	return 1
}
