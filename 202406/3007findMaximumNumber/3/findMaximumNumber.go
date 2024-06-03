package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMaximumNumber(9, 1))
	fmt.Println(findMaximumNumber(7, 2))
	fmt.Println(findMaximumNumber(4096, 6))
	fmt.Println(findMaximumNumber(3278539330613, 5))
}

// 没有能理解
func findMaximumNumber(k int64, x int) int64 {
	N := 52
	t := make([]int64, N)
	for i := 0; i < N; i++ {
		if (i+1)%x == 0 {
			t[i]++
		}
		if i == 0 {
			continue
		}
		t[i] += 2 * t[i-1]

		if i%x == 0 {
			t[i] += (1 << (i - 1)) - 2
		}
	}

	var ans, cnt int64
	for i := N - 1; i > -1; i-- {
		cost := t[i] + cnt*(1<<i)
		if cost <= k {
			k -= cost
			ans |= 1 << i
			if (i+1)%x == 0 {
				cnt++
			}
		}
	}

	return ans
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
