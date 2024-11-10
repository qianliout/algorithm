package main

import (
	"fmt"
)

func main() {
	fmt.Println(groupAnagrams([]string{"ape", "and", "cat"}))
}

func groupAnagrams(strs []string) [][]string {
	cnt := make(map[ana][]string)
	for _, s := range strs {
		a := gen(s)
		cnt[a] = append(cnt[a], s)
	}
	ans := make([][]string, 0)
	for _, v := range cnt {
		if len(v) > 0 {
			ans = append(ans, v)
		}
	}
	return ans
}

type ana [26]int

func gen(s string) ana {
	var a ana
	for _, v := range s {
		b := int(v - 'a')
		a[b]++
	}
	return ana(a)
}
