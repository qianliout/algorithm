package main

import (
	"sort"
)

func main() {

}

// 这种方式不能得到正确的答案，因为有负数
func checkIfExist1(arr []int) bool {
	sort.Ints(arr)
	i, j := 0, len(arr)-1

	for i < j {
		if arr[i]*2 == arr[j] {
			return true
		} else if arr[i]*2 < arr[j] {
			i++
		} else if arr[i]*2 > arr[j] {
			j--
		}
	}
	return false
}

func checkIfExist(arr []int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]*2 == arr[j] || arr[j]*2 == arr[i] {
				return true
			}
		}
	}
	return false
}
