package main

import (
	"fmt"
)

func main() {

}

func compress(chars []byte) int {
	nums := make([]int, 0)
	bys := make([]byte, 0)
	p, c := 1, chars[0]
	for i := 1; i < len(chars); i++ {
		if chars[i-1] != chars[i] {
			nums = append(nums, p)
			bys = append(bys, c)
			p, c = 1, chars[i]
			continue
		}
		p++
	}
	if p > 0 {
		nums = append(nums, p)
		bys = append(bys, c)
	}
	start := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		s := com(bys[i], nums[i])
		for j := 0; j < len(s); j++ {
			chars[j+start] = s[j]
		}
		start += len(s)
		ans += len(s)
	}
	return ans
}

// 拼接字符串是个耗时的操作，但是还没有超时
func com(c byte, b int) string {
	if b <= 0 {
		return ""
	}
	if b == 1 {
		return string(c)
	}

	return fmt.Sprintf("%s%d", string(c), b)
}
