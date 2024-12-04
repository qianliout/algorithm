package main

import (
	"slices"
)

func main() {

}

func deleteAndEarn(nums []int) int {
	mx := slices.Max(nums)
	f := make([]int, mx+1)
	for _, ch := range nums {
		f[ch] += ch
	}
	dp1 := make([]int, mx+1) // yes
	dp2 := make([]int, mx+1) // no

	for i := 1; i <= mx; i++ {
		dp2[i+1] = dp1[i] - dp2[i-1] - i*f[i]

	}

}
