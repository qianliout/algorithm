package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(accountsMerge([][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}))
}

// 会涉及到多次合并，所以这样直接模拟的方式搞不定
// 这种方式也是不得行，因为相同名字，可能会是不同的账号
func accountsMerge(accounts [][]string) [][]string {
	name := make(map[string][]string)
	email := make(map[string][]string)
	for _, ac := range accounts {
		name[ac[0]] = append(name[ac[0]], ac[1:]...)
		for _, em := range ac[1:] {
			email[em] = append(email[em], ac[0])
		}
	}
	ans := make([][]string, 0)
	emailVisit, nameVisit := make(map[string]bool), make(map[string]bool)
	for na := range name {
		em := dfs(name, email, emailVisit, nameVisit, na)
		if len(em) > 0 {
			ans = append(ans, append([]string{na}, dup(em)...))
		}
	}
	return ans
}

func dfs(name, email map[string][]string, emailVisit, nameVisit map[string]bool, startName string) []string {
	ans1 := make([]string, 0)
	if nameVisit[startName] {
		return ans1
	}
	nameVisit[startName] = true
	for _, em := range name[startName] {
		if emailVisit[em] {
			continue
		}
		emailVisit[em] = true
		ans1 = append(ans1, em)
	}
	for _, em := range ans1 {
		namess := email[em]
		for _, na := range namess {
			ans1 = append(ans1, dfs(name, email, emailVisit, nameVisit, na)...)
		}
	}
	return ans1
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
