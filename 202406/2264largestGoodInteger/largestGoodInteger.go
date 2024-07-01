package main

import (
	"strings"
)

func main() {

}

func largestGoodInteger(num string) string {
	ans := ""
	n := len(num)
	ss := []byte(num)
	for i := 0; i+3 < n; i++ {
		if string(ss[i:i+3]) == strings.Repeat(string(ss[i]), 3) {
			if ans == "" || (ans != "" && ans < string(ss[i:i+3])) {
				ans = string(ss[i : i+3])
			}
		}
	}

	return ans
}
