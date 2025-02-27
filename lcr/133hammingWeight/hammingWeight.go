package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func hammingWeight2(n uint32) int {
	cnt := 0
	for n > 0 {
		if n&1 == 1 {
			cnt++
		}
		n = n >> 1
	}
	return cnt
}

func hammingWeight(n uint32) int {
	cnt := 0
	for n > 0 {
		low := n & -n
		if low > 0 {
			cnt++
		}
		n -= low
	}
	return cnt
}
