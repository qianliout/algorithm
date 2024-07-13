package main

import (
	"slices"
)

func main() {

}

func getLastMoment(n int, left []int, right []int) int {

	if len(left) > 0 && len(right) > 0 {
		return max(slices.Max(left), n-slices.Min(right))
	}
	if len(left) > 0 {
		return slices.Max(left)
	}
	if len(right) > 0 {
		return n - slices.Min(right)
	}
	return 0
}
