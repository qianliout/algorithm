package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingWeight(11))
	fmt.Println(hammingWeight(128))
	fmt.Println(hammingWeight(2147483645))
}

func hammingWeight(n int) int {
	ans := 0
	for n != 0 {
		if n&1 == 1 {
			ans++
		}
		n = n >> 1
	}
	return ans
}

func hammingWeight2(num uint32) int {
	ans := 0
	for num != 0 {
		ans++
		num = num & (num - 1)
	}

	return ans
}
