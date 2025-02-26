package main

import (
	"fmt"
)

func main() {
	fmt.Println(countDigitOne(13))
	fmt.Println(countDigitOne(20))
}

func countDigitOne1(n int) int {
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

// 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
// 示例 1：
// 输入：n = 13
// 输出：6
// 示例 2：
// 输入：n = 0
// 输出：0

// 为方便计算，首先把 n 转换成字符串 s。
// 定义 dfs(i,cnt,isLimit) 表示在前 i 位有 cnt个 1 的前提下，我们能构造出的数中的 1 的个数总和。
// 例如 n=9999，如果前三位我们都填了 1，那么填到最后一位，此时 dfs 计算的就是 1110,1111,1112,…,1119 这 10 个数中一共有多少个 1（一共有 31 个 1）。

func countDigitOne(n int) int {
	ss := fmt.Sprintf("%d", n)
	var dfs func(i, cnt int, isLimit bool) int
	mem1 := make([]map[int]int, len(ss)) // isLimit=true,记忆化这个没有意义，但是做题目时可以不用考虑这么深
	mem2 := make([]map[int]int, len(ss)) // isLimit=false
	for i := range mem1 {
		mem1[i] = make(map[int]int)
		mem2[i] = make(map[int]int)
	}

	// dfs 中的 isLimit 表示当前是否受到了 n 的约束（我们要构造的数字不能超过 n）
	dfs = func(i, cnt int, isLimit bool) int {
		if i == len(ss) {
			return cnt
		}
		if isLimit {
			if va, ok := mem1[i][cnt]; ok {
				return va
			}
		}
		if !isLimit {
			if va, ok := mem2[i][cnt]; ok {
				return va
			}
		}

		res := 0
		up := 9
		if isLimit {
			up = int(ss[i]) - int('0')
		}
		for d := 0; d <= up; d++ {
			a := dfs(i+1, cnt+cal(d), isLimit && d == up)
			res += a
		}
		if isLimit {
			mem1[i][cnt] = res
		}
		if !isLimit {
			mem2[i][cnt] = res
		}
		return res
	}
	return dfs(0, 0, true)
}

func cal(d int) int {
	if d == 1 {
		return 1
	}
	return 0
}
