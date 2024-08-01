package main

import (
	"fmt"
)

func main() {
	fmt.Println(areNumbersAscending("hello world 5 x 5"))
}

func areNumbersAscending(s string) bool {
	nums := find(s)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			return false
		}
	}
	return true
}

// 每个 token 要么是一个由数字 0-9 组成的不含前导零的 正整数
func find(s string) []int {
	pre := 0
	has := false
	ans := make([]int, 0)
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			has = true
			pre = pre*10 + int(ch) - int('0')
			continue
		} else {
			if has {
				ans = append(ans, pre)
				has = false
				pre = 0
			}
		}
	}
	if has {
		ans = append(ans, pre)
	}
	return ans
}
