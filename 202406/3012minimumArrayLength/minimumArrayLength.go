package main

import (
	"slices"
)

func main() {

}

/*
考虑这个例子：nums=[2,3,4,5,6]\textit{nums}=[2,3,4,5,6]nums=[2,3,4,5,6]每次操作都可以选择 2 和另一个数字 x，由于 x>2，所以 2mod x = 2，于是操作等价于：
移除 x。所以最后必定只会剩下 222。
*/
func minimumArrayLength(nums []int) int {
	mn := slices.Min(nums)
	cnt := 0
	for _, ch := range nums {
		if ch%mn > 0 {
			return 1
		}
		if ch == mn {
			cnt++
		}
	}
	return (cnt + 1) / 2
}
