package main

import "fmt"

func main() {
	fmt.Println(maxLength([]string{"cha", "r", "act", "ers"}))
	fmt.Println(maxLength([]string{"aa", "bb"}))
}

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
	ans := 0
	var dfs func(i, pre int) int
	// visit := make([]bool, n)
	dfs = func(i, pre int) int {
		if i < 0 || i >= len(str) {
			return 0
		}
		// 选
		if nums[i]&pre == 0 {
			not := dfs(i+1, pre)
			do := dfs(i+1, pre|nums[i]) + len(str[i])
			ans = max(ans, not, do)
			return max(not, do)
		} else {
			not := dfs(i+1, pre)
			ans = max(ans, not)
			return max(not)
		}
	}
	dfs(0, 0)
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
