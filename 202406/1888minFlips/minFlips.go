package main

import (
	"fmt"
)

func main() {
	fmt.Println(minFlips("111000"))
	fmt.Println(minFlips("0001100010101000111101000110101111000000101100000001001"))
}

// 很朴素的做法，但是会超时
func minFlips(s string) int {
	n := len(s)
	cnt := 0
	target := "01"
	// 假如以 「0101010.....」的方式进行匹配，相差 cnt 个元素，那么以 [101010101...]的方式就是 n-cnt
	for i := 0; i < n; i++ {
		if s[i] != target[i%2] {
			cnt++
		}
	}
	ans := min(cnt, n-cnt)
	// 解决第一种：类型 1 ：删除 字符串 s 的第一个字符并将它 添加 到字符串结尾，其实就是循环字符串
	for i := 0; i < n; i++ {
		if s[i] != target[i%2] {
			cnt--
		}
		if s[i] != target[(i+n)%2] {
			cnt++
		}
		ans = min(ans, cnt, n-cnt)
	}

	return ans
}

func cac(by []byte) int {
	ans1, ans2 := 0, 0
	for i := 0; i < len(by); i++ {
		if int(by[i]-'0') != i%2 {
			ans1++
		}
		if int(by[i]-'0') != (i+1)%2 {
			ans2++
		}
	}
	return min(ans1, ans2)
}
