package main

import (
	"math/bits"
)

func main() {

}

func countTriplets(arr []int) int {
	n := len(arr)
	sum := make([]int, n+1)
	for i, c := range arr {
		sum[i+1] = sum[i] ^ c
	}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			for k := j; k <= n; k++ {
				left := sum[j-1] ^ sum[i-1]
				right := sum[k] ^ sum[j-1]
				if left == right {
					ans++
				}
			}
		}
	}
	return ans
}
