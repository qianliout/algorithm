package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println((-1) % (-2))
}

// n 是正数
func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	ans := make([]string, 0)
	for n != 0 {
		mod := n % -2
		n = n / -2
		if mod == -1 {
			// -1%-2 == -1 也就是 -1 = （-2）^0+(-1),但是二进制不可以是-10,所以要把-1变成1，当我们把余数从-1变成1也就是多加了一个2
			// 那么商也要多+1多用一个-2,把加的2抵消
			n++
			mod = 1
		}
		ans = append(ans, fmt.Sprintf("%d", mod))
	}

	le, ri := 0, len(ans)-1
	for le < ri {
		ans[le], ans[ri] = ans[ri], ans[le]
		le++
		ri--
	}

	return strings.Join(ans, "")
}

func prefixesDivBy5(nums []int) []bool {
	sum := 0

	n := len(nums)
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		sum = (sum*2 + nums[i]) % 5
		if sum%5 == 0 {
			ans[i] = true
		}
	}
	return ans
}
