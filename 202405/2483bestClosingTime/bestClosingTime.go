package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(bestClosingTime("NNN"))
	fmt.Println(bestClosingTime("YYY"))
	fmt.Println(bestClosingTime("YYNY"))
}

// 会超时
func bestClosingTime1(customers string) int {
	c := customers
	ans := len(c)
	idx := len(c)
	for i := 0; i < len(c)+1; i++ {
		cnt := count(c, i)
		if cnt < ans {
			idx = i
			ans = cnt
		}
	}

	return idx
}

func count(c string, start int) int {
	ans := 0
	for i := 0; i < len(c)+1; i++ {
		if i < start {
			if i < len(c) && c[i] == 'N' {
				ans++
			}
		}
		if i >= start {
			if i < len(c) && c[i] == 'Y' {
				ans++
			}
		}
	}
	return ans
}

func bestClosingTime(customers string) int {
	// 假如从0点就关门,代价就是所有的 Y
	cost := strings.Count(customers, "Y")
	mi := cost
	idx := 0
	for i := 1; i <= len(customers); i++ {
		ch := customers[i-1]
		// 假如知道了 i-1的时间关门的代价是 c，那么在 i 这个时间关门的代码是多少呢：
		// - 如果i-i 是 Y 那么 i的代码就是 c-1,如是 i-i 是 N，那么就是 c+1

		if ch == 'Y' {
			cost--
			if cost < mi {
				mi = cost
				idx = i
			}
		} else {
			cost++
		}
	}
	return idx
}
