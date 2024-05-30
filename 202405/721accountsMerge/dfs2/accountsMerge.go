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
	name := make(map[int][]string)  // 名字到 email
	email := make(map[string][]int) // email 到名字,这里一定要用名字的下标，因为相同的名字也有可能属于不同的组
	for i, ac := range accounts {
		name[i] = append(name[i], ac[1:]...)
		for j := 1; j < len(ac); j++ {
			email[ac[j]] = append(email[ac[j]], i)
		}
	}
	ans := make([][]string, 0)
	emailVisit, nameVisit := make(map[string]bool), make(map[int]bool)
	for i, ac := range accounts {
		em := dfs(name, email, emailVisit, nameVisit, i)
		if len(em) > 0 {
			ans = append(ans, append([]string{ac[0]}, dup(em)...))
		}
	}

	return ans
}

func dfs(name map[int][]string, email map[string][]int, emailVisit map[string]bool, nameVisit map[int]bool, startName int) []string {
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
