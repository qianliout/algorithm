package main

import (
	"sort"
)

func main() {

}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = min(arr[i-1]+1, arr[i])
	}

	return arr[n-1]
}
