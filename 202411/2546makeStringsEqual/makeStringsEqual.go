package main

import (
	"strings"
)

func main() {

}

func makeStringsEqual1(s string, target string) bool {
	//  s[i] = s[i]|s[j]  说明中只要有一个1 就能把 s 全变成 1
	//  s[j] = s[i]^s[j]  说明中只要有一个1 经过上一步全变成1，然后把其他数变成0（最后还会剩下一个1）
	//  也就是说 s 中只有有1就不能变成全（n-1）个0

	cnt1 := strings.Count(s, "1")

	cnt2 := strings.Count(target, "1")
	if cnt1 == 0 {
		return cnt2 == 0
	}
	return cnt2 > 0
}

func makeStringsEqual(s string, target string) bool {
	// 更简洁的写法
	return strings.Contains(s, "1") == strings.Contains(target, "1")
}
