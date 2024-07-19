package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBalls(1, 10))
	fmt.Println(countBalls(5, 15))
}

func countBalls(lowLimit int, highLimit int) int {
	cnt := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		cnt[cal(i)]++
	}
	mx := 0
	ans := 0
	for _, v := range cnt {
		// 如果有多个盒子都满足放有最多小球，只需返回其中任一盒子的小球数量。
		if v <= mx {
			continue
		} else {
			mx = v
			ans = v
		}
	}
	return ans
}

func cal(n int) int {
	ans := 0
	for n > 0 {
		ans += n % 10
		n = n / 10
	}
	return ans
}
