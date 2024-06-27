package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(smallestRepunitDivByK(23))
	fmt.Println(smallestRepunitDivByK(3))
}

func smallestRepunitDivByK2(k int) int {
	seen := make(map[int]int)
	i := 1 % k
	for i > 0 && seen[i] == 0 {
		seen[i]++
		i = (i*10 + 1) % k
	}
	if i == 0 {
		return len(seen) + 1
	}
	return -1
}

// 直接模拟的结果竟然是不对的，为什么呢
// 因为数字会比 math.MaxInt64还要大
func smallestRepunitDivByK(k int) int {
	for i := 1; i <= (math.MaxInt64-1)/10; i = i*10 + 1 {
		if i%k == 0 {
			return len(strconv.Itoa(i))
		}
	}
	return -1
}
