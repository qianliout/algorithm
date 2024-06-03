package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(findMaximumNumber(9, 1))
	fmt.Println(findMaximumNumber(7, 2))
	fmt.Println(findMaximumNumber(4096, 6))
	fmt.Println(findMaximumNumber2(4096, 6))
	fmt.Println(findMaximumNumber(3278539330613, 5))
}

// 没有能理解
func findMaximumNumber2(k1 int64, x int) int64 {
	k := int(k1)
	var num int
	// pre1： 假设我们遍历到一个数a的第 i 位时，pre1代码这个数左边有多个少1
	var pre1 int
	// start 是一个初值，这里定义一个很大的数也可以，只是会有无效计算,可能会超时
	start := bits.Len(uint(k+1) << x)
	// start := math.MaxInt

	for i := start; i >= 0; i-- {
		cnt := pre1<<i + ((i/x)<<i)>>1
		if cnt > k {
			continue
		}
		k = k - cnt
		num = num | (1 << i)
		if (i+1)%x == 0 {
			pre1++
		}
	}

	return int64(num - 1)
}

func findMaximumNumber(K int64, x int) int64 {
	k := int(K)
	num, pre1 := 0, 0
	for i := bits.Len(uint((k+1)<<x)) - 1; i >= 0; i-- {
		cnt := pre1<<i + i/x<<i>>1
		if cnt > k {
			continue
		}
		k -= cnt
		num |= 1 << i
		if (i+1)%x == 0 {
			pre1++
		}
	}
	return int64(num - 1)
}

// 计算下于等于 n的数有多少个1
func countDigitOne(n int) int {
	count := 0
	for i := 1; i <= n; i = i * 10 {
		abc := n % i
		xyzd := n / i
		d := xyzd % 10
		xyz := xyzd / 10
		if d == 0 {
			count = count + xyz*i
		}
		if d == 1 {
			count = count + xyz*i + abc + 1
		}
		if d > 1 {
			count = count + xyz*i + i
		}

		//	防止溢出
		if xyz == 0 {
			break
		}
	}
	return count
}
