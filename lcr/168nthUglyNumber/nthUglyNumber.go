package main

import (
	"fmt"
)

func main() {
	fmt.Println(nthUglyNumber(1690))
	fmt.Println(nthUglyNumber(16))
}

func nthUglyNumber2(n int) int {
	f := make([]int, n)
	f[0] = 1
	i2, i3, i5 := 0, 0, 0
	exit := make(map[int]bool)
	cnt := 1
	for cnt < n {
		le := min(f[i2]*2, f[i3]*3, f[i5]*5)
		if !exit[le] {
			f[cnt] = le
			exit[le] = true
			cnt++
		}

		if le == f[i2]*2 {
			i2++
		}
		if le == f[i3]*3 {
			i3++
		}
		if le == f[i5]*5 {
			i5++
		}
	}
	return f[n-1]
}

func nthUglyNumber(n int) int {
	if n <= 1 {
		return 1
	}
	f := make([]int, n)
	f[0] = 1
	i2, i3, i5 := 0, 0, 0
	exit := make(map[int]bool)
	cnt := 1
	for cnt < n {
		next := min(f[i2]*2, f[i3]*3, f[i5]*5)
		// 一定要去重
		if !exit[next] {
			f[cnt] = next
			exit[next] = true
			cnt++
		}
		// 如果已重复了，也要做下边这些步骤
		if f[i2]*2 == next {
			i2++
		}
		if f[i3]*3 == next {
			i3++
		}
		if f[i5]*5 == next {
			i5++
		}
	}
	return f[n-1]
}

// 根据题意，每个丑数都可以由其他较小的丑数通过乘以2或3或5得到。
// 所以，可以考虑使用一个优先队列保存所有的丑数，每次取出最小的那个，然后乘以2,3,5后放回队列。然而，这样做会出现重复的丑数。例如：
// 初始化丑数列表 [1]
// 第一轮： 1 -> 2, 3, 5 ，丑数列表变为 [1, 2, 3, 5]
// 第二轮： 2 -> 4, 6, 10 ，丑数列表变为 [1, 2, 3, 4, 6, 10]
// 第三轮： 3 -> 6, 9, 15 ，出现重复的丑数 6
// 为了避免重复，我们可以用三个指针a,b,c，分别表示下一个丑数是当前指针指向的丑数乘以2,3,5。
// 利用三个指针生成丑数的算法流程：
// 初始化丑数列表res，首个丑数为1，三个指针a,b,c都指向首个丑数。
// 开启循环生成丑数：
// 计算下一个丑数的候选集res[a]⋅2,res[b]⋅3,res[c]⋅5。
// 选择丑数候选集中最小的那个作为下一个丑数，填入res。
// 将被选中的丑数对应的指针向右移动一格。
// 返回res的最后一个元素即可。
