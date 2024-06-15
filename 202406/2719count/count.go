package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(count("1", "12", 1, 8))
	fmt.Println(count("4179205230", "7748704426", 8, 46))
}

func count(num1 string, num2 string, miSum int, mxSum int) int {
	mod := int(math.Pow10(9)) + 7
	a := cac(miSum, mxSum, num2)
	b := cac(miSum, mxSum, num1)
	ans := a - b + mod
	// ans := cac(miSum, mxSum, num2) - cac(miSum, mxSum, num1) + mod // 避免负数
	sum := 0
	for _, c := range num1 {
		sum += int(c - '0')
	}

	if miSum <= sum && sum <= mxSum { // num1 是合法的，补回来
		ans++
	}
	return ans % mod
}

func cac(miSum int, mxSum int, s string) int {
	mod := int(math.Pow10(9)) + 7
	mem := make([]map[int]int, len(s))
	for i := range mem {
		mem[i] = make(map[int]int)
	}

	var dfs func(i, sum int, limit bool) int
	dfs = func(i, sum int, limit bool) int {
		if sum > mxSum {
			return 0
		}
		if i > len(s) {
			return 0
		}

		if i == len(s) {
			if sum >= miSum {
				return 1
			}
			return 0
		}
		// 如果拿不准，就直接都缓存得了
		// 这里是一个很容器出错的点，我们没有对 limit 这个维度做缓存，是因为 limit=false 的情况不会访问第二次
		// 所以我们不会缓存 limit=false 的情况，只缓存了 limit=true 的情况
		if !limit {
			if va, ok := mem[i][sum]; ok {
				return va
			}
		}

		res := 0

		up := 9
		if limit {
			up = int(s[i] - '0')
		}
		for j := 0; j <= up; j++ {
			res = (res + dfs(i+1, sum+j, limit && j == up)) % mod
		}
		// 只缓存 limit=true 的情况,但是我们不能这样写，这样写之后就
		// if limit {
		mem[i][sum] = res % mod
		// }
		return res % mod
	}
	return dfs(0, 0, true)
}
