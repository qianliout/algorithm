package main

import "fmt"

func main() {
	fmt.Println(maxLength([]string{"cha", "r", "act", "ers"}))
}

// 这样写为啥是错的，有问题，不知道是啥原因
func maxLength(arr []string) int {
	nums := make([]int, 0)
	str := make([]string, 0)
	for i := range arr {
		a, b := gen(arr[i])
		if b {
			nums = append(nums, a)
			str = append(str, arr[i])
		}
	}
	arr = str
	n := len(arr)
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
			ans = max(ans, cnt+len(arr[start]))
			dfs(pre|nums[start], cnt+len(arr[start]))
			visit[start] = false
		}
	}
	for i := 0; i < n; i++ {
		dfs(0, 0)
	}
	return ans
}

func gen(word string) (int, bool) {
	ans := 0
	for _, ch := range word {
		if ans&(1<<(int(ch)-'a')) != 0 {
			return 0, false
		}
		ans = ans | 1<<(int(ch)-'a')
	}
	return ans, true
}
