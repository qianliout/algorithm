package main

import (
	"fmt"
	"strings"
)

func main() {
	a := strings.Split("00|00|00|00", "|")
	fmt.Println(a)
}

func maximumXOR(nums []int) int {
	// nums[i] XOR x，异或操作可以把0变成1，也可以把1变成0，所以当 x 取任意数时，也就是说 nums[i]可以取任意数
	// & 操作可以把1变成0，但是不能把0变成1，
	// 要想 XOR 的和最大，那么只需要每一位尽量是1就可以，也就是说，对每一位，只要有奇数个1，那么这一位的 xor 和就是1
	// 又因为nums[i] 可能是任何数，
	ans := 0
	for _, ch := range nums {
		ans = ans | ch
	}
	return ans
}

func countAsterisks(s string) int {
	// |0|0| split 后的结果是["","0","0",""]

	split := strings.Split(s, "|")

	n := len(split)
	all := strings.Count(s, "*")
	for i := 1; i < n; i = i + 2 {
		all -= strings.Count(split[i], "*")
	}
	return all
}
