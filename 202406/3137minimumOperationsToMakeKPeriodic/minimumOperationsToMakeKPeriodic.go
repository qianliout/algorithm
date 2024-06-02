package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumOperationsToMakeKPeriodic("leetcoleet", 2))
	fmt.Println(minimumOperationsToMakeKPeriodic("leetcodeleet", 4))
}

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	// 题目中说，子串只能是从0开始，我们找到出现次数最多的子串，用这个子串去替换其他的子串
	ss := []byte(word)
	cnt := make(map[string]int)
	for i := 0; i+k <= len(ss); i = i + k {
		cnt[string(ss[i:i+k])]++
	}
	mx := 0
	for _, v := range cnt {
		mx = max(mx, v)
	}

	return len(word)/k - mx
}
