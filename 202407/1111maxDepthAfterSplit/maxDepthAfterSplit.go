package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDepthAfterSplit("()(())()"))
}

// 这个方法是求连续
func maxDepthAfterSplit2(seq string) []int {
	ss := []byte(seq)
	n := len(ss)
	res := make([]int, n)
	for i := range res {
		res[i] = 1
	}
	ans := n
	start, end := 0, 0
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if !check(ss[i+1 : j]) {
				continue
			}
			dep1 := cal(ss[i+1 : j])
			pre := append([]byte{}, ss[:i+1]...)
			pre = append(pre, ss[j:]...)

			dep2 := cal(pre)
			mx := max(dep1, dep2)
			if mx < ans {
				start, end = i, j
				ans = mx
			}
		}
	}
	for i := 0; i <= min(start, n-1); i++ {
		res[i] = 0
	}
	for i := n - 1; i >= max(0, end); i-- {
		res[i] = 0
	}

	return res
}

func check(seq []byte) bool {
	if len(seq) == 0 {
		return false
	}
	cnt := 0
	for _, ch := range seq {
		if ch == '(' {
			cnt++
		} else {
			cnt--
		}
	}
	return cnt == 0
}

func cal(seq []byte) int {

	if len(seq) == 0 {
		return 0
	}
	ans := 0

	cnt := 0

	for _, ch := range seq {
		if ch == '(' {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt--
		}
	}
	return ans
}

func maxDepthAfterSplit1(seq string) []int {
	n := len(seq)
	res := make([]int, n)
	stark := make([]int, 0)
	// dep := 0
	for i, ch := range seq {
		if ch == '(' {
			res[i] = len(stark) % 2
			stark = append(stark, i)
		} else {
			stark = stark[:len(stark)-1]
			res[i] = len(stark) % 2
		}
	}

	return res
}

// 字符串随便选，所以可以把奇数深度的括号放在 A，偶数深度放在 B
func maxDepthAfterSplit3(seq string) []int {
	n := len(seq)
	res := make([]int, n)
	dep := 0
	for i, ch := range seq {
		if ch == '(' {
			res[i] = dep % 2
			dep++
		} else {
			dep--
			res[i] = dep % 2
		}
	}

	return res
}

// 字符串随便选，所以可以把奇数深度的括号放在 A，偶数深度放在 B
func maxDepthAfterSplit(seq string) []int {
	n := len(seq)
	res := make([]int, n)
	dep := 0
	for i, ch := range seq {
		if ch == '(' {
			dep++
			// 奇数深度是1，偶数深度是0，题目没有区分0和1，这样写好理解一点
			res[i] = dep % 2
		} else {
			// 对于 ） 是看一个（ 的奇数还是偶数，所以先求值，再做--操作
			res[i] = dep % 2
			dep--
		}
	}

	return res
}
