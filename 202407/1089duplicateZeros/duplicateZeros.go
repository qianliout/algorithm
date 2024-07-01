package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
}

func duplicateZeros(arr []int) {
	cnt := 0
	for _, ch := range arr {
		if ch == 0 {
			cnt++
		}
	}
	n := len(arr)
	// 前面的变动会影响到后面，所以从后向前遍历
	for i := n - 1; i >= 0; i-- {
		if i+cnt < n {
			arr[i+cnt] = arr[i] // 移动
		}
		if arr[i] == 0 {
			cnt--
			if i+cnt < n {
				arr[i+cnt] = 0
			}
		}
	}
}
