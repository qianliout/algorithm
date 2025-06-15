package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))

}

// 错的：3,30,34,5,9
func largestNumber1(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a := fmt.Sprintf("%d", nums[i])
		b := fmt.Sprintf("%d", nums[j])
		return a > b
	})
	n := len(nums)
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = fmt.Sprintf("%d", nums[i])
	}
	return strings.Join(ans, "")
}

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a := fmt.Sprintf("%d", nums[i])
		b := fmt.Sprintf("%d", nums[j])
		return a+b > b+a
	})
	n := len(nums)
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = fmt.Sprintf("%d", nums[i])
	}
	res := strings.Join(ans, "")
	if strings.HasPrefix(res, "0") {
		return "0"
	}
	return res
}
