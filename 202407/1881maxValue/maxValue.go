package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxValue("-13", 2))
	fmt.Println(maxValue("-132", 3))
}

func maxValue(n string, x int) string {
	if n[0] == '-' {
		return "-" + cal1(n[1:], x)
	}
	return cal2(n, x)
}

// 变成最小的
func cal1(s string, x int) string {
	n := len(s)
	// 比第一个都小就加在最前面
	if int(s[0])-int('0') > x {
		return fmt.Sprintf("%d%s", x, s)
	}

	for i := 0; i < n; i++ {
		if int(s[i])-int('0') > x {
			return fmt.Sprintf("%s%d%s", s[:i], x, s[i:])
		}
	}

	return fmt.Sprintf("%s%d", s, x)
}

// 变成最大的
func cal2(s string, x int) string {
	n := len(s)
	// 比第一个都小就加在最前面
	if int(s[0])-int('0') < x {
		return fmt.Sprintf("%d%s", x, s)
	}

	for i := 0; i < n; i++ {
		if int(s[i])-int('0') < x {
			return fmt.Sprintf("%s%d%s", s[:i], x, s[i:])
		}
	}

	return fmt.Sprintf("%s%d", s, x)
}
