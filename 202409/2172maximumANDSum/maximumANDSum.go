package main

import (
	"math/bits"
	"slices"
)

func main() {

}

func maximumANDSum(nums []int, numSlots int) int {
	n := len(nums)
	f := make([]int, 1<<(numSlots*2))
	for i, fi := range f {
		c := bits.OnesCount(uint(i))
		if c >= n {
			continue
		}
		for j := 0; j < numSlots*2; j++ {
			if i&(1<<j) == 0 {
				s := i | (1 << j)
				f[s] = max(f[s], fi+(j/2+1)&nums[c])
			}
		}
	}
	return slices.Max(f)
}
