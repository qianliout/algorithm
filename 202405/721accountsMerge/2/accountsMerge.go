package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(accountsMerge([][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}))
}

// 会涉及到多次合并，所以这样直接模拟的方式搞不定
func accountsMerge(accounts [][]string) [][]string {

	used := make(map[int]bool)
	email := make(map[string][]int)

	for i, acc := range accounts {
		for _, em := range acc {
			email[em] = append(email[em], i)
		}
	}

	ans := make([][]string, 0)
	for _, v := range email {
		tem := make([]string, 0)
		name := ""
		for _, ch := range v {
			if used[ch] {
				continue
			}
			used[ch] = true
			// 加名字
			if name == "" {
				name = accounts[ch][0] // 防止名字有不一样的
			}
			// 加邮箱
			tem = append(tem, accounts[ch][1:]...)
		}
		tem = dup(tem)
		if len(tem) > 0 && name != "" {
			ans = append(ans, append([]string{name}, tem...))
		}
	}
	return ans
}

func dup(strs []string) []string {
	exit := make(map[string]bool)
	ans := make([]string, 0)
	for _, ch := range strs {
		if !exit[ch] {
			ans = append(ans, ch)
			exit[ch] = true
		}
	}
	sort.Strings(ans)
	return ans
}
