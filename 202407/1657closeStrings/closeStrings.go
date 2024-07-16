package main

import (
	"fmt"
)

func main() {
	// fmt.Println(closeStrings("cabbba", "abbccc"))
	fmt.Println(closeStrings("aaabbbbccddeeeeefffff", "aaaaabbcccdddeeeeffff"))
}

// 如果 s 和 t 的字符一样，并且字符出现次数的集合是相同的（比如上面这个例子都是集合 {1,2,3}），那么可以结合操作 1 和操作 2，把 s 变成 t
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	cnt1 := make(map[int]int)
	cnt2 := make(map[int]int)
	n := len(word1)
	for i := 0; i < n; i++ {
		cnt1[int(word1[i])]++
		cnt2[int(word2[i])]++
	}
	if !same1(cnt1, cnt2) {
		return false
	}
	cnt3 := make(map[int]int)
	cnt4 := make(map[int]int)
	for _, v := range cnt1 {
		cnt3[v]++
	}
	for _, v := range cnt2 {
		cnt4[v]++
	}

	return same2(cnt3, cnt4)
}

func same1(cnt1, cnt2 map[int]int) bool {
	if len(cnt1) != len(cnt2) {
		return false
	}
	for k := range cnt1 {
		if cnt2[k] <= 0 {
			return false
		}
	}
	return true
}

func same2(cnt1, cnt2 map[int]int) bool {
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
