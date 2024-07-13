package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minInteger("4321", 4))
}

func minInteger(num string, k int) string {
	for k > 0 {
		rev := reverse12(num)
		fmt.Println(rev)
		if rev == num {
			return rev
		}
		num = rev
		k--
	}
	return num
}

// 这样贪心的做法还搞不定
func reverse(num string) string {
	ss := []byte(num)
	n := len(ss)
	for i := 1; i < n; i++ {
		if ss[i] < ss[i-1] {
			ss[i], ss[i-1] = ss[i-1], ss[i]
			break
		}
	}
	return string(ss)
}

// 从最后找一小的。每次都往前面加
func reverse12(num string) string {
	ss := []byte(num)
	n := len(ss)
	mi := slices.Min(ss)
	if ss[0] == mi {
		return num
	}
	for i := 1; i < n; i++ {
		if ss[i] == mi {
			ss[i], ss[i-1] = ss[i-1], ss[i]
			break
		}
	}

	return string(ss)
}
