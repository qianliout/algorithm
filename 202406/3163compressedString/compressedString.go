package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(compressedString("abcde"))
	fmt.Println(compressedString("aaaaaaaaaaaaaabb"))
}

func compressedString2(word string) string {
	nums := make([]int, 0)
	bys := make([]byte, 0)
	p, c := 1, word[0]
	for i := 1; i < len(word); i++ {
		if word[i-1] != word[i] {
			nums = append(nums, p)
			bys = append(bys, c)
			p, c = 1, word[i]
			continue
		}
		p++
	}
	if p > 0 {
		nums = append(nums, p)
		bys = append(bys, c)
	}

	ans := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		ans = append(ans, com(bys[i], nums[i]))
	}

	return strings.Join(ans, "")
}

func compressedString(word string) string {
	start := 0
	ans := make([]string, 0)
	for i := 0; i < len(word); i++ {
		if i+1 == len(word) || word[i+1] != word[i] {
			ans = append(ans, com(word[i], i-start+1))
			start = i + 1
		}
	}
	return strings.Join(ans, "")
}

// 拼接字符串是个耗时的操作，但是还没有超时
func com(c byte, b int) string {
	if b <= 0 {
		return ""
	}
	if b <= 9 {
		return fmt.Sprintf("%d%s", b, string(c))
	}
	return com(c, 9) + com(c, b-9)
}
