package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sumSubarrayMins2([]int{3, 1, 2, 4}))
	// fmt.Println(sumSubarrayMins2([]int{11, 81, 94, 43, 3}))
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
	// fmt.Println(sumSubarrayMins([]int{11, 81, 94, 43, 3}))
}

// 从单调队列就学到的“弹它”图，以及及时移除无用数据的核心原则，灵神是我的偶像！
func sumSubarrayMins2(arr []int) int {
	base := int(math.Pow(10, 9)) + 7
	left := make([]int, len(arr))
	right := make([]int, len(arr))
	st := make([]int, 0)
	st = append(st, -1)
	// 找右边界，严格小于等于的边界,也就是说找到等于之后就不向右
	for i, ch := range arr {
		for len(st) > 1 && arr[st[len(st)-1]] >= ch {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]

		st = append(st, i)
	}
	st = st[:0]
	st = append(st, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		for len(st) > 1 && arr[st[len(st)-1]] > arr[i] {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}
	ans := 0
	fmt.Println("left", left, "right", right)
	for i, ch := range arr {
		ans += ch * (i - left[i]) * (right[i] - i)
	}
	return ans % base
}

func sumSubarrayMins(arr []int) int {
	base := int(math.Pow(10, 9)) + 7
	right := make([]int, len(arr))

	st := make([]int, 0)
	st = append(st, -1)

	for i, ch := range arr {
		for len(st) > 1 && ch <= arr[st[len(st)-1]] {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}

	st = st[:0]
	st = append(st, len(arr))
	left := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		for len(st) > 1 && arr[i] < arr[st[len(st)-1]] {
			left[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}

	fmt.Println("left", left, "right", right)
	ans := 0
	for i, ch := range arr {
		ans += ch * (i - left[i]) * (right[i] - i)
	}
	return ans % base
}
