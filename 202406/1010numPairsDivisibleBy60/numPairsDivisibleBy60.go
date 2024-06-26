package main

import (
	"fmt"
)

func main() {
	fmt.Println(canThreePartsEqualSum([]int{6, 1, 1, 13, -1, 0, -10, 20}))
}

func numPairsDivisibleBy60(nums []int) int {
	cnt := make(map[int]int)
	ans := 0
	for _, ch := range nums {
		ans += cnt[(60-ch%60)%60]
		cnt[ch%60]++
	}

	return ans
}

func canThreePartsEqualSum(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	sum := 0
	for _, ch := range arr {
		sum += ch
	}
	if sum%3 != 0 {
		return false
	}
	le := 0
	ri := n - 1
	// 为啥这里要先赋值呢，因为数组中有负值，可能sum/3就是0
	left := arr[0]
	right := arr[n-1]
	for le+1 < ri {
		if left == sum/3 && right == sum/3 {
			return true
		}

		if left != sum/3 {
			le++
			left += arr[le]
		}
		if right != sum/3 {
			ri--
			right += arr[ri]
		}
	}
	return false
}
