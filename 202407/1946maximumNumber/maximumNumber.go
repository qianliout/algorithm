package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumNumber("021", []int{9, 4, 3, 5, 7, 2, 1, 9, 0, 6}))
	fmt.Println(maximumNumber("214010", []int{6, 7, 9, 7, 4, 0, 3, 4, 4, 7}))
	fmt.Println(maximumNumber("334111", []int{0, 9, 2, 3, 3, 2, 5, 5, 5, 5}))
}

func maximumNumber(num string, change []int) string {
	n := len(num)
	ss := []byte(num)
	for i := 0; i < n; i++ {

		if canChang(ss, i, change) {
			// 你可以选择 突变  num 的任一子字符串
			for j := i; j < n; j++ {
				b := int(ss[j]) - int('0')
				c := change[b]
				// 这里b<=c 都可以继续突变，为的就是突变更多。
				if b > c {
					break
				}
				ss[j] = byte('0' + change[int(ss[j])-int('0')])
			}

			return string(ss)
		}
	}
	return num
}

func canChang(num []byte, i int, change []int) bool {
	b := int(num[i]) - int('0')
	c := change[b]
	// 这里判断是否开始交换时只能用 b<c,因为只能变换任一一个子字符串
	if b < c {
		return true
	}
	return false
}
