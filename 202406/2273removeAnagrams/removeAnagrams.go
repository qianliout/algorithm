package main

import (
	"sort"
)

func main() {

}

func removeAnagrams(words []string) []string {
	n := len(words)
	ans := make([]string, 0)
	i := 0
	for i < n {
		ans = append(ans, words[i])
		j := i + 1

		for j < n && check(words[i], words[j]) {
			j++
		}
		i = j
	}
	return ans
}

func check(pre, cur string) bool {
	ss := []byte(pre)
	sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })

	ss2 := []byte(cur)
	sort.Slice(ss2, func(i, j int) bool { return ss2[i] < ss2[j] })
	return string(ss) == string(ss2)
}
