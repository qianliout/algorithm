package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkStrings("abcdba", "cabdab"))
}

func checkStrings(s1 string, s2 string) bool {
	if len(s2) != len(s1) {
		return false
	}

	cnt1 := make([]int, 26)
	cnt2 := make([]int, 26)
	n := len(s1)

	for i := 0; i < n; i = i + 2 {
		idx1 := int(s1[i]) - int('a')
		idx2 := int(s2[i]) - int('a')
		cnt1[idx1]++
		cnt2[idx2]++
	}
	if !same(cnt1, cnt2) {
		return false
	}

	cnt1 = make([]int, 26)
	cnt2 = make([]int, 26)

	for i := 1; i < n; i = i + 2 {
		idx1 := int(s1[i]) - int('a')
		idx2 := int(s2[i]) - int('a')
		cnt1[idx1]++
		cnt2[idx2]++
	}
	if !same(cnt1, cnt2) {
		return false
	}
	return true
}

func same(cnt1, cnt2 []int) bool {
	for i := 0; i < len(cnt2); i++ {
		if cnt1[i] != cnt2[i] {
			return false
		}
	}
	return true
}
