package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestCommonSupersequence("abac", "cba"))
	fmt.Println(shortestCommonSupersequence("bbbaaaba", "bbababbb"))
}

// 这种字符串的拼接很消耗内存，不能通过全部测试用例
func shortestCommonSupersequence1(str1 string, str2 string) string {
	n1, n2 := len(str1), len(str2)
	var dfs func(i, j int) string
	mem := make([][]string, n1)
	for i := range mem {
		mem[i] = make([]string, n2)
	}

	dfs = func(i, j int) string {
		res := ""
		if i < 0 {
			return str2[:j+1]
		}
		if j < 0 {
			return str1[:i+1]
		}
		if mem[i][j] != "" {
			return mem[i][j]
		}
		if str1[i] == str2[j] {
			nex := dfs(i-1, j-1)
			res = nex + string(str1[i])
		} else {
			nex1 := dfs(i-1, j)
			nex2 := dfs(i, j-1)
			if len(nex1) < len(nex2) {
				res = nex1 + string(str1[i])
			} else {
				res = nex2 + string(str2[j])
			}
		}
		mem[i][j] = res
		return res
	}
	a := dfs(n1-1, n2-1)
	return string(a)
}

func shortestCommonSupersequence(str1 string, str2 string) string {
	n1, n2 := len(str1), len(str2)
	// 返回最短路径的长度
	var dfs func(i, j int) int
	mem := make([][]int, n1)
	for i := range mem {
		mem[i] = make([]int, n2)
		for k := range mem[i] {
			mem[i][k] = -1
		}
	}

	dfs = func(i, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		res := n1 + n2
		if str1[i] == str2[j] {
			res = min(res, dfs(i-1, j-1)+1)
		} else {
			nex1 := dfs(i-1, j)
			nex2 := dfs(i, j-1)
			res = min(res, nex1+1, nex2+1)
		}
		mem[i][j] = res
		return res
	}

	var makAns func(i, j int) string

	makAns = func(i, j int) string {
		if i < 0 {
			return str2[:j+1]
		}
		if j < 0 {
			return str1[:i+1]
		}
		if str1[i] == str2[j] {
			return makAns(i-1, j-1) + string(str1[i])
		}
		if dfs(i, j) == dfs(i-1, j)+1 {
			return makAns(i-1, j) + string(str1[i])
		}
		return makAns(i, j-1) + string(str2[j])
	}
	return makAns(n1-1, n2-1)
}
