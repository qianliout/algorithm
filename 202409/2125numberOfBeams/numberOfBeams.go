package main

import (
	"strings"
)

func main() {

}

func numberOfBeams(bank []string) int {
	pre := 0
	ans := 0
	for _, ch := range bank {
		cnt := strings.Count(ch, "1")
		if cnt > 0 {
			ans += pre * cnt
			pre = cnt
		}
	}
	return ans
}
