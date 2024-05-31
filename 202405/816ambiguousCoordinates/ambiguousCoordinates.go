package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(ambiguousCoordinates("(123)"))
	fmt.Println(ambiguousCoordinates("(0123)"))
	fmt.Println(ambiguousCoordinates("(00011)"))
	fmt.Println(parse([]byte("0001")))
}

func ambiguousCoordinates(s string) []string {
	ans := make([]string, 0)
	if strings.Count(s, "(") != 1 || strings.Count(s, ")") != 1 {
		return ans
	}
	s = strings.ReplaceAll(strings.ReplaceAll(s, "(", ""), ")", "")
	ss := []byte(s)

	for i := 1; i < len(ss); i++ {
		left := parse(ss[:i])
		right := parse(ss[i:])
		for _, le := range left {
			for _, ri := range right {
				// 输出的答案有个空格
				ans = append(ans, fmt.Sprintf("(%s, %s)", le, ri))
			}
		}
	}
	return ans
}

// ["(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"]
func parse(ss []byte) []string {
	ans := make([]string, 0)
	for i := 1; i <= len(ss); i++ {
		fir := ss[:i]
		a, _ := strconv.Atoi(string(fir))

		aa := strconv.Itoa(a) == string(fir)
		if !aa {
			continue
		}

		if aa {
			if len(ss[i:]) == 0 {
				ans = append(ans, fmt.Sprintf("%s", string(ss[:i])))
				continue
			}
		}

		sec := rever(ss[i:])
		b, _ := strconv.Atoi(string(sec))

		if strconv.Itoa(b) == string(sec) && b != 0 {
			ans = append(ans, fmt.Sprintf("%s.%s", string(ss[:i]), string(ss[i:])))
		}
	}
	return ans
}

func rever(ans []byte) []byte {
	str := make([]byte, 0)
	str = append(str, ans...)
	le, ri := 0, len(str)-1
	for le < ri {
		str[le], str[ri] = str[ri], str[le]
		le++
		ri--
	}
	return str
}
