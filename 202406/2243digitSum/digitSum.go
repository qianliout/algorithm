package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(digitSum("11111222223", 3))
	fmt.Println(findClosestNumber([]int{-4, -2, 1, 4, 8}))
}

func digitSum(s string, k int) string {
	if len(s) <= k {
		return s
	}
	ans := make([]string, 0)

	for i := 0; i < len(s); i = i + k {
		ans = append(ans, s[i:min(len(s), i+k)])
	}
	for i, ch := range ans {
		ans[i] = sum(ch)
	}
	for i, ch := range ans {
		if len(ch) > k {
			ans[i] = digitSum(ch, k)
		}
	}
	nw := strings.Join(ans, "")
	return digitSum(nw, k)
}

func sum(s string) string {
	a := 0
	for _, ch := range s {
		a += int(ch) - '0'
	}
	return fmt.Sprintf("%d", a)
}

func findClosestNumber(nums []int) int {
	ans := make([]pair, 0)
	for _, ch := range nums {
		ans = append(ans, pair{v: ch, a: abs(ch)})
	}
	sort.Slice(ans, func(i, j int) bool {
		if ans[i].a < ans[j].a {
			return true
		} else if ans[i].a > ans[j].a {
			return false
		}
		return ans[i].v > ans[j].v
	})
	return ans[0].v
}

type pair struct {
	v, a int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
