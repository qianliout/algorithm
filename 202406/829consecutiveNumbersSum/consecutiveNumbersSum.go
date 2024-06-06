package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// 给定一个正整数 n，返回 连续正整数满足所有数字之和为 n 的组数 。
func consecutiveNumbersSum2(n int) int {
	n = n * 2
	ans := 0
	for k := 1; k*k < n; k++ {
		if n%k != 0 {
			continue
		}
		if (n/k-(k+1))%2 == 0 {
			ans++
		}
	}
	return ans
}

// 等差数列公式
/* 假设该数列最小值是a,数列个数是 k，那么 (a+(a+k-1))/2 =n
a

*/
func consecutiveNumbersSum(n int) int {
	ans := 0
	k := 1
	for k*(k+1) <= n*2 {
		if isKConsecutive(n, k) {
			ans++
		}
		k++
	}
	return ans
}

// k表示这个数列中元素的个数
func isKConsecutive(n, k int) bool {
	if k%2 > 0 {
		return n%k == 0
	}
	return n%k > 0 && 2*n%k == 0
}
