package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(numDupDigitsAtMostN(20))
	fmt.Println(numDupDigitsAtMostN(100))
	fmt.Println(numDupDigitsAtMostN(1000))
	fmt.Println(numDupDigitsAtMostN(415))
}

/*
给定正整数 n，返回在 [1, n] 范围内具有 至少 1 位 重复数字的正整数的个数。

正难则反，求一个都不重复的数量
*/
// 这是数位 dfs 的模板
func numDupDigitsAtMostN(n int) int {
	var dfs func(i, mask int, isLimit bool, isNum bool) int
	s := strconv.Itoa(n)
	mem := make([][]int, len(s))
	for i := range mem {
		// 这里是一个容易出错的点，这个数组在开多大，不能写1<<n，这样写会超限投影，,也不能写 1<<len(s)，会不够
		// 因为不能有重复的数据，所以这些数字中最多10个数字
		// 这里 mask 的意思是记录 0-9 这10个数用了那些，所以只能用1<<10 就行
		// 如果是不同的题目，这里要看数据结构，如果实在拿不准，可以直接用 map，但是效率不高，可能会超时
		mem[i] = make([]int, 1<<10)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	dfs = func(i, mask int, isLimit bool, isNum bool) int {
		if i >= len(s) {
			if isNum {
				return 1
			}
			return 0
		}
		res := 0
		// 这里的缓存是一个容易出错的点，这里只是记录前面是数字，并且不受限制的，因为，其他的情况下，不会再有第二次递归
		if isNum && !isLimit {
			if va := mem[i][mask]; va != -1 {
				return va
			}
		}

		if !isNum {
			// 跳过这一位
			res += dfs(i+1, mask, false, false)
		}
		low := 0 // 不能有前导0
		if !isNum {
			low = 1
		}
		up := int(s[i] - '0')
		if !isLimit {
			up = 9
		}
		for d := low; d <= up; d++ {
			// 判断是否在集合中容易出错的，他和下面 mask 的定义要一致，我们可以定义在集合中是0，本题定义的是在集合中是1，这也是最好的方式
			if mask>>d&1 == 0 {
				// isLimit 的赋值，也是容易出错的，这里要注意
				// 比如本题,假设n=234,如果前面填了2，这一位填3时，那一位就会受限制，但是如果这一位填的是2，那么下一位还是可以填0-9，是不受限制的
				res += dfs(i+1, mask|(1<<d), isLimit && d == up, true)
			}
		}
		mem[i][mask] = res
		return res
	}
	return n - dfs(0, 0, true, false)
}
