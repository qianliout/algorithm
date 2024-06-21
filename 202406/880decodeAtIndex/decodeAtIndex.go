package main

import (
	"fmt"
)

func main() {
	fmt.Println(decodeAtIndex("ha22", 5))
	fmt.Println(decodeAtIndex("vzpp636m8y", 2920))
}

func decodeAtIndex(s string, k int) string {
	for k >= 0 {
		l := 0
		for i := 0; i < len(s); i++ {
			if s[i] >= 'a' && s[i] <= 'z' {
				l++
				if l == k {
					return string(s[i])
				}
			} else if l*int(s[i]-'0') >= k {
				// k是从1开始的
				// 题目保证 k 小于或等于解码字符串的长度。
				k = (k-1)%l + 1
				// 从头开始遍历
				break
			} else {
				l = l * (int(s[i] - '0'))
			}
		}
	}
	return ""
}
