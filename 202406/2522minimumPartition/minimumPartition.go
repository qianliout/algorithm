package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumPartition("165462", 60))
}

func minimumPartition(s string, k int) int {
	ans := 0
	x := 0
	for _, c := range s {
		num := int(c - '0')
		if num > k {
			return -1
		}
		x = x*10 + num
		if x > k {
			ans++
			x = num
		}
	}
	if x > k {
		return -1
	}
	// 为啥要加一呢，因为 上面求的 ans 是有多少个分隔线
	return ans + 1
}
