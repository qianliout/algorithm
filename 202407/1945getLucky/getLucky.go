package main

import (
	"fmt"
)

func main() {
	// fmt.Println(getLucky("iiii", 1))
	fmt.Println(getLucky("vbyytoijnbgtyrjlsc", 2))
}

func getLucky(s string, k int) int {
	ans := 0
	for _, ch := range s {
		a := int(ch) - int('a') + 1
		for a > 0 {
			ans += a % 10
			a = a / 10
		}
	}

	// 这里要k>1,因为上面操作已经做了一次转换了
	// 1 <= k <= 10 k是大于等1的，所以可以这样做，但是如果 k 的范围可以取到0就不能这样做了
	for k > 1 {
		next := cal(ans)
		if ans == next {
			return next
		}
		ans = next
		k--
	}
	return ans
}

func same(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ch := range a {
		if b[i] != ch {
			return false
		}
	}
	return true
}

func cal(a int) int {
	ans := 0
	for a > 0 {
		ans += a % 10
		a = a / 10
	}
	return ans
}
