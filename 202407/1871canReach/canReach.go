package main

import (
	"fmt"
)

func main() {
	fmt.Println(canReach("011010", 2, 3))
}

// 超时
func canReach2(s string, minJump int, maxJump int) bool {
	n := len(s)
	ans := make([]bool, n)
	ans[0] = true
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			continue
		}
		if !ans[i] {
			continue
		}
		for k := i + minJump; k <= min(i+maxJump, n-1); k++ {
			if s[k] == '0' {
				ans[k] = true
			}
		}
	}
	return ans[n-1]
}

// 使用前缀和的方式优化
func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	ans := make([]int, n)
	pre := make([]int, n)
	ans[0] = 1 // 1 表示能跳到，0 表示不能跳到
	pre[0] = 1
	for i := 1; i < minJump; i++ {
		pre[i] = 1
	}

	for i := minJump; i < n; i++ {
		left, right := i-maxJump, i-minJump
		if s[i] == '0' {
			total := pre[right] - 0
			if left-1 >= 0 {
				total = pre[right] - pre[left-1]
			}
			if total != 0 {
				ans[i] = 1
			}
		}
		pre[i] = pre[i-1] + ans[i]
	}
	return ans[n-1] == 1
}
