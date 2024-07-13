package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMaxLen([]int{1, -2, -3, 4}))
	fmt.Println(getMaxLen([]int{-1, -2, -3, 0, 1}))
}

// 请你求出乘积为正数的最长子数组的长度。
func getMaxLen(nums []int) int {
	first := -1 // 第一次出现负数的 index
	pos, neg := 0, 0
	ans := 0
	for i, ch := range nums {
		if ch == 0 {
			pos, neg, first = 0, 0, -1
		} else if ch > 0 {
			pos++
		} else if ch < 0 {
			neg++
			if first == -1 {
				first = i
			}
		}
		if neg&1 == 0 {
			ans = max(ans, pos+neg)
		} else {
			ans = max(ans, i-first)
		}
	}

	return ans
}

/*
子数组乘积为正数，即要求该段子数组中没有0且负数的个数为偶数，这样我们可以用三个变量：
pos:该段正数个数
neg:负数个数
first:第一个负数出现的下标，初始化 -1
来记录需要的数量，然后对数组进行遍历：
1.如果当前neg % 2 = 0，说明该段组数组的所有元素相乘为正，
那么ans = max(ans, pos + neg)。
2.如果当前neg % 2 != 0，我们可以贪心的进行选取组数组，只要去掉该段字数组的一个负数便可以使负数个数为偶数，即乘积为正，这时，即从第一个负数开始，后面的位置到当前位置所有数的乘积可以为正，
此时:ans = max(ans, 当前位置下标 - fisrt).
3.如果当前数为0，则将所有标记初始化，因为0不可能包含在任何字数组中，是使得乘积为正
*/
