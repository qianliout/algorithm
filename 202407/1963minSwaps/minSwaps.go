package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSwaps("]]][[["))
}

func minSwaps1(s string) int {
	ans := 0
	cnt := 0
	for _, ch := range s {
		if ch == '[' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			cnt++
			ans++
		}
	}
	return ans
}

func minSwaps2(s string) int {
	ans := 0
	cnt := 0
	for _, ch := range s {
		if ch == '[' {
			cnt++
		} else if ch == ']' {
			if cnt > 0 {
				cnt--
			} else {
				ans++
				cnt++
			}
		}
	}
	return ans
}

// 使用栈,题目中说了是n个"["和n 个"]",所以把 n 个 "["匹配完成就好了
func minSwaps(s string) int {
	ans := 0
	st := make([]byte, 0)
	for _, ch := range s {
		if ch == '[' {
			st = append(st, '[')
		} else {
			if len(st) == 0 {
				ans++
				st = append(st, '[')
			} else {
				st = st[:len(st)-1]
			}
		}
	}
	return ans
}
