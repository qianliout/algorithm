package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(maxValue(6773685, 5166078, 49851224))
	fmt.Println(maxValue(995610677, 934568761, 999009430))
}

// nums[index] 的值被 最大化,返回 nums[index]
func maxValue(n int, index int, maxSum int) int {
	j := sort.Search(maxSum, func(k int) bool { return cal(n, index, k) >= maxSum })
	fmt.Println(j)
	// todo 这里如果自已实现，应该怎么写
	le, ri := 1, maxSum+1
	for le < ri {
		mid := le + (ri-le)/2
		if mid >= 1 && mid < maxSum+1 && cal(n, index, mid) >= maxSum {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	// 不做验证就能通过
	// nums 中所有元素之和不超过 maxSum 说明这样写还是有问题的
	fmt.Println(cal(n, index, le)-maxSum, le)
	if cal(n, index, le) <= maxSum {
		return le
	}
	return 1
}

// 一个数组，最大值是big，最小值是1，那么这个数组的和
func f(big, length int) int {
	if length <= 0 {
		return 0
	}
	if length <= big {
		return (2*big + 1 - length) * length / 2
	}
	return (big+1)*big/2 + length - big
}

// 当 idx 的值是 mx 时，数组和的最小值是多少
// 这样计会超时
func cal(n, idx, mx int) int {
	left := idx
	right := n - idx - 1
	return mx + f(mx, left) + f(mx, right)

	// 这样计算会超时
	// ans := mx
	// a, b := mx-1, mx-1
	// for i := idx - 1; i >= 0; i-- {
	// 	if a > 1 {
	// 		ans += a
	// 		a--
	// 	} else {
	// 		ans += i + 1
	// 		break
	// 	}
	// }
	// for i := idx + 1; i < n; i++ {
	// 	if b > 1 {
	// 		ans += b
	// 		b--
	// 	} else {
	// 		ans += n - i
	// 		break
	// 	}
	// }
	// return ans
}
