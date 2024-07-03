package main

import (
	"sort"
)

func main() {

}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	ids := make(map[int]int)
	for i, c := range arr2 {
		ids[c] = i + 1
	}

	n := len(arr1)
	ids2 := make(map[int]int)
	for _, ch := range arr1 {
		ids2[ch] = ids[ch]
		if ids2[ch] == 0 {
			ids2[ch] = n
		}
	}

	sort.Slice(arr1, func(i, j int) bool {
		if ids2[arr1[i]] < ids2[arr1[j]] {
			return true
		} else if ids2[arr1[i]] > ids2[arr1[j]] {
			return false
		} else {
			return arr1[i] < arr1[j]
		}
	})
	return arr1
}
