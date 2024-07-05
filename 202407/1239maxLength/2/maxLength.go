package main

import "fmt"

func main() {
	fmt.Println(maxLength([]string{"cha", "r", "act", "ers"}))
}

func maxLength(arr []string) int {
	n := len(arr)
	nums := make([]int, n)
	for i := range arr {
		nums[i] = gen(arr[i])
	}
	ans := 0
	var dfs func(pre int, cnt int)
	visit := make([]bool, n)
	dfs = func(pre int, cnt int) {
		for start := 0; start < n; start++ {
			if visit[start] {
				continue
			}
			visit[start] = true
			if pre&nums[start] != 0 {
				continue
			}
			pre = pre | nums[start]
			ans = max(ans, cnt+len(arr[start]))
			dfs(pre, cnt+len(arr[start]))
		}
	}
	dfs(0, 0)
	return ans
}

func gen(word string) int {
	ans := 0
	for _, ch := range word {
		ans = ans | 1<<(int(ch)-'a')
	}
	return ans
}
