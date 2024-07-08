package main

import (
	"math/bits"
	"sort"
)

func main() {

}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a := bits.OnesCount(uint(arr[i]))
		b := bits.OnesCount(uint(arr[j]))
		if a < b {
			return true
		} else if a > b {
			return false
		} else {
			return arr[i] < arr[j]
		}
	})
	return arr
}
