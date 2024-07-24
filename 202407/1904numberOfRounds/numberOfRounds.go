package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(numberOfRounds("12:01", "12:44"))
	fmt.Println(numberOfRounds("00:00", "23:59"))
	fmt.Println(numberOfRounds("20:00", "06:00"))
	fmt.Println(numberOfRounds("23:46", "00:01"))
	fmt.Println(numberOfRounds("09:31", "10:14"))
}

func numberOfRounds(loginTime string, logoutTime string) int {
	start := cal(loginTime)
	end := cal(logoutTime)
	if start > end {
		end += 60 * 24
	}
	// start 向上取整，找到第一个开始的时间
	start = (start + 15 - 1) / 15 * 15
	// 向下取整，找到最近的结束时间
	end = end / 15 * 15
	// 上面的操作可能使 start 小于 end，所以用个 max
	return max(0, end-start) / 15
}

func cal(s string) int {
	split := strings.Split(s, ":")
	h, _ := strconv.Atoi(split[0])
	m, _ := strconv.Atoi(split[1])
	return h*60 + m
}

// 这样写太麻烦了
// 没有过夜
func cal2(ti []int, start, end int) int {
	if end >= start {
		a := 0
		for i, ch := range ti {
			if ch >= start {
				a = i
				break
			}
		}
		b := len(ti) - 1
		for i := len(ti) - 1; i >= 0; i-- {
			if ti[i] <= end {
				b = i
				break
			}
		}
		return b - a
	}
	return 0
}

// 过夜了
func cal3(ti []int, start, end int) int {
	a := cal2(ti, start, cal("23:59"))
	b := cal2(ti, cal("00:00"), end)
	return a + b
}
