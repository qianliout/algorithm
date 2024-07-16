package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(numOfSubarrays([]int{1, 3, 5}))
	fmt.Println(numOfSubarrays([]int{1, 2, 3, 4, 5, 6, 7}))
}

func numOfSubarrays(arr []int) int {
	mod := int(math.Pow10(9)) + 7
	// 这里 even的初值是1,是容易出错的
	// 1 <= arr[i] <= 100
	odd, even, sum := 0, 1, 0
	ans := 0
	for _, ch := range arr {
		sum += ch
		if sum&1 == 0 {
			ans += odd
		} else {
			ans += even // 都是正数，不然不可以这样写
		}
		if sum&1 == 0 {
			even++
		} else {
			odd++
		}
	}
	return ans % mod
}
