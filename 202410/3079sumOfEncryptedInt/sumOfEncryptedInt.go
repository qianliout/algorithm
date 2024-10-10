package main

import "fmt"

func main() {
	fmt.Println(sumOfEncryptedInt([]int{1, 2, 3}))
}

func sumOfEncryptedInt(nums []int) int {
	s := 0
	for _, ch := range nums {
		s += encrypt(ch)
	}
	return s
}

func encrypt(x int) int {
	a := 0
	cnt := 0
	for x > 0 {
		cnt++
		a = max(a, x%10)
		x = x / 10
	}
	b := 0
	for cnt > 0 {
		b = b*10 + a
		cnt--
	}
	return b
}
