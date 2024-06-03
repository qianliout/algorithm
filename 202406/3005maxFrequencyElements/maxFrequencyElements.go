package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxFrequencyElements([]int{1, 2, 2, 3, 1, 4}))
}

func maxFrequencyElements(nums []int) int {
	frq := make(map[int]int) // 数字--->频率
	ans := 0
	mx := 0
	for _, ch := range nums {
		frq[ch]++
		c := frq[ch]
		if c > mx {
			mx = c
			ans = c
		} else if c == mx {
			ans += c
		}
	}
	return ans
}
