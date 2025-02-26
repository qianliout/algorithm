package main

import (
	"fmt"
)

func main() {

}

func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		start = start * 10
		digit++
		count = 9 * digit * start
	}
	num := start + (n-1)/digit
	a := fmt.Sprintf("%d", num)
	return int(a[(n-1)%digit]) - int('0')
}

// 根据以上分析，可将求解分为三步：
//
// 确定 n 所在 数字 的 位数 ，记为 digit 。
// 确定 n 所在的 数字 ，记为 num 。
// 确定 n 是 num 中的哪一数位，并返回结果。

/*
class Solution:
    def findNthDigit(self, n: int) -> int:
        digit, start, count = 1, 1, 9
        while n > count: # 1.
            n -= count
            start *= 10
            digit += 1
            count = 9 * start * digit
        num = start + (n - 1) // digit # 2.
        return int(str(num)[(n - 1) % digit]) # 3.
*/
