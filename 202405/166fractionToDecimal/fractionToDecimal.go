package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(fractionToDecimal(4, -333))
	fmt.Println(fractionToDecimal(1, -2))
	fmt.Println(fractionToDecimal(1, 6))
}

func fractionToDecimal(numerator int, denominator int) string {
	return dec(int64(numerator), int64(denominator))
}

// 正数需要转负数，这里用int64防止越界
func dec(a, b int64) string {
	if a < 0 && b < 0 {
		return dec(-a, -b)
	}
	if a < 0 && b > 0 {
		return "-" + dec(-a, b)
	}
	if a > 0 && b < 0 {
		return "-" + dec(a, -b)
	}
	// 到这里，两个数都是正数了
	c := a / b

	if a%b == 0 {
		return strconv.Itoa(int(c))
	}

	exit := make(map[int64]int) // 余数是否出现过
	residue := a % b
	ss := make([]string, 0)
	for residue > 0 {
		exit[residue] = len(ss)
		residue = residue * 10
		ss = append(ss, strconv.Itoa(int(residue/b)))
		residue = residue % b
		// 说明有循环了
		if pre, ok := exit[residue]; ok {
			return fmt.Sprintf("%d.%s(%s)", c, strings.Join(ss[:pre], ""), strings.Join(ss[pre:], ""))
		}
	}
	return fmt.Sprintf("%d.%s", c, strings.Join(ss, ""))
}
