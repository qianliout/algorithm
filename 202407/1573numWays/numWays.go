package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(numWays("100100010100110"))
}

/*
	方法

1001001000
从前往后取满一之后，后面就可以补0，
第一组可以取1，10，100，
第二组可以取，1，10，100
*/
func numWays(s string) int {
	n := len(s)
	mod := int(math.Pow10(9)) + 7
	cnt := strings.Count(s, "1")
	if cnt%3 != 0 || n <= 2 {
		return 0
	}
	if cnt == 0 {
		return (n - 1) * (n - 2) / 2 % mod
	}
	a := cnt / 3
	return cal1(s, a) * cal2(s, a) % mod
}

// 从前向后取
func cal1(s string, cnt int) int {
	cur := 0
	ret := 1
	for i, ch := range s {
		if ch == '1' {
			cur++
		}
		if cur == cnt {
			for j := i + 1; j < len(s) && s[j] == '0'; j++ {
				ret++
			}
			break
		}
	}
	return ret
}

// 从后向前取
func cal2(s string, cnt int) int {
	cur := 0
	ret := 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			cur++
		}
		if cur == cnt {
			for j := i - 1; j >= 0 && s[j] == '0'; j-- {
				ret++
			}
			break
		}
	}
	return ret
}
