package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(kMirror(3, 7))
	fmt.Println(kMirror(7, 17))
	fmt.Println(toK(121, 3))
	fmt.Println(toK(151, 3))
	fmt.Println(toK(212, 3))
}

// 给你进制 k 和一个数字 n ，请你返回 k 镜像数字中 最小 的 n 个数 之和 。

func kMirror(k int, n int) int64 {
	ans := 0
	num := 0
	for n > 0 {
		num = next(num)
		tk := toK(num, k)

		if check(tk) {
			ans += num
			n--
		}
	}
	return int64(ans)
}

// 已知 n 是一个10进制的回文数，找下一个回文数

func next(m int) int {
	ss := []byte(strconv.Itoa(m))
	n := len(ss)

	for i := n / 2; i >= 0; i-- {
		if ss[i] != '9' {
			ss[i]++
			// 如果是最中间的数就不能增加两次
			if n-1-i != i {
				ss[n-1-i]++
			}
			// 中间的全部变成0
			for j := i + 1; j <= n/2; j++ {
				ss[j] = '0'
				ss[n-j-1] = '0'
			}
			ne, _ := strconv.Atoi(string(ss))
			return ne
		}
	}
	// 如果没有找到，就在增加一位
	res := make([]byte, n+1)
	for i := range res {
		res[i] = '0'
	}
	res[0] = '1'
	res[n] = '1'
	ne, _ := strconv.Atoi(string(res))
	return ne
}

func toK(n, k int) string {
	ans := make([]byte, 0)
	for n > 0 {
		a := n % k
		n = n / k
		ans = append(ans, byte('0'+int(a)))
	}
	// 这是一步容易出错的地方，容易忘记反转
	l, r := 0, len(ans)-1
	for l < r {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}

	return string(ans)
}

func check(ss string) bool {
	l, r := 0, len(ss)-1
	for l < r {
		if ss[l] != ss[r] {
			return false
		}
		l++
		r--
	}
	return true
}
