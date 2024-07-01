package main

import (
	"fmt"
)

func main() {
	fmt.Println(appealSum("abbca"))
}

func appealSum(s string) int64 {
	ans := 0
	sum := 0
	last := make(map[byte]int) // 上次出现的下标
	for i, ch := range s {
		// 不能通过 last[byte(ch)]>0来判断，因为下标可能是0
		if _, ok := last[byte(ch)]; !ok {
			sum += i + 1
			ans += sum
		} else {
			pre := last[byte(ch)]
			sum += (i - pre - 1) + 1
			ans += sum
		}
		last[byte(ch)] = i
	}

	return int64(ans)
}
