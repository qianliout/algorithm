package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findDifferentBinaryString([]string{"0000000111", "0000001001", "0000000100", "0000000001", "0000000010", "1111111111", "0000000101", "0000000000", "0000001000", "0000000110"}))
}

func findDifferentBinaryString(nums []string) string {
	exist := make(map[int64]bool)
	for _, ch := range nums {
		nu := cal(ch)
		exist[nu] = true
	}
	n := len(nums[0])
	for i := 0; i < 1<<n; i++ {
		if !exist[int64(i)] {
			a := strconv.FormatInt(int64(i), 2)
			return strings.Repeat("0", n-len(a)) + a
		}
	}

	return strconv.FormatInt(1<<n, 2)
}

func cal(str string) int64 {
	ans := 0
	for i := 0; i < len(str); i++ {
		ans = ans*2 + int(str[i]) - int('0')
	}
	return int64(ans)
}
