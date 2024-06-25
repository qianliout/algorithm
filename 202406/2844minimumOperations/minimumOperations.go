package main

import (
	"strings"
)

func main() {

}

/*
一个数能被 25 整除，有如下五种情况：

	这个数是 0。
	这个数的末尾是 00。
	这个数的末尾是 25。
	这个数的末尾是 50。
	这个数的末尾是 75。
*/

func minimumOperations(num string) int {
	ans := len(num)
	if strings.Contains(num, "0") {
		ans--
	}

	findTal(num, "00", &ans)
	findTal(num, "25", &ans)
	findTal(num, "50", &ans)
	findTal(num, "75", &ans)
	return ans
}

func findTal(nums string, tail string, ans *int) {
	i := strings.LastIndex(nums, string(tail[1]))
	if i < 0 {
		return
	}
	i = strings.LastIndex(string(nums[:i]), string(tail[0]))
	if i < 0 {
		return
	}
	*ans = min(*ans, len(nums)-i-2)
}
